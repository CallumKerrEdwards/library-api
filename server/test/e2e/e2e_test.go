package e2e

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

var (
	postBody = `{
		"author": "Brandon Sanderson",
		"title": "The Final Empire",
		"series": {
			"title": "Mistborn",
			"sequence": 1
		}
	}
`
	expectedGetBody = `{"id":"%s","title":"The Final Empire","author":"Brandon Sanderson","series":{"sequence":1,"title":"Mistborn"}}
`
)

type loginResponse struct {
	JWT string `json:"jwt"`
}

type postResponse struct {
	ID string `json:"id"`
}

func TestEndToEndWorkflow(t *testing.T) {
	client := resty.New()
	apiHost := "http://localhost:8081/api/v1"

	// login
	resp, err := client.R().Get(apiHost + "/auth/login")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	var login loginResponse
	err = json.Unmarshal(resp.Body(), &login)
	assert.Nil(t, err)
	assert.NotEmpty(t, login.JWT)

	// create new book
	resp, err = client.R().
		SetAuthToken(login.JWT).
		SetBody(postBody).
		Post(apiHost + "/book")
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())

	var post postResponse
	err = json.Unmarshal(resp.Body(), &post)
	assert.Nil(t, err)
	assert.NotEmpty(t, post.ID)

	// get newly created book
	resp, err = client.R().SetAuthToken(login.JWT).Get(apiHost + "/book/" + post.ID)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, fmt.Sprintf(expectedGetBody, post.ID), string(resp.Body()))

}
