package user

import "database/sql"

type PostgresRepo struct {
	DB *sql.DB
}

// Login : Postgres function to validate user crederntials
func (pg *PostgresRepo) Login(login *Login) (*User, error) {
	var user User
	query := `
		SELECT 
			id, username 
		FROM 
			users 
		WHERE 
			username = $1 
		AND 
			password = $2;`
	row := pg.DB.QueryRow(query, login.Username, login.Password)
	err := row.Scan(&user.ID, &user.Username)
	return &user, err
}