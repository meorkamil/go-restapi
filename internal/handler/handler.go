package handler

import (
	"fmt"
	"go-restapi/internal/database"
	"go-restapi/internal/model"
	"go-restapi/internal/util"
	"log"
	"net/http"
)

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {

	emp, err := util.Decode[model.Employee](r)

	if err != nil {

		http.Error(w, "Invalid request payload", http.StatusBadRequest)

	}

	req := &model.Employee{
		Empid: emp.Empid,
		Name:  emp.Name,
		Dept:  emp.Dept,
	}

	log.Println(r.Method, r.URL, ": Requeest data from db")

	data := database.Select(req)

	if err := util.Encode(w, r, http.StatusOK, data); err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

}

func ListEmployee(w http.ResponseWriter, r *http.Request) {

	log.Println(r.Method, r.URL, ": Request data from db")

	data := database.ListEmployee()

	if err := util.Encode(w, r, http.StatusOK, data); err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

}

func DefaultRouter(w http.ResponseWriter, r *http.Request) {

	log.Println(r.URL)
	fmt.Fprintf(w, "RESTful API")

}
