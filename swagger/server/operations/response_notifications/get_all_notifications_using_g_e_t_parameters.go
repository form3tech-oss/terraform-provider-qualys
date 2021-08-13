/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package response_notifications

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

// NewGetAllNotificationsUsingGETParams creates a new GetAllNotificationsUsingGETParams object
// with the default values initialized.
func NewGetAllNotificationsUsingGETParams() GetAllNotificationsUsingGETParams {

	var (
		// initialize parameters with default values

		pageNoDefault   = int32(0)
		pageSizeDefault = int32(50)
	)

	return GetAllNotificationsUsingGETParams{
		PageNo: pageNoDefault,

		PageSize: pageSizeDefault,
	}
}

// GetAllNotificationsUsingGETParams contains all the bound params for the get all notifications using g e t operation
// typically these are obtained from a http.Request
//
// swagger:parameters getAllNotificationsUsingGET
type GetAllNotificationsUsingGETParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Cloud Type
	  Required: true
	  In: query
	*/
	CloudType string
	/*Filter the Activities list by providing a query using Qualys syntax.
	  In: query
	*/
	Filter *string
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
	/*Sort the results based on field name
	  In: query
	*/
	SortField *string
	/*SortOrder
	  In: query
	*/
	SortOrder *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetAllNotificationsUsingGETParams() beforehand.
func (o *GetAllNotificationsUsingGETParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qCloudType, qhkCloudType, _ := qs.GetOK("cloudType")
	if err := o.bindCloudType(qCloudType, qhkCloudType, route.Formats); err != nil {
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

	qSortField, qhkSortField, _ := qs.GetOK("sortField")
	if err := o.bindSortField(qSortField, qhkSortField, route.Formats); err != nil {
		res = append(res, err)
	}

	qSortOrder, qhkSortOrder, _ := qs.GetOK("sortOrder")
	if err := o.bindSortOrder(qSortOrder, qhkSortOrder, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindCloudType binds and validates parameter CloudType from query.
func (o *GetAllNotificationsUsingGETParams) bindCloudType(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("cloudType", "query", rawData)
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// AllowEmptyValue: false
	if err := validate.RequiredString("cloudType", "query", raw); err != nil {
		return err
	}

	o.CloudType = raw

	if err := o.validateCloudType(formats); err != nil {
		return err
	}

	return nil
}

// validateCloudType carries on validations for parameter CloudType
func (o *GetAllNotificationsUsingGETParams) validateCloudType(formats strfmt.Registry) error {

	if err := validate.EnumCase("cloudType", "query", o.CloudType, []interface{}{"AWS", "AZURE", "GCP"}, true); err != nil {
		return err
	}

	return nil
}

// bindFilter binds and validates parameter Filter from query.
func (o *GetAllNotificationsUsingGETParams) bindFilter(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetAllNotificationsUsingGETParams) bindPageNo(rawData []string, hasKey bool, formats strfmt.Registry) error {
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
func (o *GetAllNotificationsUsingGETParams) bindPageSize(rawData []string, hasKey bool, formats strfmt.Registry) error {
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

// bindSortField binds and validates parameter SortField from query.
func (o *GetAllNotificationsUsingGETParams) bindSortField(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.SortField = &raw

	return nil
}

// bindSortOrder binds and validates parameter SortOrder from query.
func (o *GetAllNotificationsUsingGETParams) bindSortOrder(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.SortOrder = &raw

	return nil
}
