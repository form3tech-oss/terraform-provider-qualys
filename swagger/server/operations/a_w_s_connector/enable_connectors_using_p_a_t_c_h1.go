/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package a_w_s_connector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// EnableConnectorsUsingPATCH1HandlerFunc turns a function with the right signature into a enable connectors using p a t c h 1 handler
type EnableConnectorsUsingPATCH1HandlerFunc func(EnableConnectorsUsingPATCH1Params, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn EnableConnectorsUsingPATCH1HandlerFunc) Handle(params EnableConnectorsUsingPATCH1Params, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// EnableConnectorsUsingPATCH1Handler interface for that can handle valid enable connectors using p a t c h 1 params
type EnableConnectorsUsingPATCH1Handler interface {
	Handle(EnableConnectorsUsingPATCH1Params, interface{}) middleware.Responder
}

// NewEnableConnectorsUsingPATCH1 creates a new http.Handler for the enable connectors using p a t c h 1 operation
func NewEnableConnectorsUsingPATCH1(ctx *middleware.Context, handler EnableConnectorsUsingPATCH1Handler) *EnableConnectorsUsingPATCH1 {
	return &EnableConnectorsUsingPATCH1{Context: ctx, Handler: handler}
}

/*EnableConnectorsUsingPATCH1 swagger:route PATCH /rest/v1/aws/connectors/connectors/enable AWS Connector enableConnectorsUsingPATCH1

Enable the provided connectors

*/
type EnableConnectorsUsingPATCH1 struct {
	Context *middleware.Context
	Handler EnableConnectorsUsingPATCH1Handler
}

func (o *EnableConnectorsUsingPATCH1) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewEnableConnectorsUsingPATCH1Params()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
