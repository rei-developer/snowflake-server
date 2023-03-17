package server

import (
	"fmt"

	"github.com/snowflake-server/src/user"
)

type UserObserver interface {
	UpdateUserList(map[uint32]*user.User)
}

func (s *Server) RegisterUserObserver(o UserObserver) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.userObservers = append(s.userObservers, o)
}

func (s *Server) UnregisterUserObserver(o UserObserver) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, observer := range s.userObservers {
		if observer == o {
			s.userObservers = append(s.userObservers[:i], s.userObservers[i+1:]...)
			break
		}
	}
}

func (s *Server) NotifyUserList() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, observer := range s.userObservers {
		observer.UpdateUserList(s.users)
	}
}

type userObserverWrapper struct {
	user *user.User
}

func (u *userObserverWrapper) UpdateUserList(users map[uint32]*user.User) {
	// no-op
}

func (s *Server) addUser(user *user.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[uint32(user.ID)]; ok {
		return fmt.Errorf("user with ID %d already exists", user.ID)
	}

	s.users[uint32(user.ID)] = user

	observer := &userObserverWrapper{user: user}
	s.RegisterUserObserver(observer)

	s.NotifyUserList()

	return nil
}

func (s *Server) removeUser(id uint32) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.users[id]; !ok {
		return fmt.Errorf("user with ID %d does not exist", id)
	}

	delete(s.users, id)

	s.NotifyUserList()

	return nil
}
