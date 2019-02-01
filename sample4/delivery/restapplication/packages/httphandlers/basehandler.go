package httphandlers

import (
	"net/http"

	"github.com/priteshgudge/gohttpexamples/sample4/delivery/restapplication/packages/resputl"
)

//BaseHandler is base for all Handlers
type BaseHandler struct{}

//GetOne returns single object with http GET method
func (p *BaseHandler) GetOne(r *http.Request, id string) resputl.SrvcRes {
	return resputl.ResponseNotImplemented(nil)
}

// Get http Get method
func (p *BaseHandler) Get(r *http.Request) resputl.SrvcRes {
	return resputl.ResponseNotImplemented(nil)
}

//Put http PUT method
func (p *BaseHandler) Put(r *http.Request) resputl.SrvcRes {
	return resputl.ResponseNotImplemented(nil)
}

//Post http POST method
func (p *BaseHandler) Post(r *http.Request) resputl.SrvcRes {
	return resputl.ResponseNotImplemented(nil)
}

//Delete http DELETE method
func (p *BaseHandler) Delete(r *http.Request) resputl.SrvcRes {
	return resputl.ResponseNotImplemented(nil)
}

//Patch http PATCH method
func (p *BaseHandler) Patch(r *http.Request) resputl.SrvcRes {
	return resputl.ResponseNotImplemented(nil)
}

//Options http OPTIONS method
func (p *BaseHandler) Options(r *http.Request) resputl.SrvcRes {
	return resputl.OptionsResponseOK("OK")
}
