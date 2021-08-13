/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package response_rules

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DisableRuleUsingPOSTHandlerFunc turns a function with the right signature into a disable rule using p o s t handler
type DisableRuleUsingPOSTHandlerFunc func(DisableRuleUsingPOSTParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DisableRuleUsingPOSTHandlerFunc) Handle(params DisableRuleUsingPOSTParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DisableRuleUsingPOSTHandler interface for that can handle valid disable rule using p o s t params
type DisableRuleUsingPOSTHandler interface {
	Handle(DisableRuleUsingPOSTParams, interface{}) middleware.Responder
}

// NewDisableRuleUsingPOST creates a new http.Handler for the disable rule using p o s t operation
func NewDisableRuleUsingPOST(ctx *middleware.Context, handler DisableRuleUsingPOSTHandler) *DisableRuleUsingPOST {
	return &DisableRuleUsingPOST{Context: ctx, Handler: handler}
}

/*DisableRuleUsingPOST swagger:route POST /rest/v1/rules/disable Response Rules disableRuleUsingPOST

Disable enabled rules

*/
type DisableRuleUsingPOST struct {
	Context *middleware.Context
	Handler DisableRuleUsingPOSTHandler
}

func (o *DisableRuleUsingPOST) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDisableRuleUsingPOSTParams()

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
