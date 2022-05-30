package models

import "time"

/*Transaction */
type Transaction struct {
	Identifier int32     `bson:"identifier" json:"identifier,omitempty"`
	Date       time.Time `bson:"date" json:"date,omitempty"`
	Amount     float64   `bson:"amount" json:"amount,omitempty"`
	Type       string    `bson:"type" json:"type,omitempty"`
}
