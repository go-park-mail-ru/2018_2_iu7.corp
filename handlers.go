package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func RegisterRequestHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			writeErrorResponse(w, http.StatusInternalServerError, err)
			return
		}
		r.Body.Close()

		var p Profile
		if err = json.Unmarshal(reqBody, &p); err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		id, err := profileRepository.SaveNew(p)
		if err != nil {
			switch err.(type) {
			case *AlreadyExistsError:
				writeErrorResponse(w, http.StatusConflict, err)
			default:
				writeErrorResponse(w, http.StatusInternalServerError, err)
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

func writeErrorResponse(w http.ResponseWriter, status int, err error) {
	jsonResp, err := json.Marshal(err.Error())
	if err != nil {
		panic(err)
	}

	w.WriteHeader(status)
	w.Write(jsonResp)
}

func writeSuccessResponse(w http.ResponseWriter, status int, resp interface{}) {
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(status)
	w.Write(jsonResp)
}
