// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v5.29.2
// source: envoy/extensions/filters/http/health_check/v3/health_check.proto

package health_checkv3

import (
	v31 "github.com/cilium/proxy/go/envoy/config/route/v3"
	v3 "github.com/cilium/proxy/go/envoy/type/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// [#next-free-field: 6]
type HealthCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies whether the filter operates in pass through mode or not.
	PassThroughMode *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=pass_through_mode,json=passThroughMode,proto3" json:"pass_through_mode,omitempty"`
	// If operating in pass through mode, the amount of time in milliseconds
	// that the filter should cache the upstream response.
	CacheTime *durationpb.Duration `protobuf:"bytes,3,opt,name=cache_time,json=cacheTime,proto3" json:"cache_time,omitempty"`
	// If operating in non-pass-through mode, specifies a set of upstream cluster
	// names and the minimum percentage of servers in each of those clusters that
	// must be healthy or degraded in order for the filter to return a 200.
	//
	// .. note::
	//
	//	This value is interpreted as an integer by truncating, so 12.50% will be calculated
	//	as if it were 12%.
	ClusterMinHealthyPercentages map[string]*v3.Percent `protobuf:"bytes,4,rep,name=cluster_min_healthy_percentages,json=clusterMinHealthyPercentages,proto3" json:"cluster_min_healthy_percentages,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Specifies a set of health check request headers to match on. The health check filter will
	// check a request’s headers against all the specified headers. To specify the health check
	// endpoint, set the “:path“ header to match on.
	Headers []*v31.HeaderMatcher `protobuf:"bytes,5,rep,name=headers,proto3" json:"headers,omitempty"`
}

func (x *HealthCheck) Reset() {
	*x = HealthCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_health_check_v3_health_check_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheck) ProtoMessage() {}

func (x *HealthCheck) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_health_check_v3_health_check_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheck.ProtoReflect.Descriptor instead.
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_health_check_v3_health_check_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCheck) GetPassThroughMode() *wrapperspb.BoolValue {
	if x != nil {
		return x.PassThroughMode
	}
	return nil
}

func (x *HealthCheck) GetCacheTime() *durationpb.Duration {
	if x != nil {
		return x.CacheTime
	}
	return nil
}

func (x *HealthCheck) GetClusterMinHealthyPercentages() map[string]*v3.Percent {
	if x != nil {
		return x.ClusterMinHealthyPercentages
	}
	return nil
}

func (x *HealthCheck) GetHeaders() []*v31.HeaderMatcher {
	if x != nil {
		return x.Headers
	}
	return nil
}

var File_envoy_extensions_filters_http_health_check_v3_health_check_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_health_check_v3_health_check_proto_rawDesc = []byte{
	0x0a, 0x40, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x76, 0x33, 0x2f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76,
	0x33, 0x1a, 0x2c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64,
	0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x04, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c,
	0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x50, 0x0a, 0x11, 0x70, 0x61, 0x73, 0x73, 0x5f,
	0x74, 0x68, 0x72, 0x6f, 0x75, 0x67, 0x68, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x54, 0x68,
	0x72, 0x6f, 0x75, 0x67, 0x68, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x63, 0x61, 0x63,
	0x68, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x63, 0x61, 0x63, 0x68, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0xa3, 0x01, 0x0a, 0x1f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6d, 0x69, 0x6e, 0x5f, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x5f, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x5c, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x68, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x76, 0x33, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4d, 0x69, 0x6e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x50, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x1c, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x73, 0x12, 0x3e, 0x0a, 0x07, 0x68, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x2e,
	0x76, 0x33, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x67, 0x0a, 0x21, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x4d, 0x69, 0x6e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x79, 0x50, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x33, 0x2e,
	0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x3a, 0x3b, 0x9a, 0xc5, 0x88, 0x1e, 0x36, 0x0a, 0x34, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x2e, 0x76, 0x32, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4a,
	0x04, 0x08, 0x02, 0x10, 0x03, 0x42, 0xbe, 0x01, 0x0a, 0x3b, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73,
	0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x2e, 0x76, 0x33, 0x42, 0x10, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x63, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2f, 0x76, 0x33, 0x3b,
	0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x76, 0x33, 0xba, 0x80,
	0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_health_check_v3_health_check_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_health_check_v3_health_check_proto_rawDescData = file_envoy_extensions_filters_http_health_check_v3_health_check_proto_rawDesc
)

func file_envoy_extensions_filters_http_health_check_v3_health_check_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_health_check_v3_health_check_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_health_check_v3_health_check_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_health_check_v3_health_check_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_health_check_v3_health_check_proto_rawDescData
}

var file_envoy_extensions_filters_http_health_check_v3_health_check_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_http_health_check_v3_health_check_proto_goTypes = []interface{}{
	(*HealthCheck)(nil),          // 0: envoy.extensions.filters.http.health_check.v3.HealthCheck
	nil,                          // 1: envoy.extensions.filters.http.health_check.v3.HealthCheck.ClusterMinHealthyPercentagesEntry
	(*wrapperspb.BoolValue)(nil), // 2: google.protobuf.BoolValue
	(*durationpb.Duration)(nil),  // 3: google.protobuf.Duration
	(*v31.HeaderMatcher)(nil),    // 4: envoy.config.route.v3.HeaderMatcher
	(*v3.Percent)(nil),           // 5: envoy.type.v3.Percent
}
var file_envoy_extensions_filters_http_health_check_v3_health_check_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.http.health_check.v3.HealthCheck.pass_through_mode:type_name -> google.protobuf.BoolValue
	3, // 1: envoy.extensions.filters.http.health_check.v3.HealthCheck.cache_time:type_name -> google.protobuf.Duration
	1, // 2: envoy.extensions.filters.http.health_check.v3.HealthCheck.cluster_min_healthy_percentages:type_name -> envoy.extensions.filters.http.health_check.v3.HealthCheck.ClusterMinHealthyPercentagesEntry
	4, // 3: envoy.extensions.filters.http.health_check.v3.HealthCheck.headers:type_name -> envoy.config.route.v3.HeaderMatcher
	5, // 4: envoy.extensions.filters.http.health_check.v3.HealthCheck.ClusterMinHealthyPercentagesEntry.value:type_name -> envoy.type.v3.Percent
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_health_check_v3_health_check_proto_init() }
func file_envoy_extensions_filters_http_health_check_v3_health_check_proto_init() {
	if File_envoy_extensions_filters_http_health_check_v3_health_check_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_health_check_v3_health_check_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheck); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_extensions_filters_http_health_check_v3_health_check_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_health_check_v3_health_check_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_health_check_v3_health_check_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_health_check_v3_health_check_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_health_check_v3_health_check_proto = out.File
	file_envoy_extensions_filters_http_health_check_v3_health_check_proto_rawDesc = nil
	file_envoy_extensions_filters_http_health_check_v3_health_check_proto_goTypes = nil
	file_envoy_extensions_filters_http_health_check_v3_health_check_proto_depIdxs = nil
}
