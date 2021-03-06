// Code generated by go-swagger; DO NOT EDIT.

package logs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// SaveLogHandlerFunc turns a function with the right signature into a save log handler
type SaveLogHandlerFunc func(SaveLogParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SaveLogHandlerFunc) Handle(params SaveLogParams) middleware.Responder {
	return fn(params)
}

// SaveLogHandler interface for that can handle valid save log params
type SaveLogHandler interface {
	Handle(SaveLogParams) middleware.Responder
}

// NewSaveLog creates a new http.Handler for the save log operation
func NewSaveLog(ctx *middleware.Context, handler SaveLogHandler) *SaveLog {
	return &SaveLog{Context: ctx, Handler: handler}
}

/*SaveLog swagger:route POST /log logs saveLog

Saves a log to the datastore

*/
type SaveLog struct {
	Context *middleware.Context
	Handler SaveLogHandler
}

func (o *SaveLog) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewSaveLogParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
