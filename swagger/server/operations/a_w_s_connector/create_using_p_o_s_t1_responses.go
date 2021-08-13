/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package a_w_s_connector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/form3tech-oss/terraform-provider-qualys/swagger/models"
)

// CreateUsingPOST1OKCode is the HTTP code returned for type CreateUsingPOST1OK
const CreateUsingPOST1OKCode int = 200

/*CreateUsingPOST1OK OK

swagger:response createUsingPOST1OK
*/
type CreateUsingPOST1OK struct {

	/*
	  In: Body
	*/
	Payload *models.AwsConnectorResponse `json:"body,omitempty"`
}

// NewCreateUsingPOST1OK creates CreateUsingPOST1OK with default headers values
func NewCreateUsingPOST1OK() *CreateUsingPOST1OK {

	return &CreateUsingPOST1OK{}
}

// WithPayload adds the payload to the create using p o s t1 o k response
func (o *CreateUsingPOST1OK) WithPayload(payload *models.AwsConnectorResponse) *CreateUsingPOST1OK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create using p o s t1 o k response
func (o *CreateUsingPOST1OK) SetPayload(payload *models.AwsConnectorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUsingPOST1OK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUsingPOST1CreatedCode is the HTTP code returned for type CreateUsingPOST1Created
const CreateUsingPOST1CreatedCode int = 201

/*CreateUsingPOST1Created Created

swagger:response createUsingPOST1Created
*/
type CreateUsingPOST1Created struct {
}

// NewCreateUsingPOST1Created creates CreateUsingPOST1Created with default headers values
func NewCreateUsingPOST1Created() *CreateUsingPOST1Created {

	return &CreateUsingPOST1Created{}
}

// WriteResponse to the client
func (o *CreateUsingPOST1Created) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// CreateUsingPOST1BadRequestCode is the HTTP code returned for type CreateUsingPOST1BadRequest
const CreateUsingPOST1BadRequestCode int = 400

/*CreateUsingPOST1BadRequest Bad Request

swagger:response createUsingPOST1BadRequest
*/
type CreateUsingPOST1BadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIErrorFacade `json:"body,omitempty"`
}

// NewCreateUsingPOST1BadRequest creates CreateUsingPOST1BadRequest with default headers values
func NewCreateUsingPOST1BadRequest() *CreateUsingPOST1BadRequest {

	return &CreateUsingPOST1BadRequest{}
}

// WithPayload adds the payload to the create using p o s t1 bad request response
func (o *CreateUsingPOST1BadRequest) WithPayload(payload *models.APIErrorFacade) *CreateUsingPOST1BadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create using p o s t1 bad request response
func (o *CreateUsingPOST1BadRequest) SetPayload(payload *models.APIErrorFacade) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUsingPOST1BadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUsingPOST1UnauthorizedCode is the HTTP code returned for type CreateUsingPOST1Unauthorized
const CreateUsingPOST1UnauthorizedCode int = 401

/*CreateUsingPOST1Unauthorized Unauthorized user

swagger:response createUsingPOST1Unauthorized
*/
type CreateUsingPOST1Unauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.APIErrorFacade `json:"body,omitempty"`
}

// NewCreateUsingPOST1Unauthorized creates CreateUsingPOST1Unauthorized with default headers values
func NewCreateUsingPOST1Unauthorized() *CreateUsingPOST1Unauthorized {

	return &CreateUsingPOST1Unauthorized{}
}

// WithPayload adds the payload to the create using p o s t1 unauthorized response
func (o *CreateUsingPOST1Unauthorized) WithPayload(payload *models.APIErrorFacade) *CreateUsingPOST1Unauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create using p o s t1 unauthorized response
func (o *CreateUsingPOST1Unauthorized) SetPayload(payload *models.APIErrorFacade) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateUsingPOST1Unauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// CreateUsingPOST1ForbiddenCode is the HTTP code returned for type CreateUsingPOST1Forbidden
const CreateUsingPOST1ForbiddenCode int = 403

/*CreateUsingPOST1Forbidden Forbidden

swagger:response createUsingPOST1Forbidden
*/
type CreateUsingPOST1Forbidden struct {
}

// NewCreateUsingPOST1Forbidden creates CreateUsingPOST1Forbidden with default headers values
func NewCreateUsingPOST1Forbidden() *CreateUsingPOST1Forbidden {

	return &CreateUsingPOST1Forbidden{}
}

// WriteResponse to the client
func (o *CreateUsingPOST1Forbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// CreateUsingPOST1NotFoundCode is the HTTP code returned for type CreateUsingPOST1NotFound
const CreateUsingPOST1NotFoundCode int = 404

/*CreateUsingPOST1NotFound Not Found

swagger:response createUsingPOST1NotFound
*/
type CreateUsingPOST1NotFound struct {
}

// NewCreateUsingPOST1NotFound creates CreateUsingPOST1NotFound with default headers values
func NewCreateUsingPOST1NotFound() *CreateUsingPOST1NotFound {

	return &CreateUsingPOST1NotFound{}
}

// WriteResponse to the client
func (o *CreateUsingPOST1NotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
