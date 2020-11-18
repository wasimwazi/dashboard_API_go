package user

import "database/sql"

//Repo : User repo interface
type Repo interface {
	Login(*Login) (*User, error)
}

// NewRepo : Returns User Repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
