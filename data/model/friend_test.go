package model

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFriendDeserializeFromJSON(t *testing.T) {
	c := Friends{}

	err := c.FromJSON(bytes.NewReader([]byte(friendsData)))
	assert.NoError(t, err)

	assert.Len(t, c, 2)
	assert.Equal(t, 1, c[0].ID)
	assert.Equal(t, 2, c[1].ID)
}

func TestFriendsSerializesToJSON(t *testing.T) {
	c := Friends{
		Friend{ID: 1, Name: "test", Address: "test"},
	}

	d, err := c.ToJSON()
	assert.NoError(t, err)

	cd := make([]map[string]interface{}, 0)
	err = json.Unmarshal(d, &cd)
	assert.NoError(t, err)

	assert.Equal(t, float64(1), cd[0]["id"])
	assert.Equal(t, "test", cd[0]["name"])
	assert.Equal(t, float64(120.12), cd[0]["address"])
}

var friendsData = `
[
	{
		"id": 1,
		"name": "suwnu",
		"address": "seoul"
	},
	{
		"id": 2,
		"name": "won",
		"address": "sunwon"
	}
]
`