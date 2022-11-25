package internal

import (
	"bytes"
	"fmt"
	"html/template"
	"strconv"

	"github.com/BreCkver/st-codeChallenge/models"
	gomail "gopkg.in/gomail.v2"
)

/*SendEmail */
func SendEmail(template string, destinatary string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", "breckver.dll@gmail.com")
	m.SetHeader("To", destinatary)
	m.SetHeader("Subject", "Stori - transaction summary!")
	m.Embed("./images/firma.png")
	m.Embed("./images/header.png")
	m.Embed("./images/ico_fb_circ.png")
	m.Embed("./images/ico_in_circ.png")
	m.Embed("./images/ico_instagram_circ.png")
	m.Embed("./images/ico_tw_circ.png")
	m.Embed("./images/ico_you_circ.png")

	m.SetBody("text/html", template)

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "", "")

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

/*GetTemplateEmail to send in the email */
func GetTemplateEmail(txList []*models.Transaction, accountName string) (string, *models.Summary, error) {

	debitAverage := transactionTypeCalculateAverage(txList, "debit")
	creditAverage := transactionTypeCalculateAverage(txList, "credit")
	totalBalance := transactionTotalBalance(txList)

	summary := models.Summary{
		AccountName:      accountName,
		TotalBalance:     fmt.Sprintf("%.2f", totalBalance),
		Debit:            fmt.Sprintf("%.2f", debitAverage),
		Credit:           fmt.Sprintf("%.2f", creditAverage),
		TransationNumber: map[string]string{},
	}

	summary.TransationNumber = make(map[string]string)
	txbyMonth := transactionNumberByMonth(txList)

	for month, quantity := range txbyMonth {
		summary.TransationNumber[month] = "Number of transaction in " + month + ": " + strconv.Itoa(quantity)
	}

	tmpl, err := template.ParseFiles("./templates/emailTemplate.html")
	if err != nil {
		return "", nil, err
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, summary); err != nil {
		return "", nil, err

	}

	return tpl.String(), &summary, nil
}

func transactionTypeCalculateAverage(txList []*models.Transaction, typeTx string) float64 {

	var amount float64 = 0
	var count float64 = 0

	for i := 0; i < len(txList); i++ {
		if txList[i].Type == typeTx {
			count += 1
			amount += txList[i].Amount
		}
	}

	if count == 0 {
		return 0
	}

	return amount / count
}

func transactionTotalBalance(txList []*models.Transaction) float64 {

	var amount float64
	for i := 0; i < len(txList); i++ {
		amount += txList[i].Amount
	}

	return amount
}

func transactionNumberByMonth(txList []*models.Transaction) map[string]int {

	monthArray := make(map[string]int)
	for _, i := range txList {
		month := i.Date.Month().String()
		monthArray[month] += 1
	}
	return monthArray
}
