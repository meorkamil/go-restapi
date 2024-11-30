package handler

import (
	"bytes"
	"go-restapi/internal/database"
	"go-restapi/internal/model"
	"go-restapi/internal/util"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerLogin(t *testing.T) {
	file := "../../config/config.yml"
	c := util.ConfigInit(file)

	db := database.InitDB(c)

	con, err := db.Con()
	if err != nil {
		t.Errorf("TestHandlerLogin %s", err)
	}

	repo := database.InitRepo(con)

	if err := db.Migration(); err != nil {
		t.Errorf("TestHandlerLogin %s", err)
	}

	appUser := model.Employee{
		Empid:    "TEST-USER",
		Name:     "TEST-USER",
		Password: "TEST-USER",
		Dept:     "TEST-USER",
	}

	if err := repo.CreateAppUser(&appUser); err != nil {
		t.Errorf("TestHandlerLogin %s", err)
	}

	handler := Login(con)

	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader([]byte(`{"empid": "TEST-USER", "password": "TEST-USER"}`)))
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
