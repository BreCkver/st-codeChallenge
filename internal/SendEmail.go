package internal

import (
	gomail "gopkg.in/gomail.v2"
)

/*SendEmail */
func SendEmail() {
	m := gomail.NewMessage()
	m.SetHeader("From", "breckver.dll@outlook.com")
	m.SetHeader("To", "breckver.dll@gmail.com")
	m.SetHeader("Subject", "Stori Test!")
	m.Embed("./images/image.jpg")
	m.SetBody("text/html", `<img src="cid:image.jpg" alt="My image" />`)

	//d := gomail.NewPlainDialer("smtp-mail.outlook.com", 587, "breckver.dll@outlook.com", "mexico40W$W$")
	d := gomail.NewPlainDialer("smtp.gmail.com", 587, "breckver.dll@gmail.com", "rqrqmxdszmeudxmh")

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
