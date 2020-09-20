package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	db, err := sql.Open("sqlite3", "./book-requests.db")
	if err != nil {
		panic(err)
	}

	// Restart with a fresh DB, for the sake of testing
	_, err = db.Exec("DROP TABLE IF EXISTS requests")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE requests (id INTEGER AUTO_INCREMENT PRIMARY KEY, title TEXT, email TEXT)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS books (id INTEGER AUTO_INCREMENT PRIMARY KEY, title TEXT, available BOOLEAN)")
	if err != nil {
		panic(err)
	}

	err = populateBooks(db)
	if err != nil {
		panic(err)
	}

	api := RequestAPI{database: db}
	r.POST("/request", api.Requesting)
	r.GET("/request", api.FetchingAll)
	r.GET("/request/:id", api.Fetching)
	r.DELETE("/:id", api.Deleting)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


func populateBooks(db *sql.DB) error {
	books := map[string]bool{
		"ABC": true,
		"XYZ": false,
		"Adventures of Huckleberry Fin": true,
		"Harry Potter and the Sorcerers Stone": false,
		"Words of Radiance": false,
	}

	for title, avail := range books {
		stmt, err := db.Prepare("INSERT INTO books(title, available) values(?,?)")
		if err != nil {
			return err
		}

		_, err = stmt.Exec(title, avail)
		if err != nil {
			return err
		}
	}

	return nil
}
