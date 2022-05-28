package models

import "time"

/*Transaction */
type Transaction struct {
	Identifier int32
	Date       time.Time
	Amount     float32
	Type       string
	Period     string
}
