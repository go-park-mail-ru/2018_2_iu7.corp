package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func RegisterRequestHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rb, err := parseRequestBody(r)
		if err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		p := &Profile{}
		if err = p.ParseOnRegister(rb); err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		if err = profileRepository.SaveNew(*p); err != nil {
			switch err.(type) {
			case *AlreadyExistsError:
				writeErrorResponse(w, http.StatusConflict, err)
			default:
				writeErrorResponseEmpty(w, http.StatusInternalServerError)
			}
			return
		}

		writeSuccessResponseEmpty(w, http.StatusCreated)
	})
}

func LoginRequestHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rb, err := parseRequestBody(r)
		if err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		p := &Profile{}
		if err = p.ParseOnLogin(rb); err != nil {
			writeErrorResponse(w, http.StatusBadRequest, err)
			return
		}

		exp, err := profileRepository.FindByUsernameAndPassword(p.Username, p.Password)
		if err != nil {
			switch err.(type) {
			case *NotFoundError:
				writeErrorResponse(w, http.StatusNotFound, err)
			default:
				writeErrorResponseEmpty(w, http.StatusInternalServerError)
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
		vars := mux.Vars(r)

		idStr, ok := vars["id"]
		if !ok {
			panic(!ok)
		}

		id, err := strconv.ParseUint(idStr, 0, 64)
		if err != nil {
			panic(err)
		}

		p, err := profileRepository.FindByID(id)
		if err != nil {
			switch err.(type) {
			case *NotFoundError:
				writeErrorResponse(w, http.StatusNotFound, err)
			default:
				writeErrorResponseEmpty(w, http.StatusInternalServerError)
			}
			return
		}

		writeSuccessResponse(w, http.StatusOK, p.GetPublicAttributes())
	})
}

func CurrentProfileRequestHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := sessionStorage.GetSession(r)
		if err != nil {
			panic(err)
		}

		var p Profile
		if p, err = profileRepository.FindByID(session.ProfileID); err != nil {
			panic(err)
		}

		if r.Method == http.MethodGet {
			writeSuccessResponse(w, http.StatusOK, p.GetPrivateAttributes())
		} else {
			rb, err := parseRequestBody(r)
			if err != nil {
				writeErrorResponse(w, http.StatusBadRequest, err)
				return
			}

			if err = p.ParseOnEdit(rb); err != nil {
				writeErrorResponse(w, http.StatusBadRequest, err)
				return
			}

			if err = profileRepository.SaveExisting(p); err != nil {
				switch err.(type) {
				case *NotFoundError:
					writeErrorResponse(w, http.StatusNotFound, err)
				case *AlreadyExistsError:
					writeErrorResponse(w, http.StatusConflict, err)
				default:
					writeErrorResponseEmpty(w, http.StatusInternalServerError)
				}
				return
			}

			writeSuccessResponseEmpty(w, http.StatusOK)
		}
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
