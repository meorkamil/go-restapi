package server

import (
	"go-restapi/internal/database"
	"go-restapi/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerNewAPIServer(t *testing.T) {
	// Just invoke new NewAPIServer
	file := "../../config/config.yml"
	config := util.ConfigInit(file)
	server := NewAPIServer(config)
	assert.NotNil(t, server)
}

func TestServerAppUser(t *testing.T) {
	file := "../../config/config.yml"
	c := util.ConfigInit(file)

	db := database.InitDB(c)
	if err := db.Migration(); err != nil {
		t.Errorf("TestServerAppUser %s", err)
	}

	con, err := db.Con()
	if err != nil {
		t.Errorf("TestServerAppUser %s", err)
	}

	server := NewAPIServer(c)
	if err := appUser(server, con); err != nil {
		t.Errorf("TestServerAppUser %s", err)
	}
}
