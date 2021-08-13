/* #nosec */ // Code generated by go-swagger; DO NOT EDIT.

package g_c_p_evaluations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetEvaluationsStatsUsingGET2HandlerFunc turns a function with the right signature into a get evaluations stats using g e t 2 handler
type GetEvaluationsStatsUsingGET2HandlerFunc func(GetEvaluationsStatsUsingGET2Params, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetEvaluationsStatsUsingGET2HandlerFunc) Handle(params GetEvaluationsStatsUsingGET2Params, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetEvaluationsStatsUsingGET2Handler interface for that can handle valid get evaluations stats using g e t 2 params
type GetEvaluationsStatsUsingGET2Handler interface {
	Handle(GetEvaluationsStatsUsingGET2Params, interface{}) middleware.Responder
}

// NewGetEvaluationsStatsUsingGET2 creates a new http.Handler for the get evaluations stats using g e t 2 operation
func NewGetEvaluationsStatsUsingGET2(ctx *middleware.Context, handler GetEvaluationsStatsUsingGET2Handler) *GetEvaluationsStatsUsingGET2 {
	return &GetEvaluationsStatsUsingGET2{Context: ctx, Handler: handler}
}

/*GetEvaluationsStatsUsingGET2 swagger:route GET /rest/v1/gcp/evaluations/stats/{controlId}/{connectorId} GCP Evaluations getEvaluationsStatsUsingGET2

Get the stats for specified control id and resource id

*/
type GetEvaluationsStatsUsingGET2 struct {
	Context *middleware.Context
	Handler GetEvaluationsStatsUsingGET2Handler
}

func (o *GetEvaluationsStatsUsingGET2) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetEvaluationsStatsUsingGET2Params()

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
