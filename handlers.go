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
		rb, err := parseRequestBody(r)
		if err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		p, err := ParseProfileOnRegister(rb)
		if err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		id, err := profileRepository.SaveNew(*p)
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
		rb, err := parseRequestBody(r)
		if err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		p, err := ParseProfileOnLogin(rb)
		if err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		exp, err := profileRepository.FindByUsernameAndPassword(p.Username, p.Password)
		if err != nil {
			switch err.(type) {
			case *NotFoundError:
				writeErrorResponse(w, http.StatusNotFound, err)
			default:
				writeErrorResponse(w, http.StatusInternalServerError, internalServerError)
			}
			return
		}

		session := Session{
			Authorized: true,
			ProfileID:  exp.ID,
		}

		if err = sessionStorage.SaveSession(w, r, session); err != nil {
			panic(err)
		}

		writeSuccessResponseEmpty(w, http.StatusOK)
	})
}

func LogoutRequestHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := Session{Authorized: false}
		if err := sessionStorage.SaveSession(w, r, session); err != nil {
			panic(err)
		}
		writeSuccessResponseEmpty(w, http.StatusOK)
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

func AuthenticatedMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := sessionStorage.GetSession(r)
		if err != nil {
			if err = sessionStorage.SaveSession(w, r, Session{Authorized: false}); err != nil {
				panic(err)
			}
			writeErrorResponseEmpty(w, http.StatusUnauthorized)
			return
		}

		if !session.Authorized {
			writeErrorResponseEmpty(w, http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func NotAuthenticatedMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := sessionStorage.GetSession(r)
		if err != nil {
			h.ServeHTTP(w, r)
			return
		}

		if session.Authorized {
			writeErrorResponse(w, http.StatusForbidden, errors.New("already authorized"))
			return
		}

		h.ServeHTTP(w, r)
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

func parseRequestBody(r *http.Request) (map[string]interface{}, error) {
	rb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	r.Body.Close()

	var data map[string]interface{}
	if err = json.Unmarshal(rb, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func writeErrorResponseEmpty(w http.ResponseWriter, status int) {
	http.Error(w, "", status)
}

func writeSuccessResponseEmpty(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
}

func writeErrorResponse(w http.ResponseWriter, status int, err error) {
	http.Error(w, err.Error(), status)
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
