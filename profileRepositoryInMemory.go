package main

import (
	"sync"
)

type InMemoryProfileRepository struct {
	idSequence inMemoryProfileRepositoryIDSequence
	storage    []Profile
	rwMutex    *sync.RWMutex
}

func NewInMemoryProfileRepository() *InMemoryProfileRepository {
	idSequence := newInMemoryProfileRepositoryIDSequence()
	if idSequence == nil {
		return nil
	}

	return &InMemoryProfileRepository{
		idSequence: *idSequence,
		rwMutex:    &sync.RWMutex{},
	}
}

func (r *InMemoryProfileRepository) SaveNew(p Profile) (id uint64, err error) {
	r.rwMutex.Lock()
	defer r.rwMutex.Unlock()

	if r.findByUsername(p.Username) != nil {
		return 0, NewAlreadyExistsError("username already taken")
	}
	if r.findByEmail(p.Email) != nil {
		return 0, NewAlreadyExistsError("profile with the email already exists")
	}

	p.ID = r.idSequence.nextValue()
	r.storage = append(r.storage, p)

	return p.ID, nil
}

func (r *InMemoryProfileRepository) SaveExisting(p Profile) (err error) {
	r.rwMutex.Lock()
	defer r.rwMutex.Unlock()

	pPtr := r.findByID(p.ID)
	if pPtr == nil {
		return NewNotFoundError("profile not found")
	}

	*pPtr = p

	return nil
}

func (r *InMemoryProfileRepository) DeleteByID(id uint64) (err error) {
	r.rwMutex.Lock()
	defer r.rwMutex.Unlock()

	index := r.findIndexByID(id)
	if index == -1 {
		return NewNotFoundError("profile not found")
	}

	r.storage = append(r.storage[:index], r.storage[index+1:]...)

	return nil
}

func (r *InMemoryProfileRepository) FindByID(id uint64) (p Profile, err error) {
	r.rwMutex.RLock()
	defer r.rwMutex.RUnlock()

	pPtr := r.findByID(id)
	if pPtr == nil {
		return Profile{}, NewNotFoundError("profile not found")
	}

	p = *pPtr
	return p, nil
}

func (r *InMemoryProfileRepository) FindByUsernameAndPassword(username, password string) (p Profile, err error) {
	r.rwMutex.RLock()
	defer r.rwMutex.RUnlock()

	pPtr := r.findByUsernameAndPassword(username, password)
	if pPtr == nil {
		return Profile{}, NewNotFoundError("profile not found")
	}

	p = *pPtr
	return p, nil
}

func (r *InMemoryProfileRepository) findByID(id uint64) (p *Profile) {
	for _, v := range r.storage {
		if v.ID == id {
			return &v
		}
	}
	return nil
}

func (r *InMemoryProfileRepository) findIndexByID(id uint64) int {
	for i, v := range r.storage {
		if v.ID == id {
			return i
		}
	}
	return -1
}

func (r *InMemoryProfileRepository) findByUsername(username string) (p *Profile) {
	for _, v := range r.storage {
		if v.Username == username {
			return &v
		}
	}
	return nil
}

func (r *InMemoryProfileRepository) findByUsernameAndPassword(username, password string) (p *Profile) {
	for _, v := range r.storage {
		if v.Username == username && v.Password == password {
			return &v
		}
	}
	return nil
}

func (r *InMemoryProfileRepository) findByEmail(email string) (p *Profile) {
	for _, v := range r.storage {
		if v.Email == email {
			return &v
		}
	}
	return nil
}

type inMemoryProfileRepositoryIDSequence struct {
	currentValue uint64
}

func newInMemoryProfileRepositoryIDSequence() *inMemoryProfileRepositoryIDSequence {
	return &inMemoryProfileRepositoryIDSequence{
		currentValue: 0,
	}
}

func (s *inMemoryProfileRepositoryIDSequence) nextValue() uint64 {
	s.currentValue++
	return s.currentValue
}
