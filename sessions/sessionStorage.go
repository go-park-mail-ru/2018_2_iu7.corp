package sessions

import "net/http"

type SessionStorage interface {
	GetSession(r *http.Request) (*Session, error)
	SaveSession(w http.ResponseWriter, r *http.Request, session Session) error
}
