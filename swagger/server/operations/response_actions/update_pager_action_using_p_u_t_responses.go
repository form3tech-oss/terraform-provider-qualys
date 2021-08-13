/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package response_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/form3tech-oss/terraform-provider-qualys/swagger/models"
)

// UpdatePagerActionUsingPUTOKCode is the HTTP code returned for type UpdatePagerActionUsingPUTOK
const UpdatePagerActionUsingPUTOKCode int = 200

/*UpdatePagerActionUsingPUTOK OK

swagger:response updatePagerActionUsingPUTOK
*/
type UpdatePagerActionUsingPUTOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseEntity `json:"body,omitempty"`
}

// NewUpdatePagerActionUsingPUTOK creates UpdatePagerActionUsingPUTOK with default headers values
func NewUpdatePagerActionUsingPUTOK() *UpdatePagerActionUsingPUTOK {

	return &UpdatePagerActionUsingPUTOK{}
}

// WithPayload adds the payload to the update pager action using p u t o k response
func (o *UpdatePagerActionUsingPUTOK) WithPayload(payload *models.ResponseEntity) *UpdatePagerActionUsingPUTOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update pager action using p u t o k response
func (o *UpdatePagerActionUsingPUTOK) SetPayload(payload *models.ResponseEntity) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePagerActionUsingPUTOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePagerActionUsingPUTCreatedCode is the HTTP code returned for type UpdatePagerActionUsingPUTCreated
const UpdatePagerActionUsingPUTCreatedCode int = 201

/*UpdatePagerActionUsingPUTCreated Created

swagger:response updatePagerActionUsingPUTCreated
*/
type UpdatePagerActionUsingPUTCreated struct {
}

// NewUpdatePagerActionUsingPUTCreated creates UpdatePagerActionUsingPUTCreated with default headers values
func NewUpdatePagerActionUsingPUTCreated() *UpdatePagerActionUsingPUTCreated {

	return &UpdatePagerActionUsingPUTCreated{}
}

// WriteResponse to the client
func (o *UpdatePagerActionUsingPUTCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// UpdatePagerActionUsingPUTBadRequestCode is the HTTP code returned for type UpdatePagerActionUsingPUTBadRequest
const UpdatePagerActionUsingPUTBadRequestCode int = 400

/*UpdatePagerActionUsingPUTBadRequest Bad Request

swagger:response updatePagerActionUsingPUTBadRequest
*/
type UpdatePagerActionUsingPUTBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIErrorFacade `json:"body,omitempty"`
}

// NewUpdatePagerActionUsingPUTBadRequest creates UpdatePagerActionUsingPUTBadRequest with default headers values
func NewUpdatePagerActionUsingPUTBadRequest() *UpdatePagerActionUsingPUTBadRequest {

	return &UpdatePagerActionUsingPUTBadRequest{}
}

// WithPayload adds the payload to the update pager action using p u t bad request response
func (o *UpdatePagerActionUsingPUTBadRequest) WithPayload(payload *models.APIErrorFacade) *UpdatePagerActionUsingPUTBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update pager action using p u t bad request response
func (o *UpdatePagerActionUsingPUTBadRequest) SetPayload(payload *models.APIErrorFacade) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePagerActionUsingPUTBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePagerActionUsingPUTUnauthorizedCode is the HTTP code returned for type UpdatePagerActionUsingPUTUnauthorized
const UpdatePagerActionUsingPUTUnauthorizedCode int = 401

/*UpdatePagerActionUsingPUTUnauthorized Unauthorized user

swagger:response updatePagerActionUsingPUTUnauthorized
*/
type UpdatePagerActionUsingPUTUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.APIErrorFacade `json:"body,omitempty"`
}

// NewUpdatePagerActionUsingPUTUnauthorized creates UpdatePagerActionUsingPUTUnauthorized with default headers values
func NewUpdatePagerActionUsingPUTUnauthorized() *UpdatePagerActionUsingPUTUnauthorized {

	return &UpdatePagerActionUsingPUTUnauthorized{}
}

// WithPayload adds the payload to the update pager action using p u t unauthorized response
func (o *UpdatePagerActionUsingPUTUnauthorized) WithPayload(payload *models.APIErrorFacade) *UpdatePagerActionUsingPUTUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update pager action using p u t unauthorized response
func (o *UpdatePagerActionUsingPUTUnauthorized) SetPayload(payload *models.APIErrorFacade) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdatePagerActionUsingPUTUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// UpdatePagerActionUsingPUTForbiddenCode is the HTTP code returned for type UpdatePagerActionUsingPUTForbidden
const UpdatePagerActionUsingPUTForbiddenCode int = 403

/*UpdatePagerActionUsingPUTForbidden Forbidden

swagger:response updatePagerActionUsingPUTForbidden
*/
type UpdatePagerActionUsingPUTForbidden struct {
}

// NewUpdatePagerActionUsingPUTForbidden creates UpdatePagerActionUsingPUTForbidden with default headers values
func NewUpdatePagerActionUsingPUTForbidden() *UpdatePagerActionUsingPUTForbidden {

	return &UpdatePagerActionUsingPUTForbidden{}
}

// WriteResponse to the client
func (o *UpdatePagerActionUsingPUTForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// UpdatePagerActionUsingPUTNotFoundCode is the HTTP code returned for type UpdatePagerActionUsingPUTNotFound
const UpdatePagerActionUsingPUTNotFoundCode int = 404

/*UpdatePagerActionUsingPUTNotFound Not Found

swagger:response updatePagerActionUsingPUTNotFound
*/
type UpdatePagerActionUsingPUTNotFound struct {
}

// NewUpdatePagerActionUsingPUTNotFound creates UpdatePagerActionUsingPUTNotFound with default headers values
func NewUpdatePagerActionUsingPUTNotFound() *UpdatePagerActionUsingPUTNotFound {

	return &UpdatePagerActionUsingPUTNotFound{}
}

// WriteResponse to the client
func (o *UpdatePagerActionUsingPUTNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
