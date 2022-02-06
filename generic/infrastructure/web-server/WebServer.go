package webserver

import (
	"log"
	"net/http"
	"strconv"
)

type WebServerController interface {
	Handler(http.ResponseWriter, *http.Request)
	GetPattern() string
}

type WebServer struct{}

func (webServer WebServer) Listen(port int) {
	portParam := ":" + strconv.Itoa(port)

	log.Fatal(http.ListenAndServe(portParam, nil))
}

func (webServer WebServer) Register(controller WebServerController) {
	http.HandleFunc(controller.GetPattern(), controller.Handler)
}

func Create() WebServer {
	return WebServer{}
}
