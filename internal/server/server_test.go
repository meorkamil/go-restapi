package server

import (
	"go-restapi/internal/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAPIServer(t *testing.T) {
	// Just invoke new NewAPIServer
	file := "../../config/config.yml"
	config := util.ConfigInit(file)
	server := NewAPIServer(config)
	assert.NotNil(t, server)
}
