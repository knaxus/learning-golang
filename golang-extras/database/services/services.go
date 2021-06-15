package services

import (
	"database/sql"

	"github.com/ashokdey/learning-golang/dbutil"
)

func ReadData(db *sql.DB) []dbutil.TestData {
	var dataset = []dbutil.TestData{}

	res, err := db.Query("SELECT * FROM test_data LIMIT 3")
	dbutil.Check(err)

	for res.Next() {
		data := dbutil.TestData{}
		err := res.Scan(&data.ID, &data.Name)
		dbutil.Check(err)

		dataset = append(dataset, data)
	}
	return dataset
}

func CreateData(db *sql.DB, data dbutil.TestData) int64 {
	query := "INSERT INTO test_data (id, name) VALUES (?, ?)"
	stmt, err := db.Prepare(query)
	dbutil.Check(err)

	// execute the statement
	res, err := stmt.Exec(data.ID, data.Name)
	dbutil.Check(err)

	id, err := res.LastInsertId()
	dbutil.Check(err)

	return id
}

func UpdateData(db *sql.DB, data dbutil.TestData) int64 {
	query := "UPDATE test_data SET name =? WHERE id =?"
	stmt, err := db.Prepare(query)
	dbutil.Check(err)

	res, err := stmt.Exec(data.Name, data.ID)
	dbutil.Check(err)

	rowsChanged, err := res.RowsAffected()
	dbutil.Check(err)

	return rowsChanged
}

func DeleteData(db *sql.DB, data dbutil.TestData) int64 {
	query := "DELETE FROM test_data WHERE id =? AND name = ?"

	stmt, err := db.Prepare(query)
	dbutil.Check(err)

	res, err := stmt.Exec(data.ID, data.Name)
	dbutil.Check(err)

	rowsChanged, err := res.RowsAffected()
	dbutil.Check(err)
	return rowsChanged
}

func Cleanup(db *sql.DB) {
	_, err := db.Query("TRUNCATE TABLE test_data")
	dbutil.Check(err)
}
