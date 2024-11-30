package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestModelEmployee(t *testing.T) {
	emp := Employee{
		Empid:     "KAMIOPS",
		Name:      "Kamil",
		Dept:      "DevOps",
		Password:  "beLSepolUtAR",
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
		DeletedAt: time.Now(),
	}

	assert.NotNil(t, emp)
}

func TestModelProduct(t *testing.T) {
	prd := Product{
		ID:    1,
		Name:  "Kamil",
		Price: 1.223,
	}

	assert.NotNil(t, prd)
}

func TestModelConfig(t *testing.T) {
	config := Config{
		Server: struct {
			Addr string `yaml:"Addr"`
			Port int    `yaml:"Port"`
		}{
			Addr: "0.0.0.1",
			Port: 5001,
		},
		Database: struct {
			Host    string `yaml:"Host"`
			User    string `yaml:"User"`
			Pass    string `yaml:"Pass"`
			Port    string `yaml:"Port"`
			DBName  string `yaml:"DBName"`
			DBFlags string `yaml:"DBFlags"`
			Type    string `yaml:"Type"`
		}{
			Host:    "localhost",
			User:    "admin",
			Pass:    "password",
			Port:    "5432",
			DBName:  "mydb",
			DBFlags: "charset=utf8",
		},
		Jwt: struct {
			SecretKey string `yaml:"SecretKey"`
		}{
			SecretKey: "supersecretkey",
		},
		AppAdmin: struct {
			Enable bool   `yaml:"enable"`
			User   string `yaml:"enable"`
			Pass   string `yaml:"pass"`
		}{
			Enable: true,
			User:   "admin",
			Pass:   "adminpassword",
		},
	}

	assert.NotNil(t, config)
}

func TestModelToken(t *testing.T) {
	token := Token{
		Raw: "setrscx2341234d",
	}

	assert.NotNil(t, token)
}
