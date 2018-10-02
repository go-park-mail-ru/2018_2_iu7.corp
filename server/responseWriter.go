package server

import (
	"net/http"
)

type WriteHeaderFunc func(statusCode int)

type ResponseWriter struct {
	wrFunc WriteHeaderFunc
	writer http.ResponseWriter
}

func (rw ResponseWriter) Header() http.Header {
	return rw.writer.Header()
}

func (rw ResponseWriter) WriteHeader(statusCode int) {
	rw.wrFunc(statusCode)
	rw.writer.WriteHeader(statusCode)
}

func (rw ResponseWriter) Write(data []byte) (int, error) {
	return rw.writer.Write(data)
}
