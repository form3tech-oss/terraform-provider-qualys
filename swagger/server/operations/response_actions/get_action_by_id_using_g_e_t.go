/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package response_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetActionByIDUsingGETHandlerFunc turns a function with the right signature into a get action by Id using g e t handler
type GetActionByIDUsingGETHandlerFunc func(GetActionByIDUsingGETParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetActionByIDUsingGETHandlerFunc) Handle(params GetActionByIDUsingGETParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetActionByIDUsingGETHandler interface for that can handle valid get action by Id using g e t params
type GetActionByIDUsingGETHandler interface {
	Handle(GetActionByIDUsingGETParams, interface{}) middleware.Responder
}

// NewGetActionByIDUsingGET creates a new http.Handler for the get action by Id using g e t operation
func NewGetActionByIDUsingGET(ctx *middleware.Context, handler GetActionByIDUsingGETHandler) *GetActionByIDUsingGET {
	return &GetActionByIDUsingGET{Context: ctx, Handler: handler}
}

/*GetActionByIDUsingGET swagger:route GET /rest/v1/actions/{actionId} Response Actions getActionByIdUsingGET

Get action by Id

*/
type GetActionByIDUsingGET struct {
	Context *middleware.Context
	Handler GetActionByIDUsingGETHandler
}

func (o *GetActionByIDUsingGET) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetActionByIDUsingGETParams()

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
