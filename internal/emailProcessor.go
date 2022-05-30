package internal

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strconv"

	"github.com/BreCkver/st-codeChallenge/models"
	gomail "gopkg.in/gomail.v2"
)

/*SendEmail */
func SendEmail(template string, destinatary string) error {

	m := gomail.NewMessage()
	m.SetHeader("From", "breckver.dll@gmail.com")
	m.SetHeader("To", destinatary)
	m.SetHeader("Subject", "Stori Test!")
	m.Embed("./images/firma.png")
	m.Embed("./images/header.png")
	m.Embed("./images/ico_fb_circ.png")
	m.Embed("./images/ico_in_circ.png")
	m.Embed("./images/ico_instagram_circ.png")
	m.Embed("./images/ico_tw_circ.png")
	m.Embed("./images/ico_you_circ.png")

	m.SetBody("text/html", template)

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "breckver.dll@gmail.com", "rqrqmxdszmeudxmh")

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error enviando email %v", err.Error())
		return err
	}
	return nil
}

func GetTemplateEmail(txList []*models.Transaction, accountName string) (string, *models.Summary, error) {

	debitAmount, debitAccount := transactionTypeCalculateAverage(txList, "debit")
	creditAmount, creditAccount := transactionTypeCalculateAverage(txList, "credit")
	totalBalance := transactionTotalBalance(txList)

	summary := models.Summary{
		AccountName:      accountName,
		TotalBalance:     fmt.Sprintf("%.2f", totalBalance),
		Debit:            fmt.Sprintf("%.2f", (debitAmount / debitAccount)),
		Credit:           fmt.Sprintf("%.2f", (creditAmount / creditAccount)),
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

func transactionTypeCalculateAverage(txList []*models.Transaction, typeTx string) (float64, float64) {

	var amount float64
	var count float64

	for i := 0; i < len(txList); i++ {
		if txList[i].Type == typeTx {
			count += 1
			amount += txList[i].Amount
		}
	}

	return amount, count
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
