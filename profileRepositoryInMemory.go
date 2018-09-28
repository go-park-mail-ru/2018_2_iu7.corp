package main

import (
	"sync"
)

type InMemoryProfileRepository struct {
	ProfileRepository

	storage []Profile
	rwMutex *sync.RWMutex
}

func NewInMemoryProfileRepository() *InMemoryProfileRepository {
	return &InMemoryProfileRepository{}
}

func (r *InMemoryProfileRepository) SaveNew(p Profile) (id uint64, err error) {
	panic("Not implemented yet!")
}

func (r *InMemoryProfileRepository) SaveExisting(p Profile) (err error) {
	panic("Not implemented yet!")
}

func (r *InMemoryProfileRepository) DeleteByID(id uint64) (err error) {
	panic("Not implemented yet!")
}

func (r *InMemoryProfileRepository) FindByID(id uint64) (p Profile, err error) {
	panic("Not implemented yet!")
}

func (r *InMemoryProfileRepository) FindByUsernameAndPassword(username, password string) (p Profile, err error) {
	panic("Not implemented yet!")
}
