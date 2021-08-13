/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package a_w_s_evaluations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/form3tech-oss/terraform-provider-qualys/swagger/models"
)

// GetEvaluatedControlsUsingGETOKCode is the HTTP code returned for type GetEvaluatedControlsUsingGETOK
const GetEvaluatedControlsUsingGETOKCode int = 200

/*GetEvaluatedControlsUsingGETOK OK

swagger:response getEvaluatedControlsUsingGETOK
*/
type GetEvaluatedControlsUsingGETOK struct {

	/*
	  In: Body
	*/
	Payload *models.PageImplAwsEvaluationSummaryListResponse `json:"body,omitempty"`
}

// NewGetEvaluatedControlsUsingGETOK creates GetEvaluatedControlsUsingGETOK with default headers values
func NewGetEvaluatedControlsUsingGETOK() *GetEvaluatedControlsUsingGETOK {

	return &GetEvaluatedControlsUsingGETOK{}
}

// WithPayload adds the payload to the get evaluated controls using g e t o k response
func (o *GetEvaluatedControlsUsingGETOK) WithPayload(payload *models.PageImplAwsEvaluationSummaryListResponse) *GetEvaluatedControlsUsingGETOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get evaluated controls using g e t o k response
func (o *GetEvaluatedControlsUsingGETOK) SetPayload(payload *models.PageImplAwsEvaluationSummaryListResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEvaluatedControlsUsingGETOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEvaluatedControlsUsingGETBadRequestCode is the HTTP code returned for type GetEvaluatedControlsUsingGETBadRequest
const GetEvaluatedControlsUsingGETBadRequestCode int = 400

/*GetEvaluatedControlsUsingGETBadRequest Bad Request

swagger:response getEvaluatedControlsUsingGETBadRequest
*/
type GetEvaluatedControlsUsingGETBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIErrorFacade `json:"body,omitempty"`
}

// NewGetEvaluatedControlsUsingGETBadRequest creates GetEvaluatedControlsUsingGETBadRequest with default headers values
func NewGetEvaluatedControlsUsingGETBadRequest() *GetEvaluatedControlsUsingGETBadRequest {

	return &GetEvaluatedControlsUsingGETBadRequest{}
}

// WithPayload adds the payload to the get evaluated controls using g e t bad request response
func (o *GetEvaluatedControlsUsingGETBadRequest) WithPayload(payload *models.APIErrorFacade) *GetEvaluatedControlsUsingGETBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get evaluated controls using g e t bad request response
func (o *GetEvaluatedControlsUsingGETBadRequest) SetPayload(payload *models.APIErrorFacade) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEvaluatedControlsUsingGETBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEvaluatedControlsUsingGETUnauthorizedCode is the HTTP code returned for type GetEvaluatedControlsUsingGETUnauthorized
const GetEvaluatedControlsUsingGETUnauthorizedCode int = 401

/*GetEvaluatedControlsUsingGETUnauthorized Unauthorized user

swagger:response getEvaluatedControlsUsingGETUnauthorized
*/
type GetEvaluatedControlsUsingGETUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.APIErrorFacade `json:"body,omitempty"`
}

// NewGetEvaluatedControlsUsingGETUnauthorized creates GetEvaluatedControlsUsingGETUnauthorized with default headers values
func NewGetEvaluatedControlsUsingGETUnauthorized() *GetEvaluatedControlsUsingGETUnauthorized {

	return &GetEvaluatedControlsUsingGETUnauthorized{}
}

// WithPayload adds the payload to the get evaluated controls using g e t unauthorized response
func (o *GetEvaluatedControlsUsingGETUnauthorized) WithPayload(payload *models.APIErrorFacade) *GetEvaluatedControlsUsingGETUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get evaluated controls using g e t unauthorized response
func (o *GetEvaluatedControlsUsingGETUnauthorized) SetPayload(payload *models.APIErrorFacade) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEvaluatedControlsUsingGETUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetEvaluatedControlsUsingGETForbiddenCode is the HTTP code returned for type GetEvaluatedControlsUsingGETForbidden
const GetEvaluatedControlsUsingGETForbiddenCode int = 403

/*GetEvaluatedControlsUsingGETForbidden Forbidden

swagger:response getEvaluatedControlsUsingGETForbidden
*/
type GetEvaluatedControlsUsingGETForbidden struct {
}

// NewGetEvaluatedControlsUsingGETForbidden creates GetEvaluatedControlsUsingGETForbidden with default headers values
func NewGetEvaluatedControlsUsingGETForbidden() *GetEvaluatedControlsUsingGETForbidden {

	return &GetEvaluatedControlsUsingGETForbidden{}
}

// WriteResponse to the client
func (o *GetEvaluatedControlsUsingGETForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetEvaluatedControlsUsingGETNotFoundCode is the HTTP code returned for type GetEvaluatedControlsUsingGETNotFound
const GetEvaluatedControlsUsingGETNotFoundCode int = 404

/*GetEvaluatedControlsUsingGETNotFound Not Found

swagger:response getEvaluatedControlsUsingGETNotFound
*/
type GetEvaluatedControlsUsingGETNotFound struct {
}

// NewGetEvaluatedControlsUsingGETNotFound creates GetEvaluatedControlsUsingGETNotFound with default headers values
func NewGetEvaluatedControlsUsingGETNotFound() *GetEvaluatedControlsUsingGETNotFound {

	return &GetEvaluatedControlsUsingGETNotFound{}
}

// WriteResponse to the client
func (o *GetEvaluatedControlsUsingGETNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
