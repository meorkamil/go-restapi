package router

import (
	"go-restapi/internal/database"
	"go-restapi/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRouter(t *testing.T) {
	file := "../../config/config.yml"
	c := util.ConfigInit(file)

	db := database.InitDB(c)

	con, err := db.Con()
	if err != nil {
		t.Errorf("TestCreateRouter %s", err)
	}

	mux := CreateRouter(con)

	assert.NotNil(t, mux)
}
