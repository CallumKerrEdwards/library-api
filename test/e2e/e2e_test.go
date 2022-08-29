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
	"authors": [
		{
		"forenames": "Brandon",
		"sortName": "Sanderson"
		}
	],
	"title": "The Final Empire",
	"series": {
		"title": "Mistborn",
		"sequence": 1
	}
}
`
	putBody = `{
	"id": "%s",
	"authors": [
	  {
		"forenames": "Brandon",
		"sortName": "Sanderson"
	  }
	],
	"releaseDate": "2006-07-16",
	"description": "There is always another secret.",
	"title": "The Final Empire",
	"series": {
	  "title": "Mistborn",
	  "sequence": 1
	}
}
`
	expectedGetResponse1 = `{"id":"%s","title":"The Final Empire","authors":[{"forenames":"Brandon","sortName":"Sanderson"}],"series":{"sequence":1,"title":"Mistborn"}}
`
	expectedPutAndGetResponse2 = `{"id":"%s","title":"The Final Empire","authors":[{"forenames":"Brandon","sortName":"Sanderson"}],"description":"There is always another secret.","releaseDate":"2006-07-16","series":{"sequence":1,"title":"Mistborn"}}
`
	deleteResponse = `{"Message":"Successfully deleted book with ID %s"}
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
	assert.Equal(t, 201, resp.StatusCode())

	var post postResponse
	err = json.Unmarshal(resp.Body(), &post)
	assert.Nil(t, err)
	assert.NotEmpty(t, post.ID)

	// get newly created book
	resp, err = client.R().SetAuthToken(login.JWT).Get(apiHost + "/book/" + post.ID)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, fmt.Sprintf(expectedGetResponse1, post.ID), string(resp.Body()))

	// update created book
	resp, err = client.R().
		SetAuthToken(login.JWT).
		SetBody(fmt.Sprintf(putBody, post.ID)).
		Put(apiHost + "/book/" + post.ID)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, fmt.Sprintf(expectedPutAndGetResponse2, post.ID), string(resp.Body()))

	// get newly updated book
	resp, err = client.R().SetAuthToken(login.JWT).Get(apiHost + "/book/" + post.ID)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, fmt.Sprintf(expectedPutAndGetResponse2, post.ID), string(resp.Body()))

	// delete book
	resp, err = client.R().SetAuthToken(login.JWT).Delete(apiHost + "/book/" + post.ID)
	assert.Nil(t, err)
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, fmt.Sprintf(deleteResponse, post.ID), string(resp.Body()))

	// book no longer exists
	resp, err = client.R().SetAuthToken(login.JWT).Get(apiHost + "/book/" + post.ID)
	assert.Nil(t, err)
	assert.Equal(t, 404, resp.StatusCode())
}
