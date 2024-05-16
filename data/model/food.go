package model

import (
	"database/sql"
	"encoding/json"
	"io"
)

// Foods is a list of Food
type Foods []Food

// FromJSON serializes data from json
func (o *Foods) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(o)
}

// ToJSON converts the collection to json
func (o *Foods) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}

// Food defines an food in the database
type Food struct {
	ID        int            `db:"id" json:"id,omitempty"`
	UserID    int            `db:"user_id" json:"-"`
	CreatedAt string         `db:"created_at" json:"-"`
	UpdatedAt string         `db:"updated_at" json:"-"`
	DeletedAt sql.NullString `db:"deleted_at" json:"-"`
	Items     []FoodItems    `json:"items,omitempty"`
}

// FromJSON serializes data from json
func (o *Food) FromJSON(data io.Reader) error {
	de := json.NewDecoder(data)
	return de.Decode(o)
}

// ToJSON converts the collection to json
func (o *Food) ToJSON() ([]byte, error) {
	return json.Marshal(o)
}

// FoodItems is an item/quantity in an food
type FoodItems struct {
	ID        int            `db:"id" json:"-"`
	FoodID    int            `db:"food_id" json:"-"`
	Name      string         `db:"name" json:"name,omitempty"`
	Price     float64        `db:"price" json:"price,omitempty"`
	CreatedAt string         `db:"created_at" json:"-"`
	UpdatedAt string         `db:"updated_at" json:"-"`
	DeletedAt sql.NullString `db:"deleted_at" json:"-"`
}
