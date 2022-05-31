package internal

import (
	"github.com/BreCkver/st-codeChallenge/data"
	"github.com/BreCkver/st-codeChallenge/models"
)

func TransactionSave(txList []*models.Transaction, account *models.Account, errList []string) (string, error) {

	template, summary, err := GetTemplateEmail(txList, account.UserName)
	if err != nil {
		return "", err
	}

	var transactionsFile = models.TransactionsFile{
		Account:         *account,
		Summary:         summary,
		TransactionList: txList,
		ErrorList:       errList,
	}

	id, err := data.TransacionFileSave(&transactionsFile)
	if err != nil {
		return "", err
	}

	errEmail := SendEmail(template, account.Email)
	if errEmail != nil {
		return "", errEmail
	}

	return id, nil
}

func GetTrasaction(id string) (*models.Summary, error) {

	txFile, err := data.GetTransactionFile(id)
	if err != nil {
		return nil, err
	}

	var summary = txFile.Summary
	return summary, nil
}
