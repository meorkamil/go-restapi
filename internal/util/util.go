package util

import (
	"encoding/json"
	"fmt"
	"go-restapi/internal/model"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

var lg = NewLogger()

// JSON Decoding
func Decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}

	return v, nil
}

// JSON Encoding
func Encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}

	return nil
}

// Set configuration
func ConfigInit(c string) *model.Config {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(c)

	lg.Info("Configuration:", c)
	if err := viper.ReadInConfig(); err != nil {
		lg.Fatal("fatal error config file: default \n", err)
		os.Exit(1)
	}

	var config model.Config
	if err := viper.Unmarshal(&config); err != nil {
		lg.Fatal(err)
		os.Exit(1)
	}
	// Set JWT Token in environment variable
	os.Setenv("SECRETKEY", config.Jwt.SecretKey)

	// Return all config in struct
	return &config
}
