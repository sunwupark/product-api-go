package model

import (
	"database/sql"
	"encoding/json"
	"io"
)

type Friends []Friend

// FromJSON serializes data from json
func (c *Friends) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(c)
}

// ToJSON converts the collection to json
func (c *Friends) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}

// Coffee defines a coffee in the database
type Friend struct {
	ID          int            `db:"id" json:"id"`
	Name        string         `db:"name" json:"name"`
	Address     string         `db:"address" json:"address"`
	Description string         `db:"description" json:"description"`
	Image       string         `db:"image" json:"image"`
	CreatedAt   string         `db:"created_at" json:"-"`
	UpdatedAt   string         `db:"updated_at" json:"-"`
	DeletedAt   sql.NullString `db:"deleted_at" json:"-"`
}

func (c *Friend) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(c)
}

// ToJSON converts the collection to json
func (c *Friend) ToJSON() ([]byte, error) {
	return json.Marshal(c)
}
 