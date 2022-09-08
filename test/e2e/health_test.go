//go:build e2e
// +build e2e

package e2e

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	client := resty.New()
	apiHost := "http://localhost:8081/healthcheck"

	resp, err := client.R().
		Get(apiHost)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}

func TestReadinessCheck(t *testing.T) {
	client := resty.New()
	apiHost := "http://localhost:8081/readycheck"

	resp, err := client.R().
		Get(apiHost)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode())
}
