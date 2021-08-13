/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package azure_connector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

// NewDeleteConnectorsUsingDELETEParams creates a new DeleteConnectorsUsingDELETEParams object
// no default values defined in spec.
func NewDeleteConnectorsUsingDELETEParams() DeleteConnectorsUsingDELETEParams {

	return DeleteConnectorsUsingDELETEParams{}
}

// DeleteConnectorsUsingDELETEParams contains all the bound params for the delete connectors using d e l e t e operation
// typically these are obtained from a http.Request
//
// swagger:parameters deleteConnectorsUsingDELETE
type DeleteConnectorsUsingDELETEParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Specify the IDs of the connectors to be deleted.
	  Required: true
	  In: body
	*/
	ConnectorIds []string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewDeleteConnectorsUsingDELETEParams() beforehand.
func (o *DeleteConnectorsUsingDELETEParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body []string
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("connectorIds", "body", ""))
			} else {
				res = append(res, errors.NewParseError("connectorIds", "body", "", err))
			}
		} else {
			// no validation required on inline body
			o.ConnectorIds = body
		}
	} else {
		res = append(res, errors.Required("connectorIds", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
