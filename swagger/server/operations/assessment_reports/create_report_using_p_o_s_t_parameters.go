/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package assessment_reports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/form3tech-oss/terraform-provider-qualys/swagger/models"
)

// NewCreateReportUsingPOSTParams creates a new CreateReportUsingPOSTParams object
// no default values defined in spec.
func NewCreateReportUsingPOSTParams() CreateReportUsingPOSTParams {

	return CreateReportUsingPOSTParams{}
}

// CreateReportUsingPOSTParams contains all the bound params for the create report using p o s t operation
// typically these are obtained from a http.Request
//
// swagger:parameters createReportUsingPOST
type CreateReportUsingPOSTParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Create Report Request
	  Required: true
	  In: body
	*/
	CreateReportRequest *models.CreateReportRequestAPI
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewCreateReportUsingPOSTParams() beforehand.
func (o *CreateReportUsingPOSTParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.CreateReportRequestAPI
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("createReportRequest", "body", ""))
			} else {
				res = append(res, errors.NewParseError("createReportRequest", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.CreateReportRequest = &body
			}
		}
	} else {
		res = append(res, errors.Required("createReportRequest", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
