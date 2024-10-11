package database

import (
	"database/sql"
	"go-restapi/internal/model"
	"go-restapi/internal/util"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	lg = util.NewLogger()
)

func InitDB(dsn string) error {
	var err error

	lg.Info("Connecting to database:", dsn)

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		lg.Fatal(err)

	}

	return db.Ping()
}

func Select(emp *model.Employee) []model.Employee {

	rows, err := db.Query("SELECT * FROM employee WHERE empid = ($1) LIMIT 1", emp.Empid)

	if err != nil {
		lg.Fatal("Select data failed ", err)
		return nil
	}

	var listOfEmployee []model.Employee

	for rows.Next() {
		var emp model.Employee
		if err := rows.Scan(&emp.Empid, &emp.Name, &emp.Dept); err != nil {

			lg.Fatal("Error mapping data")
		}

		listOfEmployee = append(listOfEmployee, emp)
	}

	return listOfEmployee
}

func ListEmployee() []model.Employee {
	rows, err := db.Query("SELECT * FROM employee;")

	if err != nil {
		lg.Fatal("Select data failed ", err)
		return nil
	}

	var listOfEmployee []model.Employee

	for rows.Next() {

		var emp model.Employee
		if err := rows.Scan(&emp.Empid, &emp.Name, &emp.Dept); err != nil {

			lg.Fatal("Error mapping data")

		}

		listOfEmployee = append(listOfEmployee, emp)
	}

	return listOfEmployee
}
