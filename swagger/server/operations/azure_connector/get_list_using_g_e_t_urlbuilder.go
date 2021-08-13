/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package azure_connector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// GetListUsingGETURL generates an URL for the get list using g e t operation
type GetListUsingGETURL struct {
	Filter   *string
	PageNo   int32
	PageSize int32
	Sort     *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetListUsingGETURL) WithBasePath(bp string) *GetListUsingGETURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetListUsingGETURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetListUsingGETURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/rest/v1/azure/connectors"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/cloudview-api"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var filterQ string
	if o.Filter != nil {
		filterQ = *o.Filter
	}
	if filterQ != "" {
		qs.Set("filter", filterQ)
	}

	pageNoQ := swag.FormatInt32(o.PageNo)
	if pageNoQ != "" {
		qs.Set("pageNo", pageNoQ)
	}

	pageSizeQ := swag.FormatInt32(o.PageSize)
	if pageSizeQ != "" {
		qs.Set("pageSize", pageSizeQ)
	}

	var sortQ string
	if o.Sort != nil {
		sortQ = *o.Sort
	}
	if sortQ != "" {
		qs.Set("sort", sortQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetListUsingGETURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetListUsingGETURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetListUsingGETURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetListUsingGETURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetListUsingGETURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetListUsingGETURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
