package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewYamlConfigProvider(t *testing.T) {
	path := `/Users/xuzequn/OperatingPlatform/Learning/go进阶训练营/golearning/grammar/sparrow/test/client.yaml`
	ycp, err := NewYamlConfigProvider(path)
	assert.Nil(t, err)

	cfg, err := ycp.GetServiceConfig("hello")
	assert.Nil(t, err)
	assert.Equal(t, "http://127.0.0.1:8080", cfg.Endpoint)
}
