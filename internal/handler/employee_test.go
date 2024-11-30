package handler

import (
	"bytes"
	"go-restapi/internal/database"
	"go-restapi/internal/util"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandlerEmployeeListEmployee(t *testing.T) {
	file := "../../config/config.yml"
	c := util.ConfigInit(file)

	db := database.InitDB(c)

	con, err := db.Con()

	if err != nil {
		t.Errorf("TestHandlerEmployeeListEmployee %s", err)
	}

	handler := ListEmployee(con)

	req := httptest.NewRequest(http.MethodGet, "/employee", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestHandlerEmployeeCreateEmployee(t *testing.T) {
	file := "../../config/config.yml"
	c := util.ConfigInit(file)

	db := database.InitDB(c)

	con, err := db.Con()

	if err != nil {
		t.Errorf("TestHandlerEmployeeCreateEmployee %s", err)
	}

	handler := ListEmployee(con)

	req := httptest.NewRequest(http.MethodPost, "/employee", bytes.NewReader([]byte(`{"name": "TEST-USAER","empid": "TEST-USER","dept": "TEST-USER","password": "TEST-USER"}`)))
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
}
