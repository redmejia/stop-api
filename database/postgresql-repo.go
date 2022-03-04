package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/redmejia/stop/models"
)

// PostgreSQL
type DB struct {
	Db *sql.DB
}

// Retrive method for products pants or shirts
// The scanRows parameter is a function that take a *sql.Rows pointer
func (d *DB) Get(query string, scanRows func(*sql.Rows) (models.Product, error)) ([]models.Product, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var products []models.Product

	rows, err := d.Db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {

		product, err := scanRows(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, product)

	}

	return products, nil

}

func (d *DB) NewOrder() {}
