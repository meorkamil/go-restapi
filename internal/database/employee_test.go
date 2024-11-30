package database

import (
	"go-restapi/internal/model"
	"go-restapi/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRepo(t *testing.T) {
	file := "../../config/config.yml"
	c := util.ConfigInit(file)

	db := InitDB(c)

	con, err := db.Con()
	if err != nil {
		t.Errorf("TestDatabaseEmployeeInitRepo %s", err)
	}

	repo := InitRepo(con)

	assert.NotNil(t, repo)
}

func TestDatabaseEmployeeFindByID(t *testing.T) {
	file := "../../config/config.yml"
	c := util.ConfigInit(file)

	db := InitDB(c)

	con, err := db.Con()
	if err != nil {
		t.Errorf("TestDatabaseEmployeeFindByID %s", err)
	}

	repo := InitRepo(con)

	appUser := model.Employee{
		Empid:    "APP1",
		Name:     c.AppAdmin.User,
		Password: c.AppAdmin.Pass,
		Dept:     "Int-APP",
	}

	if err := repo.CreateAppUser(&appUser); err != nil {
		t.Errorf("TestDatabaseEmployeeFindByID %s", err)
	}

	emp, err := repo.FindByID("APP1", c.AppAdmin.Pass)
	if err != nil {
		t.Errorf("TestDatabaseEmployeeFindByID %s", err)
	}

	assert.NotNil(t, emp)
}

func TestDatabaseEmployeeFindAll(t *testing.T) {
	file := "../../config/config.yml"
	c := util.ConfigInit(file)

	db := InitDB(c)

	con, err := db.Con()
	if err != nil {
		t.Errorf("TestDatabaseEmployeeFindAll %s", err)
	}

	repo := InitRepo(con)

	emp, err := repo.FindAll()
	if err != nil {
		t.Errorf("TestDatabaseEmployeeFindAll %s", err)
	}

	assert.NotNil(t, emp)
}

func TestDatabaseCreateUser(t *testing.T) {
	file := "../../config/config.yml"
	c := util.ConfigInit(file)

	db := InitDB(c)

	con, err := db.Con()

	if err != nil {
		t.Errorf("TestDatabaseEmployeeFindAll %s", err)
	}
	appUser := model.Employee{
		Empid:    "TEST-USER",
		Name:     "TEST-USER",
		Password: "TEST-USER",
		Dept:     "TEST-USER",
	}

	repo := InitRepo(con)

	if err := repo.CreateAppUser(&appUser); err != nil {
		t.Errorf("TestDatabaseEmployeeFindByID %s", err)
	}
}
