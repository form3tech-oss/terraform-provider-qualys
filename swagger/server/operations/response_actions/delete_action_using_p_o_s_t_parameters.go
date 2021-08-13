/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package response_actions

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

// NewDeleteActionUsingPOSTParams creates a new DeleteActionUsingPOSTParams object
// no default values defined in spec.
func NewDeleteActionUsingPOSTParams() DeleteActionUsingPOSTParams {

	return DeleteActionUsingPOSTParams{}
}

// DeleteActionUsingPOSTParams contains all the bound params for the delete action using p o s t operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteActionUsingPOST
type DeleteActionUsingPOSTParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Specify the IDs of the actions to be deleted.
	  Required: true
	  In: body
	*/
	ActionIds *models.BulkIDParam
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteActionUsingPOSTParams() beforehand.
func (o *DeleteActionUsingPOSTParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.BulkIDParam
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("actionIds", "body", ""))
			} else {
				res = append(res, errors.NewParseError("actionIds", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.ActionIds = &body
			}
		}
	} else {
		res = append(res, errors.Required("actionIds", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
