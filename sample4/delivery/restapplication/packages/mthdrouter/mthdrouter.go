package mthdroutr

import (
	"fmt"
	"net/http"

	resputl "github.com/priteshgudge/gohttpexamples/sample4/delivery/restapplication/packages/resputl"
)

//ServiceAPIHandler all http methods
type ServiceAPIHandler interface {
	GetOne(r *http.Request, id string) resputl.SrvcRes
	Get(r *http.Request) resputl.SrvcRes
	Put(r *http.Request) resputl.SrvcRes
	Post(r *http.Request) resputl.SrvcRes
	Delete(r *http.Request) resputl.SrvcRes
	Patch(r *http.Request) resputl.SrvcRes
	Options(r *http.Request) resputl.SrvcRes
}

//RouteAPICall routing to method
func RouteAPICall(sah ServiceAPIHandler, r *http.Request) resputl.SrvcRes {
	switch r.Method {
	case "GET":
		// params := mux.Vars(r)
		// id, present := params["id"]
		// if present {
		// 	return sah.GetOne(r, id)
		// } else {
		// 	return sah.Get(r)
		// }
		return sah.Get(r)
	case "PUT":
		return sah.Put(r)
	case "POST":
		return sah.Post(r)
	case "PATCH":
		return sah.Patch(r)
	case "DELETE":
		return sah.Delete(r)
	case "OPTIONS":
		return sah.Options(r)
	}
	return resputl.SrvcRes{
		Code:     http.StatusMethodNotAllowed,
		Response: fmt.Sprintf("{\"ResponseData\" : \"%s \"}", r.Method),
		Message:  "Method not allowed",
		Headers:  nil}

}
