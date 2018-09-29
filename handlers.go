package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	internalServerError = errors.New("500 internal server error")
)

func RegisterRequestHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		r.Body.Close()

		var p Profile
		if err = json.Unmarshal(reqBody, &p); err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		if err = p.ValidateNew(); err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		id, err := profileRepository.SaveNew(p)
		if err != nil {
			switch err.(type) {
			case *AlreadyExistsError:
				writeErrorResponse(w, http.StatusConflict, err)
			default:
				writeErrorResponse(w, http.StatusInternalServerError, internalServerError)
			}
			return
		}

		resp := make(map[string]uint64)
		resp["id"] = id

		writeSuccessResponse(w, http.StatusCreated, resp)
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

func ProfileRequestHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO
	})
}

func CurrentProfileRequestHandler(uploadsPath string) http.Handler {
	_ = uploadsPath
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO
	})
}

func LeaderBoardRequestHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//TODO
	})
}

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		status := 0

		wr := ResponseWriter{
			wrFunc: func(statusCode int) {
				status = statusCode
			},
			writer: w,
		}

		h.ServeHTTP(wr, r)
		log.Println(r.Method, r.URL.Path, status)
	})
}

func writeErrorResponse(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	if _, err := w.Write([]byte(err.Error())); err != nil {
		panic(err)
	}
}

func writeSuccessResponse(w http.ResponseWriter, status int, resp interface{}) {
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(status)
	if _, err := w.Write(jsonResp); err != nil {
		panic(err)
	}
}
