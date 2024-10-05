package util

import (
	"encoding/json"
	"fmt"
	"go-restapi/internal/model"
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

func Decode[T any](r *http.Request) (T, error) {

	var v T

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}

	return v, nil
}

func Encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {

		return fmt.Errorf("encode json: %w", err)

	}

	return nil
}

func ConfigInit(c string) *model.Config {

	viper.SetConfigType("yaml")
	viper.SetConfigFile(c)

	log.Println("Loaded configuration:", c)

	if err := viper.ReadInConfig(); err != nil {

		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	var config model.Config

	if err := viper.Unmarshal(&config); err != nil {

		log.Fatal(err)
		os.Exit(1)

	}

	return &config

}
