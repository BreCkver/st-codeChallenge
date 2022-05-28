package internal

import (
	"bytes"
	"html/template"
	"log"

	gomail "gopkg.in/gomail.v2"
)

type Summary struct {
	AccountName      string
	TotalBalance     string
	Debit            string
	Credit           string
	TransationNumber map[string]string
}

/*SendEmail */
func SendEmail() {

	summary := Summary{
		AccountName:  "Alejandra y Giovanni",
		TotalBalance: "39.74",
		Debit:        "-15.38",
		Credit:       "35.25",
	}

	summary.TransationNumber = make(map[string]string)
	summary.TransationNumber["TxNumber1"] = "Number of transaction in July: 2"
	summary.TransationNumber["TxNumber2"] = "Number of transaction in August: 3"

	tmpl, err := template.ParseFiles("./templates/confirmation.html")
	if err != nil {
		log.Print("err 1")
		log.Print(err)
		return
	}

	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, summary); err != nil {
		log.Print("err 2")
		log.Print(err)
		return

	}

	result := tpl.String()
	m := gomail.NewMessage()
	m.SetHeader("From", "breckver.dll@outlook.com")
	m.SetHeader("To", "breckver.dll@gmail.com")
	m.SetHeader("Subject", "Stori Test!")
	m.Embed("./images/firma.png")
	m.Embed("./images/header.png")
	m.Embed("./images/ico_fb_circ.png")
	m.Embed("./images/ico_in_circ.png")
	m.Embed("./images/ico_instagram_circ.png")
	m.Embed("./images/ico_tw_circ.png")
	m.Embed("./images/ico_you_circ.png")

	m.SetBody("text/html", result)

	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "breckver.dll@gmail.com", "rqrqmxdszmeudxmh")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
