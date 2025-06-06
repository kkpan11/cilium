// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/cilium/cilium/api/v1/models"
)

// PutEndpointIDReader is a Reader for the PutEndpointID structure.
type PutEndpointIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutEndpointIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPutEndpointIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutEndpointIDInvalid()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutEndpointIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutEndpointIDExists()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutEndpointIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutEndpointIDFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /endpoint/{id}] PutEndpointID", response, response.Code())
	}
}

// NewPutEndpointIDCreated creates a PutEndpointIDCreated with default headers values
func NewPutEndpointIDCreated() *PutEndpointIDCreated {
	return &PutEndpointIDCreated{}
}

/*
PutEndpointIDCreated describes a response with status code 201, with default header values.

Created
*/
type PutEndpointIDCreated struct {
	Payload *models.Endpoint
}

// IsSuccess returns true when this put endpoint Id created response has a 2xx status code
func (o *PutEndpointIDCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put endpoint Id created response has a 3xx status code
func (o *PutEndpointIDCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put endpoint Id created response has a 4xx status code
func (o *PutEndpointIDCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this put endpoint Id created response has a 5xx status code
func (o *PutEndpointIDCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this put endpoint Id created response a status code equal to that given
func (o *PutEndpointIDCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the put endpoint Id created response
func (o *PutEndpointIDCreated) Code() int {
	return 201
}

func (o *PutEndpointIDCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdCreated %s", 201, payload)
}

func (o *PutEndpointIDCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdCreated %s", 201, payload)
}

func (o *PutEndpointIDCreated) GetPayload() *models.Endpoint {
	return o.Payload
}

func (o *PutEndpointIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Endpoint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutEndpointIDInvalid creates a PutEndpointIDInvalid with default headers values
func NewPutEndpointIDInvalid() *PutEndpointIDInvalid {
	return &PutEndpointIDInvalid{}
}

/*
PutEndpointIDInvalid describes a response with status code 400, with default header values.

Invalid endpoint in request
*/
type PutEndpointIDInvalid struct {
	Payload models.Error
}

// IsSuccess returns true when this put endpoint Id invalid response has a 2xx status code
func (o *PutEndpointIDInvalid) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put endpoint Id invalid response has a 3xx status code
func (o *PutEndpointIDInvalid) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put endpoint Id invalid response has a 4xx status code
func (o *PutEndpointIDInvalid) IsClientError() bool {
	return true
}

// IsServerError returns true when this put endpoint Id invalid response has a 5xx status code
func (o *PutEndpointIDInvalid) IsServerError() bool {
	return false
}

// IsCode returns true when this put endpoint Id invalid response a status code equal to that given
func (o *PutEndpointIDInvalid) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the put endpoint Id invalid response
func (o *PutEndpointIDInvalid) Code() int {
	return 400
}

func (o *PutEndpointIDInvalid) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdInvalid %s", 400, payload)
}

func (o *PutEndpointIDInvalid) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdInvalid %s", 400, payload)
}

func (o *PutEndpointIDInvalid) GetPayload() models.Error {
	return o.Payload
}

func (o *PutEndpointIDInvalid) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutEndpointIDForbidden creates a PutEndpointIDForbidden with default headers values
func NewPutEndpointIDForbidden() *PutEndpointIDForbidden {
	return &PutEndpointIDForbidden{}
}

/*
PutEndpointIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PutEndpointIDForbidden struct {
}

// IsSuccess returns true when this put endpoint Id forbidden response has a 2xx status code
func (o *PutEndpointIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put endpoint Id forbidden response has a 3xx status code
func (o *PutEndpointIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put endpoint Id forbidden response has a 4xx status code
func (o *PutEndpointIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put endpoint Id forbidden response has a 5xx status code
func (o *PutEndpointIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put endpoint Id forbidden response a status code equal to that given
func (o *PutEndpointIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the put endpoint Id forbidden response
func (o *PutEndpointIDForbidden) Code() int {
	return 403
}

func (o *PutEndpointIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdForbidden", 403)
}

func (o *PutEndpointIDForbidden) String() string {
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdForbidden", 403)
}

func (o *PutEndpointIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutEndpointIDExists creates a PutEndpointIDExists with default headers values
func NewPutEndpointIDExists() *PutEndpointIDExists {
	return &PutEndpointIDExists{}
}

/*
PutEndpointIDExists describes a response with status code 409, with default header values.

Endpoint already exists
*/
type PutEndpointIDExists struct {
}

// IsSuccess returns true when this put endpoint Id exists response has a 2xx status code
func (o *PutEndpointIDExists) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put endpoint Id exists response has a 3xx status code
func (o *PutEndpointIDExists) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put endpoint Id exists response has a 4xx status code
func (o *PutEndpointIDExists) IsClientError() bool {
	return true
}

// IsServerError returns true when this put endpoint Id exists response has a 5xx status code
func (o *PutEndpointIDExists) IsServerError() bool {
	return false
}

// IsCode returns true when this put endpoint Id exists response a status code equal to that given
func (o *PutEndpointIDExists) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the put endpoint Id exists response
func (o *PutEndpointIDExists) Code() int {
	return 409
}

func (o *PutEndpointIDExists) Error() string {
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdExists", 409)
}

func (o *PutEndpointIDExists) String() string {
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdExists", 409)
}

func (o *PutEndpointIDExists) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutEndpointIDTooManyRequests creates a PutEndpointIDTooManyRequests with default headers values
func NewPutEndpointIDTooManyRequests() *PutEndpointIDTooManyRequests {
	return &PutEndpointIDTooManyRequests{}
}

/*
PutEndpointIDTooManyRequests describes a response with status code 429, with default header values.

Rate-limiting too many requests in the given time frame
*/
type PutEndpointIDTooManyRequests struct {
}

// IsSuccess returns true when this put endpoint Id too many requests response has a 2xx status code
func (o *PutEndpointIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put endpoint Id too many requests response has a 3xx status code
func (o *PutEndpointIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put endpoint Id too many requests response has a 4xx status code
func (o *PutEndpointIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put endpoint Id too many requests response has a 5xx status code
func (o *PutEndpointIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put endpoint Id too many requests response a status code equal to that given
func (o *PutEndpointIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the put endpoint Id too many requests response
func (o *PutEndpointIDTooManyRequests) Code() int {
	return 429
}

func (o *PutEndpointIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdTooManyRequests", 429)
}

func (o *PutEndpointIDTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdTooManyRequests", 429)
}

func (o *PutEndpointIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutEndpointIDFailed creates a PutEndpointIDFailed with default headers values
func NewPutEndpointIDFailed() *PutEndpointIDFailed {
	return &PutEndpointIDFailed{}
}

/*
PutEndpointIDFailed describes a response with status code 500, with default header values.

Endpoint creation failed
*/
type PutEndpointIDFailed struct {
	Payload models.Error
}

// IsSuccess returns true when this put endpoint Id failed response has a 2xx status code
func (o *PutEndpointIDFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put endpoint Id failed response has a 3xx status code
func (o *PutEndpointIDFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put endpoint Id failed response has a 4xx status code
func (o *PutEndpointIDFailed) IsClientError() bool {
	return false
}

// IsServerError returns true when this put endpoint Id failed response has a 5xx status code
func (o *PutEndpointIDFailed) IsServerError() bool {
	return true
}

// IsCode returns true when this put endpoint Id failed response a status code equal to that given
func (o *PutEndpointIDFailed) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the put endpoint Id failed response
func (o *PutEndpointIDFailed) Code() int {
	return 500
}

func (o *PutEndpointIDFailed) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdFailed %s", 500, payload)
}

func (o *PutEndpointIDFailed) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /endpoint/{id}][%d] putEndpointIdFailed %s", 500, payload)
}

func (o *PutEndpointIDFailed) GetPayload() models.Error {
	return o.Payload
}

func (o *PutEndpointIDFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
