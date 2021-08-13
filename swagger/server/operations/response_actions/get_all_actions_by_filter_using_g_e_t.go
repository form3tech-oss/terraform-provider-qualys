/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package response_actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetAllActionsByFilterUsingGETHandlerFunc turns a function with the right signature into a get all actions by filter using g e t handler
type GetAllActionsByFilterUsingGETHandlerFunc func(GetAllActionsByFilterUsingGETParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAllActionsByFilterUsingGETHandlerFunc) Handle(params GetAllActionsByFilterUsingGETParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetAllActionsByFilterUsingGETHandler interface for that can handle valid get all actions by filter using g e t params
type GetAllActionsByFilterUsingGETHandler interface {
	Handle(GetAllActionsByFilterUsingGETParams, interface{}) middleware.Responder
}

// NewGetAllActionsByFilterUsingGET creates a new http.Handler for the get all actions by filter using g e t operation
func NewGetAllActionsByFilterUsingGET(ctx *middleware.Context, handler GetAllActionsByFilterUsingGETHandler) *GetAllActionsByFilterUsingGET {
	return &GetAllActionsByFilterUsingGET{Context: ctx, Handler: handler}
}

/*GetAllActionsByFilterUsingGET swagger:route GET /rest/v1/actions Response Actions getAllActionsByFilterUsingGET

Get actions

*/
type GetAllActionsByFilterUsingGET struct {
	Context *middleware.Context
	Handler GetAllActionsByFilterUsingGETHandler
}

func (o *GetAllActionsByFilterUsingGET) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAllActionsByFilterUsingGETParams()

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
