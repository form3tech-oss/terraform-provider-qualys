/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package remediation_activity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/form3tech-oss/terraform-provider-qualys/swagger/models"
)

// GetRemediationHistoryUsingGETOKCode is the HTTP code returned for type GetRemediationHistoryUsingGETOK
const GetRemediationHistoryUsingGETOKCode int = 200

/*GetRemediationHistoryUsingGETOK OK

swagger:response getRemediationHistoryUsingGETOK
*/
type GetRemediationHistoryUsingGETOK struct {

	/*
	  In: Body
	*/
	Payload *models.PageImplRemediationActivityResponse `json:"body,omitempty"`
}

// NewGetRemediationHistoryUsingGETOK creates GetRemediationHistoryUsingGETOK with default headers values
func NewGetRemediationHistoryUsingGETOK() *GetRemediationHistoryUsingGETOK {

	return &GetRemediationHistoryUsingGETOK{}
}

// WithPayload adds the payload to the get remediation history using g e t o k response
func (o *GetRemediationHistoryUsingGETOK) WithPayload(payload *models.PageImplRemediationActivityResponse) *GetRemediationHistoryUsingGETOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get remediation history using g e t o k response
func (o *GetRemediationHistoryUsingGETOK) SetPayload(payload *models.PageImplRemediationActivityResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRemediationHistoryUsingGETOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRemediationHistoryUsingGETBadRequestCode is the HTTP code returned for type GetRemediationHistoryUsingGETBadRequest
const GetRemediationHistoryUsingGETBadRequestCode int = 400

/*GetRemediationHistoryUsingGETBadRequest Bad Request

swagger:response getRemediationHistoryUsingGETBadRequest
*/
type GetRemediationHistoryUsingGETBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.APIErrorFacade `json:"body,omitempty"`
}

// NewGetRemediationHistoryUsingGETBadRequest creates GetRemediationHistoryUsingGETBadRequest with default headers values
func NewGetRemediationHistoryUsingGETBadRequest() *GetRemediationHistoryUsingGETBadRequest {

	return &GetRemediationHistoryUsingGETBadRequest{}
}

// WithPayload adds the payload to the get remediation history using g e t bad request response
func (o *GetRemediationHistoryUsingGETBadRequest) WithPayload(payload *models.APIErrorFacade) *GetRemediationHistoryUsingGETBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get remediation history using g e t bad request response
func (o *GetRemediationHistoryUsingGETBadRequest) SetPayload(payload *models.APIErrorFacade) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRemediationHistoryUsingGETBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRemediationHistoryUsingGETUnauthorizedCode is the HTTP code returned for type GetRemediationHistoryUsingGETUnauthorized
const GetRemediationHistoryUsingGETUnauthorizedCode int = 401

/*GetRemediationHistoryUsingGETUnauthorized Unauthorized user

swagger:response getRemediationHistoryUsingGETUnauthorized
*/
type GetRemediationHistoryUsingGETUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.APIErrorFacade `json:"body,omitempty"`
}

// NewGetRemediationHistoryUsingGETUnauthorized creates GetRemediationHistoryUsingGETUnauthorized with default headers values
func NewGetRemediationHistoryUsingGETUnauthorized() *GetRemediationHistoryUsingGETUnauthorized {

	return &GetRemediationHistoryUsingGETUnauthorized{}
}

// WithPayload adds the payload to the get remediation history using g e t unauthorized response
func (o *GetRemediationHistoryUsingGETUnauthorized) WithPayload(payload *models.APIErrorFacade) *GetRemediationHistoryUsingGETUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get remediation history using g e t unauthorized response
func (o *GetRemediationHistoryUsingGETUnauthorized) SetPayload(payload *models.APIErrorFacade) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRemediationHistoryUsingGETUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRemediationHistoryUsingGETForbiddenCode is the HTTP code returned for type GetRemediationHistoryUsingGETForbidden
const GetRemediationHistoryUsingGETForbiddenCode int = 403

/*GetRemediationHistoryUsingGETForbidden Forbidden

swagger:response getRemediationHistoryUsingGETForbidden
*/
type GetRemediationHistoryUsingGETForbidden struct {
}

// NewGetRemediationHistoryUsingGETForbidden creates GetRemediationHistoryUsingGETForbidden with default headers values
func NewGetRemediationHistoryUsingGETForbidden() *GetRemediationHistoryUsingGETForbidden {

	return &GetRemediationHistoryUsingGETForbidden{}
}

// WriteResponse to the client
func (o *GetRemediationHistoryUsingGETForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// GetRemediationHistoryUsingGETNotFoundCode is the HTTP code returned for type GetRemediationHistoryUsingGETNotFound
const GetRemediationHistoryUsingGETNotFoundCode int = 404

/*GetRemediationHistoryUsingGETNotFound Not Found

swagger:response getRemediationHistoryUsingGETNotFound
*/
type GetRemediationHistoryUsingGETNotFound struct {
}

// NewGetRemediationHistoryUsingGETNotFound creates GetRemediationHistoryUsingGETNotFound with default headers values
func NewGetRemediationHistoryUsingGETNotFound() *GetRemediationHistoryUsingGETNotFound {

	return &GetRemediationHistoryUsingGETNotFound{}
}

// WriteResponse to the client
func (o *GetRemediationHistoryUsingGETNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
