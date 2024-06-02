package openapigenonlyinterface

import (
	"errors"
	openapigen_gin_interface "github.com/GIT_USER_ID/GIT_REPO_ID/open-oapi-codegen-gin_interface"

	"sync"
	"time"
)

type UserService struct {
	users  map[int64]*openapigen_gin_interface.User
	mu     sync.RWMutex
	nextID int64
}

// NewUserService creates a new UserService
func NewUserService() *UserService {
	return &UserService{
		users:  make(map[int64]*openapigen_gin_interface.User),
		nextID: 1,
	}
}

// CreateUser creates a new user
func (s *UserService) CreateUser(createUser openapigen_gin_interface.CreateUser) (*openapigen_gin_interface.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now()
	user := &openapigen_gin_interface.User{
		CreatedAt: now,
		Email:     createUser.Email,
		Id:        s.nextID,
		Lastname:  createUser.Lastname,
		Name:      createUser.Name,
		UpdatedAt: now,
	}

	s.users[s.nextID] = user
	s.nextID++

	return user, nil
}

// GetUser returns a user by ID
func (s *UserService) GetUser(id int64) (*openapigen_gin_interface.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// UpdateUser updates an existing user
func (s *UserService) UpdateUser(id int64, email, lastname, name string) (*openapigen_gin_interface.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	now := time.Now()
	user.Email = email
	user.Lastname = lastname
	user.Name = name
	user.UpdatedAt = now

	return user, nil
}

// DeleteUser deletes a user by ID
func (s *UserService) DeleteUser(id int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return errors.New("user not found")
	}

	delete(s.users, id)
	return nil
}
