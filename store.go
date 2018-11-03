package users

type Store struct {
	users          []User
	lastInsertedID int64
}

func NewStore() *Store {
	s := Store{
		users: make([]User, 0, 0),
	}

	return &s
}

func (s *Store) AddUser(u User) int64 {
	s.lastInsertedID++
	u.ID = s.lastInsertedID
	s.users = append(s.users, u)
	return s.lastInsertedID
}

func (s *Store) ListUsers() []User {
	return s.users
}
