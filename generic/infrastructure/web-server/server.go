package webserver

import (
	"log"
	"net/http"
	"strconv"
)

type HttpHandler = func(http.ResponseWriter, *http.Request)

type WebServerController interface {
	Handler(http.ResponseWriter, *http.Request)
	GetPattern() string
}

type WebServer struct{}

func (webServer WebServer) Listen(port int) {
	PORT_PARAM := ":" + strconv.Itoa(port)

	log.Fatal(http.ListenAndServe(PORT_PARAM, nil))
}

func (webServer WebServer) Register(controller WebServerController) {
	http.HandleFunc(controller.GetPattern(), controller.Handler)
}

func Create() WebServer {
	return WebServer{}
}
