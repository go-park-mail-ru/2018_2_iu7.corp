package main

import (
	"net/http"
)

func RegisterRequestHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO
	})
}

func LoginRequestHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO
	})
}

func LogoutRequestHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO
	})
}

func ProfileRequestHandler(uploadPath string) http.Handler {
	_ = uploadPath
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO
	})
}

func LeaderBoardRequestHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO
	})
}
