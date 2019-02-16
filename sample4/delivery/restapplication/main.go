package main

import (
	logger "log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/priteshgudge/gohttpexamples/sample4/dbrepo/userrepo"
	handlerlib "github.com/priteshgudge/gohttpexamples/sample4/delivery/restapplication/packages/httphandlers"
	"github.com/priteshgudge/gohttpexamples/sample4/delivery/restapplication/usercrudhandler"
)

func init() {
	/*
	   Safety net for 'too many open files' issue on legacy code.
	   Set a sane timeout duration for the http.DefaultClient, to ensure idle connections are terminated.
	   Reference: https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	   https://stackoverflow.com/questions/37454236/net-http-server-too-many-open-files-error
	*/
	http.DefaultClient.Timeout = time.Minute * 10
}

func main() {

	dbrepo := userrepo.NewUserInMemRepository()
	usersvc := userrepo.NewService(dbrepo)

	hndlr := usercrudhandler.NewUserCrudHandler(usersvc)

	pingHandler := &handlerlib.PingHandler{}
	logger.Println("Setting up resources.")
	logger.Println("Starting service")
	h := mux.NewRouter()
	h.Handle("/ping/", pingHandler)
	h.Handle("/user/{id}", hndlr)
	h.Handle("/user/", hndlr)

	logger.Println("Resource Setup Done.")
	logger.Println(http.ListenAndServe(":8011", h))
}
