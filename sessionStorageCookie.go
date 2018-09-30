package main

import (
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"net/http"
)

type CookieSessionStorage struct {
	store *sessions.CookieStore
}

func NewCookieSessionStorage(sessionKey string) *CookieSessionStorage {
	var sk []byte
	if sessionKey == "" {
		sk = securecookie.GenerateRandomKey(32)
	} else {
		sk = []byte(sessionKey)
	}

	store := sessions.NewCookieStore(sk)
	if store == nil {
		return nil
	}

	return &CookieSessionStorage{store}
}

func (s *CookieSessionStorage) GetSession(r *http.Request) (*Session, error) {
	cookie, err := s.store.Get(r, "session")
	if err != nil {
		return nil, err
	}

	var ok bool
	session := &Session{}

	session.Authorized, ok = cookie.Values["authorized"].(bool)
	if !ok {
		return nil, NewInvalidFormatError("invalid session cookie")
	}

	session.ProfileID, ok = cookie.Values["profile_id"].(uint64)
	if !ok {
		return nil, NewInvalidFormatError("invalid session cookie")
	}

	return session, nil
}

func (s *CookieSessionStorage) SaveSession(w http.ResponseWriter, r *http.Request, session Session) error {
	cookie, err := s.store.New(r, "session")
	if err != nil {
		return err
	}

	cookie.Values["authorized"] = session.Authorized
	cookie.Values["profile_id"] = session.ProfileID

	if err = cookie.Save(r, w); err != nil {
		return err
	}

	return nil
}
