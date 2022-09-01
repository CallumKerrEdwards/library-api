package books

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	genres "github.com/CallumKerrEdwards/library-api/pkg/books/genres"
	"github.com/CallumKerrEdwards/library-api/pkg/books/text"
)

const (
	jsonBook = `{
	"id": "",
	"title": "The Way of Kings",
	"authors": [
		{
			"forenames": "Brandon",
			"sortName": "Sanderson"
		}
	],
	"description": {
		"text": "I long for the days before the Last Desolation.\nThe age before the Heralds abandoned us and the Knights Radiant turned against us. A time when there was still magic in the world and honor in the hearts of men.\nhe world became ours, and we lost it. Nothing, it appears, is more challenging to the souls of men than victory itself.\nOr was that victory an illusion all along? Did our enemies realize that the harder they fought, the stronger we resisted? Perhaps they saw that the heat and the hammer only make for a better grade of sword. But ignore the steel long enough, and it will eventually rust away.\nThere are four whom we watch. The first is the surgeon, forced to put aside healing to become a soldier in the most brutal war of our time. The second is the assassin, a murderer who weeps as he kills. The third is the liar, a young woman who wears a scholar's mantle over the heart of a thief. The last is the highprince, a warlord whose eyes have opened to the past as his thirst for battle wanes.\nThe world can change. Surgebinding and Shardwielding can return; the magics of ancient days can become ours again. These four people are key.\nOne of them may redeem us.\nAnd one of them will destroy us.",
		"format": "Plain"
	},
	"releaseDate": "2010-08-31",
	"genres": [
		"Fantasy"
	],
	"series": {
		"sequence": 1,
		"title": "The Stormlight Archive"
	},
	"audiobook": {
		"audiobookMediaId": "0001",
		"narrators": [
			{
				"forenames": "Michael",
				"sortName": "Kramer"
			},
			{
				"forenames": "Kate",
				"sortName": "Reading"
			}
		],
		"coverImageMediaId": "0002"
	}
}`
)

func TestMarshalBookJSON(t *testing.T) {
	releaseDate, err := NewReleaseDate("2010-08-31")
	if err != nil {
		assert.Nil(t, err)
	}

	twok := Book{
		Title:   "The Way of Kings",
		Authors: []Person{{Forenames: "Brandon", SortName: "Sanderson"}},
		Description: &Description{Text: `I long for the days before the Last Desolation.
The age before the Heralds abandoned us and the Knights Radiant turned against us. A time when there was still magic in the world and honor in the hearts of men.
he world became ours, and we lost it. Nothing, it appears, is more challenging to the souls of men than victory itself.
Or was that victory an illusion all along? Did our enemies realize that the harder they fought, the stronger we resisted? Perhaps they saw that the heat and the hammer only make for a better grade of sword. But ignore the steel long enough, and it will eventually rust away.
There are four whom we watch. The first is the surgeon, forced to put aside healing to become a soldier in the most brutal war of our time. The second is the assassin, a murderer who weeps as he kills. The third is the liar, a young woman who wears a scholar's mantle over the heart of a thief. The last is the highprince, a warlord whose eyes have opened to the past as his thirst for battle wanes.
The world can change. Surgebinding and Shardwielding can return; the magics of ancient days can become ours again. These four people are key.
One of them may redeem us.
And one of them will destroy us.`, Format: text.Plain},
		ReleaseDate: releaseDate,
		Genres:      []genres.Genre{genres.Fantasy},
		Series:      Series{Title: "The Stormlight Archive", Sequence: 1},
		Audiobook: &Audiobook{
			AudiobookMediaID:  "0001",
			CoverImageMediaID: "0002",
			Narrators: []Person{
				{
					Forenames: "Michael",
					SortName:  "Kramer",
				},
				{
					Forenames: "Kate",
					SortName:  "Reading",
				},
			},
		},
	}

	generatedJSON, err := json.MarshalIndent(twok, "", "\t")
	assert.Nil(t, err)
	assert.Equal(t, jsonBook, string(generatedJSON))
}

var (
	personsTests = []struct {
		title    string
		persons  []Person
		expected string
	}{
		{
			title:    "no persons",
			persons:  []Person{},
			expected: "",
		},
		{
			title: "one person",
			persons: []Person{
				{
					Forenames: "R. F.",
					SortName:  "Kuang",
				},
			},
			expected: "R. F. Kuang",
		},
		{
			title: "two persons",
			persons: []Person{
				{
					Forenames: "Amal",
					SortName:  "El-Mohtar",
				},
				{
					Forenames: "Max",
					SortName:  "Gladstone",
				},
			},
			expected: "Amal El-Mohtar & Max Gladstone",
		},
		{
			title: "three persons",
			persons: []Person{
				{
					Forenames: "Roger",
					SortName:  "Clark",
				},
				{
					Forenames: "Jay",
					SortName:  "Snyder",
				},
				{
					Forenames: "Elizabeth",
					SortName:  "Evans",
				},
			},
			expected: "Roger Clark, Jay Snyder & Elizabeth Evans",
		},
	}
)

func TestGetAuthors(t *testing.T) {
	for _, test := range personsTests {
		t.Run(test.title, func(t *testing.T) {
			assert.Equal(t, test.expected, GetPersonsString(test.persons))
		})
	}
}
