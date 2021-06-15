package main

import (
	"database/sql"
	"fmt"

	"github.com/ashokdey/learning-golang/dbutil"
	"github.com/ashokdey/learning-golang/services"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("MySQL with Golang")

	// connect to database
	db, err := sql.Open(dbutil.Driver, dbutil.DBSource)
	dbutil.Check(err)
	defer db.Close()
	fmt.Println("Successfully connected to MySQL")

	fmt.Println("Creating data...")
	services.CreateData(db, dbutil.TestData{ID: 1, Name: "John Doe"})
	services.CreateData(db, dbutil.TestData{ID: 2, Name: "Marry Cage"})
	newId := services.CreateData(db, dbutil.TestData{ID: 3, Name: "Adam Carpenter"})
	fmt.Println("Rows created: ", newId)

	fmt.Println("Reading data...")
	dataset := services.ReadData(db)
	fmt.Println("Data from DB :", dataset)

	fmt.Println("Updating data...")
	rowsAffected := services.UpdateData(db, dbutil.TestData{ID: 2, Name: "Batman"})
	fmt.Println("Rows updated: ", rowsAffected)

	fmt.Println("Deleting data...")
	rowsAffected = services.DeleteData(db, dbutil.TestData{ID: 2, Name: "Batman"})
	fmt.Println("Rows updated: ", rowsAffected)

	fmt.Println("Remaining data...")
	list := services.ReadData(db)
	fmt.Println(list)

	services.Cleanup(db)
}
