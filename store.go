package users

type Store struct {
	users []User
}

func (s *Store) AddUser(u User) {
	s.users = append(s.users, u)
}

func (s *Store) ListUsers() []User {
	return s.users
}
