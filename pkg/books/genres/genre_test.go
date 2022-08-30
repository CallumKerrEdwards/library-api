package genres

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type genreHolder struct {
	Genres []Genre `json:"genres"`
}

const (
	sciFiFantasyJSONToUnmarshall = `{"genres": ["scifi", "Fantasy"]}`
	marshalledSciFiFantasyJSON   = `{"genres":["Science Fiction","Fantasy"]}`
)

var (
	sciFiFantasyHolder = genreHolder{Genres: []Genre{SciFi, Fantasy}}
)

func TestMarshallGenres(t *testing.T) {
	j, err := json.Marshal(sciFiFantasyHolder)
	assert.Nil(t, err)
	assert.Equal(t, marshalledSciFiFantasyJSON, string(j))
}

func TestUnmarshallGenres(t *testing.T) {
	var holder genreHolder
	err := json.Unmarshal([]byte(sciFiFantasyJSONToUnmarshall), &holder)
	assert.Nil(t, err)
	assert.Equal(t, sciFiFantasyHolder, holder)
}
