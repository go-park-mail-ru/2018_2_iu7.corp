package sessions

import (
	"2018_2_iu7.corp/errs"
	"math/rand"
	"net/http"
	"sync"
)

const (
	SessionIDLength = 100
)

type InMemorySessionStorage struct {
	rSeqGen inMemorySessionStorageSequenceGenerator
	storage []Session
	rwMutex *sync.RWMutex
}

func NewInMemorySessionStorage() *InMemorySessionStorage {
	g := newInMemorySessionStorageSequenceGenerator(SessionIDLength)
	if g == nil {
		return nil
	}

	return &InMemorySessionStorage{
		rSeqGen: *g,
		storage: make([]Session, 0),
		rwMutex: &sync.RWMutex{},
	}
}

func (s *InMemorySessionStorage) GetSession(r *http.Request) (*Session, error) {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil, err
	}

	for _, session := range s.storage {
		if session.SessionID == cookie.Value {
			return &session, nil
		}
	}

	return nil, errs.NewNotFoundError("sessions not found")
}

func (s *InMemorySessionStorage) SaveSession(w http.ResponseWriter, session Session) error {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	session.SessionID = s.rSeqGen.getSequence()
	s.storage = append(s.storage, session)

	http.SetCookie(w, &http.Cookie{
		Name:  "session_id",
		Value: session.SessionID,
	})

	return nil
}

func (s *InMemorySessionStorage) DeleteSession(w http.ResponseWriter, r *http.Request) error {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	cookie, err := r.Cookie("session_id")
	if err != nil {
		return nil
	}

	index := -1
	for i, session := range s.storage {
		if session.SessionID == cookie.Value {
			index = i
			break
		}
	}

	if index == -1 {
		return nil
	}

	s.storage = append(s.storage[0:index], s.storage[index+1:]...)

	return nil
}

type inMemorySessionStorageSequenceGenerator struct {
	sequenceLen  int
	allowedRunes []rune
}

func newInMemorySessionStorageSequenceGenerator(n int) *inMemorySessionStorageSequenceGenerator {
	return &inMemorySessionStorageSequenceGenerator{
		sequenceLen:  n,
		allowedRunes: []rune(`abcdefghijklmnopqrstuvwxyz1234567890@#$^&*()_-=+`),
	}
}

func (s *inMemorySessionStorageSequenceGenerator) getSequence() string {
	b := make([]rune, s.sequenceLen)

	n := len(s.allowedRunes)
	for i := range b {
		b[i] = s.allowedRunes[rand.Intn(n)]
	}

	return string(b)
}
