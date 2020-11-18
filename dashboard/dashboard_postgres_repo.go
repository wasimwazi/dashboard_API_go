package dashboard

import (
	"database/sql"
	"fmt"
)

//PostgresRepo : Dashboard Repo struct for Postgres
type PostgresRepo struct {
	DB *sql.DB
}

//GetDataFromDB : Get the area distribution from DB
func (pg *PostgresRepo) GetDataFromDB(keyword string, requestType string) ([]DataFromDB, error) {
	query := `
		SELECT
			distribution.place, product.name, distributor.name, distribution.quantity_sold, distribution.product_id, distribution.distributor_id
		FROM 
			distribution
		INNER JOIN 
			product
		ON
			product.id = distribution.product_id
		INNER JOIN 
			distributor
		ON 
			distributor.id = distribution.distributor_id
		%s
	`
	var whereQuery string
	if len(keyword) > 0 {
		if requestType == "area" {
			whereQuery = `WHERE distribution.place ILIKE '` + keyword + `'`
		} else if requestType == "product" {
			whereQuery = `WHERE product.name ILIKE '` + keyword + `'`
		} else if requestType == "distributor" {
			whereQuery = `WHERE distributor.name ILIKE '` + keyword + `'`
		}
	} else {
		whereQuery = ``
	}
	query = fmt.Sprintf(query, whereQuery)
	rows, err := pg.DB.Query(query)
	if err != nil {
		return nil, err
	}
	var data []DataFromDB
	var singleData DataFromDB
	for rows.Next() {
		err := rows.Scan(&singleData.Place, &singleData.ProductName, &singleData.DistributorName, &singleData.Quantity, &singleData.ProductID, &singleData.DistributorID)
		if err != nil {
			return nil, err
		}
		data = append(data, singleData)
	}
	return data, nil
}

//GetTopN : Get top N from each place for a
func (pg *PostgresRepo) GetTopN(n int, distributorName string) ([]TopNDataFromDB, error) {
	query := `
		SELECT
			rank_filter.*
		FROM (
			SELECT
				distribution.place, product.name, distributor.name, distribution.quantity_sold, distribution.product_id, distribution.distributor_id, rank()
			OVER (
				PARTITION BY 
					distribution.place
				ORDER BY
					product.name DESC
			)
			FROM 
				distribution
			INNER JOIN 
				product
			ON
				product.id = distribution.product_id
			INNER JOIN 
				distributor
			ON 
				distributor.id = distribution.distributor_id
			WHERE distributor.name ILIKE '%s' 
		) rank_filter 
		WHERE RANK <= $1
	`
	query = fmt.Sprintf(query, distributorName)
	rows, err := pg.DB.Query(query, n)
	if err != nil {
		return nil, err
	}
	var data []TopNDataFromDB
	var singleData TopNDataFromDB
	for rows.Next() {
		err := rows.Scan(&singleData.Place, &singleData.ProductName, &singleData.DistributorName, &singleData.Quantity, &singleData.ProductID, &singleData.DistributorID, &singleData.Rank)
		if err != nil {
			return nil, err
		}
		data = append(data, singleData)
	}
	return data, nil
}
