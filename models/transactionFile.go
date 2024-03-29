package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*TransactionsFile*/
type TransactionsFile struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Account         Account            `bson:"account" json:"account,omitempty"`
	ErrorList      []string           `bson:"errorlist" json:"errorlist,omitempy"`
	Summary        *Summary            `bson:"summary" json:"summary,omitempty"`
	TransactionList []*Transaction    `bson:"transactionlist" json:"transactionlist,omitempty"`
}
