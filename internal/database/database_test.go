package database

import (
	"go-restapi/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDatabaseInitDb(t *testing.T) {
	file := "../../config/config.yml"
	c := util.ConfigInit(file)

	db := InitDB(c)

	con, err := db.Con()
	if err != nil {
		t.Errorf("TestCreateRouter %s", err)
	}

	assert.NotNil(t, con)
}

func TestDatabaseMigration(t *testing.T) {
	file := "../../config/config.yml"
	c := util.ConfigInit(file)

	db := InitDB(c)

	if err := db.Migration(); err != nil {
		t.Errorf("TestDatabaseMigration %s", err)
	}
}

func TestDatabaseDsn(t *testing.T) {
	file := "../../config/config.yml"
	c := util.ConfigInit(file)

	db := InitDB(c)
	dsn := db.dsn()

	assert.NotNil(t, dsn)
}
