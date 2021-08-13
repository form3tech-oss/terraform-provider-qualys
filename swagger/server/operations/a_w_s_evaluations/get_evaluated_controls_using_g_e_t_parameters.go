/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package a_w_s_evaluations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewGetEvaluatedControlsUsingGETParams creates a new GetEvaluatedControlsUsingGETParams object
// with the default values initialized.
func NewGetEvaluatedControlsUsingGETParams() GetEvaluatedControlsUsingGETParams {

	var (
		// initialize parameters with default values

		pageNoDefault   = int32(0)
		pageSizeDefault = int32(300)
	)

	return GetEvaluatedControlsUsingGETParams{
		PageNo: &pageNoDefault,

		PageSize: &pageSizeDefault,
	}
}

// GetEvaluatedControlsUsingGETParams contains all the bound params for the get evaluated controls using g e t operation
// typically these are obtained from a http.Request
//
// swagger:parameters getEvaluatedControlsUsingGET
type GetEvaluatedControlsUsingGETParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Specify the ID of a specific account in the user’s scope.
	  Required: true
	  In: path
	*/
	AccountID string
	/*Filter the resources list by providing a query using Qualys syntax. <a href="/cloudview/help/search_tips/search_ui_evaluations.htm" target="_blank">Click here</a> for help with creating your query.
	  In: query
	*/
	Filter *string
	/*The page to be returned.
	  In: query
	  Default: 0
	*/
	PageNo *int32
	/*The number of records per page to be included in the response.
	  In: query
	  Default: 300
	*/
	PageSize *int32
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetEvaluatedControlsUsingGETParams() beforehand.
func (o *GetEvaluatedControlsUsingGETParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rAccountID, rhkAccountID, _ := route.Params.GetOK("accountId")
	if err := o.bindAccountID(rAccountID, rhkAccountID, route.Formats); err != nil {
		res = append(res, err)
	}

	qFilter, qhkFilter, _ := qs.GetOK("filter")
	if err := o.bindFilter(qFilter, qhkFilter, route.Formats); err != nil {
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

// bindAccountID binds and validates parameter AccountID from path.
func (o *GetEvaluatedControlsUsingGETParams) bindAccountID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.AccountID = raw

	return nil
}

// bindFilter binds and validates parameter Filter from query.
func (o *GetEvaluatedControlsUsingGETParams) bindFilter(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Filter = &raw

	return nil
}

// bindPageNo binds and validates parameter PageNo from query.
func (o *GetEvaluatedControlsUsingGETParams) bindPageNo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetEvaluatedControlsUsingGETParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("pageNo", "query", "int32", raw)
	}
	o.PageNo = &value

	return nil
}

// bindPageSize binds and validates parameter PageSize from query.
func (o *GetEvaluatedControlsUsingGETParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewGetEvaluatedControlsUsingGETParams()
		return nil
	}

	value, err := swag.ConvertInt32(raw)
	if err != nil {
		return errors.InvalidType("pageSize", "query", "int32", raw)
	}
	o.PageSize = &value

	return nil
}
