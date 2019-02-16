package httphandlers

import (
	"log"
	"net/http"

	mthdroutr "github.com/priteshgudge/gohttpexamples/sample4/delivery/restapplication/packages/mthdrouter"
	"github.com/priteshgudge/gohttpexamples/sample4/delivery/restapplication/packages/resputl"
)

// PingHandler is a Basic ping utility for the service
type PingHandler struct {
	BaseHandler
}

func (p *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Got ping request: %v", r)
	response := mthdroutr.RouteAPICall(p, r)
	response.RenderResponse(w)
}

// Get function for PingHandler
func (p *PingHandler) Get(r *http.Request) resputl.SrvcRes {

	return resputl.Response200OK("OK")
}
