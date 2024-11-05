package user

import (
	"database/sql"
	"fmt"
	"log"
	"test-project/types"
	"time"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	log.Println("Received this from db -->", rows)
	if err != nil {
		return nil, err
	}

	fmt.Println("jwhhwhshshks")

	u := new(types.User)
	for rows.Next() {
		u, err = ScanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	log.Println("u ----->", u)
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

func ScanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	var id int64
	var firstName string
	var lastName string
	var email string
	var password string
	var created_at time.Time
	err := rows.Scan(
		&id,
		&firstName,
		&lastName,
		&email,
		&password,
		&created_at,
	)
	user := &types.User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		CreatedAt: created_at,
	}
	log.Println("scanned user-->", user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM ecommerce.users WHERE id = ?", id)
	if err != nil {
		log.Fatal("Unable to get user by id ", id)
		return nil, err
	}
	u := new(types.User)

	for rows.Next() {
		u, err = ScanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return u, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Query("INSERT INTO ecommerce.users (firstName, lastName, email, password) VALUES (?,?,?,?)", user.FirstName, user.LastName, user.Email, user.Password)
	if err != nil {
		fmt.Println("unable to create user account at the moment", err)
		return fmt.Errorf("unable to create user account at the moment")
	}
	return nil
}
