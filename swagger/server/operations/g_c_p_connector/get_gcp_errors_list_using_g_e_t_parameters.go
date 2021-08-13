/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package g_c_p_connector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NewGetGcpErrorsListUsingGETParams creates a new GetGcpErrorsListUsingGETParams object
// with the default values initialized.
func NewGetGcpErrorsListUsingGETParams() GetGcpErrorsListUsingGETParams {

	var (
		// initialize parameters with default values

		pageNoDefault   = int32(0)
		pageSizeDefault = int32(50)
	)

	return GetGcpErrorsListUsingGETParams{
		PageNo: pageNoDefault,

		PageSize: pageSizeDefault,
	}
}

// GetGcpErrorsListUsingGETParams contains all the bound params for the get gcp errors list using g e t operation
// typically these are obtained from a http.Request
//
// swagger:parameters getGcpErrorsListUsingGET
type GetGcpErrorsListUsingGETParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Specify the connector ID of a specific GCP connector in the user’s scope.
	  Required: true
	  In: path
	*/
	ConnectorID string
	/*The page to be returned.
	  Required: true
	  In: query
	  Default: 0
	*/
	PageNo int32
	/*The number of records per page to be included in the response.
	  Required: true
	  In: query
	  Default: 50
	*/
	PageSize int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetGcpErrorsListUsingGETParams() beforehand.
func (o *GetGcpErrorsListUsingGETParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rConnectorID, rhkConnectorID, _ := route.Params.GetOK("connectorId")
	if err := o.bindConnectorID(rConnectorID, rhkConnectorID, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageNo, qhkPageNo, _ := qs.GetOK("pageNo")
	if err := o.bindPageNo(qPageNo, qhkPageNo, route.Formats); err != nil {
		res = append(res, err)
	}

	qPageSize, qhkPageSize, _ := qs.GetOK("pageSize")
	if err := o.bindPageSize(qPageSize, qhkPageSize, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindConnectorID binds and validates parameter ConnectorID from path.
func (o *GetGcpErrorsListUsingGETParams) bindConnectorID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.ConnectorID = raw

	return nil
}

// bindPageNo binds and validates parameter PageNo from query.
func (o *GetGcpErrorsListUsingGETParams) bindPageNo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("pageNo", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("pageNo", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("pageNo", "query", "int32", raw)
	}
	o.PageNo = value

	return nil
}

// bindPageSize binds and validates parameter PageSize from query.
func (o *GetGcpErrorsListUsingGETParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("pageSize", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("pageSize", "query", raw); err != nil {
		return err
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("pageSize", "query", "int32", raw)
	}
	o.PageSize = value

	return nil
}
