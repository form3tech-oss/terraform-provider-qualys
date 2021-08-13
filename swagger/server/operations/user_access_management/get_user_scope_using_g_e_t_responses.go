/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package user_access_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/form3tech-oss/terraform-provider-qualys/swagger/models"
)

// GetUserScopeUsingGETOKCode is the HTTP code returned for type GetUserScopeUsingGETOK
const GetUserScopeUsingGETOKCode int = 200

/*GetUserScopeUsingGETOK OK

swagger:response getUserScopeUsingGETOK
*/
type GetUserScopeUsingGETOK struct {

	/*
	  In: Body
	*/
	Payload map[string]interface{} `json:"body,omitempty"`
}

// NewGetUserScopeUsingGETOK creates GetUserScopeUsingGETOK with default headers values
func NewGetUserScopeUsingGETOK() *GetUserScopeUsingGETOK {

	return &GetUserScopeUsingGETOK{}
}

// WithPayload adds the payload to the get user scope using g e t o k response
func (o *GetUserScopeUsingGETOK) WithPayload(payload map[string]interface{}) *GetUserScopeUsingGETOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user scope using g e t o k response
func (o *GetUserScopeUsingGETOK) SetPayload(payload map[string]interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserScopeUsingGETOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty map
		payload = make(map[string]interface{}, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// GetUserScopeUsingGETBadRequestCode is the HTTP code returned for type GetUserScopeUsingGETBadRequest
const GetUserScopeUsingGETBadRequestCode int = 400

/*GetUserScopeUsingGETBadRequest Bad Request

swagger:response getUserScopeUsingGETBadRequest
*/
type GetUserScopeUsingGETBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIErrorFacade `json:"body,omitempty"`
}

// NewGetUserScopeUsingGETBadRequest creates GetUserScopeUsingGETBadRequest with default headers values
func NewGetUserScopeUsingGETBadRequest() *GetUserScopeUsingGETBadRequest {

	return &GetUserScopeUsingGETBadRequest{}
}

// WithPayload adds the payload to the get user scope using g e t bad request response
func (o *GetUserScopeUsingGETBadRequest) WithPayload(payload *models.APIErrorFacade) *GetUserScopeUsingGETBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user scope using g e t bad request response
func (o *GetUserScopeUsingGETBadRequest) SetPayload(payload *models.APIErrorFacade) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserScopeUsingGETBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserScopeUsingGETUnauthorizedCode is the HTTP code returned for type GetUserScopeUsingGETUnauthorized
const GetUserScopeUsingGETUnauthorizedCode int = 401

/*GetUserScopeUsingGETUnauthorized Unauthorized user

swagger:response getUserScopeUsingGETUnauthorized
*/
type GetUserScopeUsingGETUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.APIErrorFacade `json:"body,omitempty"`
}

// NewGetUserScopeUsingGETUnauthorized creates GetUserScopeUsingGETUnauthorized with default headers values
func NewGetUserScopeUsingGETUnauthorized() *GetUserScopeUsingGETUnauthorized {

	return &GetUserScopeUsingGETUnauthorized{}
}

// WithPayload adds the payload to the get user scope using g e t unauthorized response
func (o *GetUserScopeUsingGETUnauthorized) WithPayload(payload *models.APIErrorFacade) *GetUserScopeUsingGETUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get user scope using g e t unauthorized response
func (o *GetUserScopeUsingGETUnauthorized) SetPayload(payload *models.APIErrorFacade) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetUserScopeUsingGETUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetUserScopeUsingGETForbiddenCode is the HTTP code returned for type GetUserScopeUsingGETForbidden
const GetUserScopeUsingGETForbiddenCode int = 403

/*GetUserScopeUsingGETForbidden Forbidden

swagger:response getUserScopeUsingGETForbidden
*/
type GetUserScopeUsingGETForbidden struct {
}

// NewGetUserScopeUsingGETForbidden creates GetUserScopeUsingGETForbidden with default headers values
func NewGetUserScopeUsingGETForbidden() *GetUserScopeUsingGETForbidden {

	return &GetUserScopeUsingGETForbidden{}
}

// WriteResponse to the client
func (o *GetUserScopeUsingGETForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetUserScopeUsingGETNotFoundCode is the HTTP code returned for type GetUserScopeUsingGETNotFound
const GetUserScopeUsingGETNotFoundCode int = 404

/*GetUserScopeUsingGETNotFound Not Found

swagger:response getUserScopeUsingGETNotFound
*/
type GetUserScopeUsingGETNotFound struct {
}

// NewGetUserScopeUsingGETNotFound creates GetUserScopeUsingGETNotFound with default headers values
func NewGetUserScopeUsingGETNotFound() *GetUserScopeUsingGETNotFound {

	return &GetUserScopeUsingGETNotFound{}
}

// WriteResponse to the client
func (o *GetUserScopeUsingGETNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
