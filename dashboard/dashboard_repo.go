package dashboard

import "database/sql"

//Repo : Repository interface
type Repo interface {
	GetDataFromDB(string, string) ([]DataFromDB, error)
	GetTopN(int, string) ([]TopNDataFromDB, error)
}

// NewRepo : Returns Dashboard Repo
func NewRepo(db *sql.DB) Repo {
	return &PostgresRepo{
		DB: db,
	}
}
