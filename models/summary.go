package models

/*Summary */
type Summary struct {
	AccountName      string `bson:"acountName" json:"acountName,omitempty"`
	TotalBalance     string `bson:"totalBalance" json:"totalBalance,omitempty"`
	Debit            string `bson:"debit" json:"debit,omitempty"`
	Credit           string `bson:"credit" json:"credit,omitempty"`
	TransationNumber map[string]string
}
