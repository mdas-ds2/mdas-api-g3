package webserver

import (
	"log"
	"net/http"
	"strconv"
)

type WebServer struct{}

func (webServer WebServer) Listen(port int) {
	PORT_PARAM := ":" + strconv.Itoa(port)

	log.Fatal(http.ListenAndServe(PORT_PARAM, nil))
}

func Create() WebServer {
	return WebServer{}
}
