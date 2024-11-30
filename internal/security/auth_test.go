package security

import (
	"encoding/json"
	"go-restapi/internal/database"
	"go-restapi/internal/model"
	"go-restapi/internal/util"
	"log"
	"testing"
)

type TestUser struct {
	Empid    string `json:"empid"`
	Password string `json:"password"`
}

func TestAuth(t *testing.T) {
	// Load config, just using docker-compose.yml database for our testing
	file := "../../config/config.yml"
	c := util.ConfigInit(file)

	// Create a database connection
	db := database.InitDB(c)
	con, err := db.Con()
	if err != nil {
		t.Errorf("Test Auth %s", err)
	}

	// Run automigrate
	con.AutoMigrate(
		&model.Employee{},
	)

	// Create a sample user
	jsonData := []byte(`{
        	"empid": "APP1",
        	"password": "admin"
    	}`)

	var user model.Employee

	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		log.Fatal(err)
	}

	// Invoke Auth()
	Auth(con, &user)
}
