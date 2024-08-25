package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"millanuka.com/stripe-checkout/models"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "product.db")

	if err != nil {
		log.Printf("There was an error when connecting to the database: ", err)
		return nil, err
	}

	return db, nil
}

func LoadAllProducts() ([]models.Product, error) {
	db, err := ConnectDB()

	if err != nil {
		log.Print("There was an error connecting to the database")
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM product")

	if err != nil {
		log.Print("Query failed to execute")
	}

	defer rows.Close()
	products := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		err = rows.Scan(&product.ID)

		if err != nil {
			log.Print("There was an error reading in from the product database")
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil

}
