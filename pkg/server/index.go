package server

import (
	"net/http"
)

type IndexHandler struct{}

func (IndexHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello"))
}
