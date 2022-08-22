//go:build e2e
// +build e2e

package e2e

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheckEndpoint(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/healthcheck")
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
}
