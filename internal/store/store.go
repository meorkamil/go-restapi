package store

import (
	"database/sql"
	"go-restapi/internal/model"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

//type Store interface {
//	Select(emp *model.Employee)
//	ListEmployee()
//}

func InitDB(dsn string) error {

	var err error

	log.Println("Connecting to database:", dsn)

	db, err = sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)

	}

	return db.Ping()
}

func Select(emp *model.Employee) []model.Employee {

	rows, err := db.Query("SELECT * FROM employee WHERE empid = ($1) LIMIT 1", emp.Empid)

	if err != nil {
		log.Fatal("Select data failed ", err)
		return nil
	}

	var listOfEmployee []model.Employee

	for rows.Next() {

		var emp model.Employee

		if err := rows.Scan(&emp.Empid, &emp.Name, &emp.Dept); err != nil {

			log.Fatal("Error mapping data")

		}

		listOfEmployee = append(listOfEmployee, emp)
	}

	return listOfEmployee
}

func ListEmployee() []model.Employee {

	rows, err := db.Query("SELECT * FROM employee;")

	if err != nil {
		log.Fatal("Select data failed ", err)
		return nil
	}

	var listOfEmployee []model.Employee

	for rows.Next() {

		var emp model.Employee

		if err := rows.Scan(&emp.Empid, &emp.Name, &emp.Dept); err != nil {

			log.Fatal("Error mapping data")

		}

		listOfEmployee = append(listOfEmployee, emp)
	}

	return listOfEmployee

}
