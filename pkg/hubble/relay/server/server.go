// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package server

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	observerpb "github.com/cilium/cilium/api/v1/observer"
	"github.com/cilium/cilium/pkg/hubble/peer"
	peerTypes "github.com/cilium/cilium/pkg/hubble/peer/types"
	"github.com/cilium/cilium/pkg/hubble/relay/defaults"
	"github.com/cilium/cilium/pkg/hubble/relay/observer"
	"github.com/cilium/cilium/pkg/hubble/relay/pool"
	"github.com/cilium/cilium/pkg/logging/logfields"
)

var (
	// ErrNoClientTLSConfig is returned when no client TLS config is set unless
	// WithInsecureClient() is provided.
	ErrNoClientTLSConfig = errors.New("no client TLS config is set")
	// ErrNoServerTLSConfig is returned when no server TLS config is set unless
	// WithInsecureServer() is provided.
	ErrNoServerTLSConfig = errors.New("no server TLS config is set")

	registry = prometheus.NewPedanticRegistry()
)

func init() {
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	registry.MustRegister(collectors.NewGoCollector())
}

// Server is a proxy that connects to a running instance of hubble gRPC server
// via unix domain socket.
type Server struct {
	server           *grpc.Server
	grpcHealthServer *grpc.Server
	pm               *pool.PeerManager
	healthServer     *healthServer
	metricsServer    *http.Server
	opts             options
}

// New creates a new Server.
func New(options ...Option) (*Server, error) {
	opts := defaultOptions // start with defaults
	options = append(options, DefaultOptions...)
	for _, opt := range options {
		if err := opt(&opts); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}
	if opts.clientTLSConfig == nil && !opts.insecureClient {
		return nil, ErrNoClientTLSConfig
	}
	if opts.serverTLSConfig == nil && !opts.insecureServer {
		return nil, ErrNoServerTLSConfig
	}

	var peerClientBuilder peerTypes.ClientBuilder = &peerTypes.LocalClientBuilder{}
	if !strings.HasPrefix(opts.peerTarget, "unix://") {
		peerClientBuilder = &peerTypes.RemoteClientBuilder{
			TLSConfig:     opts.clientTLSConfig,
			TLSServerName: peer.TLSServerName(defaults.PeerServiceName, opts.clusterName),
		}
	}

	pm, err := pool.NewPeerManager(
		registry,
		pool.WithPeerServiceAddress(opts.peerTarget),
		pool.WithPeerClientBuilder(peerClientBuilder),
		pool.WithClientConnBuilder(pool.GRPCClientConnBuilder{
			TLSConfig: opts.clientTLSConfig,
		}),
		pool.WithRetryTimeout(opts.retryTimeout),
		pool.WithLogger(opts.log),
	)
	if err != nil {
		return nil, err
	}

	var serverOpts []grpc.ServerOption

	if len(opts.grpcUnaryInterceptors) > 0 {
		serverOpts = append(serverOpts, grpc.ChainUnaryInterceptor(opts.grpcUnaryInterceptors...))
	}
	if len(opts.grpcStreamInterceptors) > 0 {
		serverOpts = append(serverOpts, grpc.ChainStreamInterceptor(opts.grpcStreamInterceptors...))
	}

	if opts.serverTLSConfig != nil {
		tlsConfig := opts.serverTLSConfig.ServerConfig(&tls.Config{
			MinVersion: MinTLSVersion,
		})
		serverOpts = append(serverOpts, grpc.Creds(credentials.NewTLS(tlsConfig)))
	}
	grpcServer := grpc.NewServer(serverOpts...)
	grpcHealthServer := grpc.NewServer()

	observerOptions := copyObserverOptionsWithLogger(opts.log, opts.observerOptions)
	observerSrv, err := observer.NewServer(pm, observerOptions...)
	if err != nil {
		return nil, fmt.Errorf("failed to create observer server: %w", err)
	}
	healthSrv := newHealthServer(pm, defaults.HealthCheckInterval)

	observerpb.RegisterObserverServer(grpcServer, observerSrv)
	healthpb.RegisterHealthServer(grpcServer, healthSrv.svc)
	healthpb.RegisterHealthServer(grpcHealthServer, healthSrv.svc)
	reflection.Register(grpcServer)

	if opts.grpcMetrics != nil {
		registry.MustRegister(opts.grpcMetrics)
		opts.grpcMetrics.InitializeMetrics(grpcServer)
	}

	var metricsServer *http.Server
	if opts.metricsListenAddress != "" {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
		metricsServer = &http.Server{
			Addr:    opts.metricsListenAddress,
			Handler: mux,
		}
	}

	return &Server{
		pm:               pm,
		server:           grpcServer,
		grpcHealthServer: grpcHealthServer,
		metricsServer:    metricsServer,
		healthServer:     healthSrv,
		opts:             opts,
	}, nil
}

// Serve starts the hubble-relay server. Serve does not return unless a
// listening fails with fatal errors. Serve will return a non-nil error if
// Stop() is not called.
func (s *Server) Serve() error {
	var eg errgroup.Group
	if s.metricsServer != nil {
		eg.Go(func() error {
			s.opts.log.Info("Starting metrics server...", logfields.Address, s.opts.metricsListenAddress)
			return s.metricsServer.ListenAndServe()
		})
	}

	eg.Go(func() error {
		s.opts.log.Info("Starting gRPC server...", logfields.Options, s.opts)
		s.pm.Start()
		s.healthServer.start()
		socket, err := net.Listen("tcp", s.opts.listenAddress)
		if err != nil {
			return fmt.Errorf("failed to listen on tcp socket %s: %w", s.opts.listenAddress, err)
		}
		return s.server.Serve(socket)
	})

	eg.Go(func() error {
		s.opts.log.Info("Starting gRPC health server...", logfields.Address, s.opts.healthListenAddress)
		socket, err := net.Listen("tcp", s.opts.healthListenAddress)
		if err != nil {
			return fmt.Errorf("failed to listen on %s: %w", s.opts.healthListenAddress, err)
		}
		return s.grpcHealthServer.Serve(socket)
	})

	return eg.Wait()
}

// Stop terminates the hubble-relay server.
func (s *Server) Stop() {
	s.opts.log.Info("Stopping server...")
	s.server.Stop()
	if s.metricsServer != nil {
		if err := s.metricsServer.Shutdown(context.Background()); err != nil {
			s.opts.log.Info("Failed to gracefully stop metrics server", logfields.Error, err)
		}
	}
	s.pm.Stop()
	s.healthServer.stop()
	s.opts.log.Info("Server stopped")
}

// observerOptions returns the configured hubble-relay observer options along
// with the hubble-relay logger.
func copyObserverOptionsWithLogger(log *slog.Logger, options []observer.Option) []observer.Option {
	newOptions := make([]observer.Option, len(options), len(options)+1)
	copy(newOptions, options)
	newOptions = append(newOptions, observer.WithLogger(log))
	return newOptions
}
