package main

import (
	"encoding/json"
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

func sendErrorResponse(w http.ResponseWriter, status int, err error) {
	jsonResp, err := json.Marshal(err.Error())
	if err != nil {
		panic(err)
	}

	w.WriteHeader(status)
	w.Write(jsonResp)
}

func sendSuccessResponse(w http.ResponseWriter, status int, v interface{}) {
	jsonResp, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(status)
	w.Write(jsonResp)
}
