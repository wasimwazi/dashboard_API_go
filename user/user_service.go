package user

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

// ServiceInterface : Service interface for user
type ServiceInterface interface {
	LoginUser(*Login) (*User, error)
}

// Service : User Service Struct
type Service struct {
	ar Repo
}

// NewService : Returns User Service
func NewService(db *sql.DB) ServiceInterface {
	return &Service{
		ar: NewRepo(db),
	}
}

// LoginUser : to login User
func (as *Service) LoginUser(user *Login) (*User, error) {
	//Converting the user password to md5 string and checking in DB
	password := []byte(user.Password)
	hash := md5.Sum(password)
	user.Password = hex.EncodeToString(hash[:])
	return as.ar.Login(user)
}