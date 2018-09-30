package sessions

import (
	"net/http"
)

type SessionStorage interface {
	GetSession(r *http.Request) (*Session, error)
	SaveSession(w http.ResponseWriter, session Session) error
	DeleteSession(w http.ResponseWriter, r *http.Request) error
}
