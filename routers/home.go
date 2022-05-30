package routers

import (
	"log"
	"net/http"

	"github.com/BreCkver/st-codeChallenge/data"
	"github.com/BreCkver/st-codeChallenge/internal"
	"github.com/BreCkver/st-codeChallenge/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	Render(w, "./templates/home.html", nil)
}

func LoadFile(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Ocurrio al initicializar "+err.Error(), http.StatusInternalServerError)
		return
	}
	file, handler, err := r.FormFile("q")

	if err != nil {
		http.Error(w, "Ocurrio al leer"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	transactionFile := &internal.Request{
		Email:    r.PostFormValue("email"),
		File:     handler.Filename,
		FileName: "./uploads/" + handler.Filename,
	}

	if !transactionFile.Validate() {
		Render(w, "./templates/home.html", transactionFile)
	}

	if err := transactionFile.LoadFile(file); err != nil {
		transactionFile.Errors["File"] = err.Error()
		Render(w, "./templates/home.html", transactionFile)
		return
	}

	if transactionList, err := transactionFile.ReadFile(); err != nil {
		transactionFile.Errors["File"] = err.Error()
		Render(w, "./templates/home.html", transactionFile)
		return
	} else {

		template, summary, err := internal.GetTemplateEmail(transactionList, "Jaime")
		if err != nil {
			transactionFile.Errors["File"] = err.Error()
			Render(w, "./templates/home.html", transactionFile)
			return
		}

		errEmail := internal.SendEmail(template, transactionFile.Email)

		if errEmail != nil {
			log.Printf("Error %v", errEmail.Error())
			transactionFile.Errors["File"] = errEmail.Error()
			Render(w, "./templates/home.html", transactionFile)
			return
		}

		var final = models.TransactionsFile{
			Account:         models.Account{UserName: "Jaime", Email: transactionFile.Email},
			Summary:         summary,
			TransactionList: transactionList,
			ErrorList:       []string{"Error 01", "Error 02"},
		}

		id, err := data.TransacionFileSave(&final)
		if err != nil {
			transactionFile.Errors["File"] = err.Error()
			Render(w, "./templates/home.html", transactionFile)
			return
		}

		http.Redirect(w, r, "/confirmation?id="+id, http.StatusSeeOther)

	}

}
