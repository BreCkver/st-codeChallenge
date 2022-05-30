package routers

import (
	"log"
	"net/http"

	"github.com/BreCkver/st-codeChallenge/data"
	"github.com/BreCkver/st-codeChallenge/internal"
	"github.com/BreCkver/st-codeChallenge/models"
)

const (
	templateHome = "./templates/home.html"
	uploadFolder = "./uploads/"
)

func Index(w http.ResponseWriter, r *http.Request) {
	Render(w, templateHome, nil)
}

func LoadFile(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Ocurrio al initicializar "+err.Error(), http.StatusInternalServerError)
		return
	}
	file, handler, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "Ocurrio al leer"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	transactionFile := &internal.Request{
		Email:       r.PostFormValue("email"),
		AccountName: r.PostFormValue("accountName"),
		FileName:    uploadFolder + handler.Filename,
	}

	if !transactionFile.Validate() {
		Render(w, templateHome, transactionFile)
	}

	if err := transactionFile.LoadFile(file); err != nil {
		transactionFile.Errors["File"] = err.Error()
		Render(w, templateHome, transactionFile)
		return
	}

	if transactionList, err := transactionFile.ReadFile(); err != nil {
		transactionFile.Errors["File"] = err.Error()
		Render(w, templateHome, transactionFile)
		return
	} else {

		template, summary, err := internal.GetTemplateEmail(transactionList, transactionFile.AccountName)
		if err != nil {
			transactionFile.Errors["File"] = err.Error()
			Render(w, templateHome, transactionFile)
			return
		}

		errEmail := internal.SendEmail(template, transactionFile.Email)

		if errEmail != nil {
			log.Printf("Error %v", errEmail.Error())
			transactionFile.Errors["File"] = errEmail.Error()
			Render(w, templateHome, transactionFile)
			return
		}

		var final = models.TransactionsFile{
			Account:         models.Account{UserName: transactionFile.AccountName, Email: transactionFile.Email},
			Summary:         summary,
			TransactionList: transactionList,
			ErrorList:       []string{"Error 01", "Error 02"},
		}

		id, err := data.TransacionFileSave(&final)
		if err != nil {
			transactionFile.Errors["File"] = err.Error()
			Render(w, templateHome, transactionFile)
			return
		}

		http.Redirect(w, r, "/confirmation?id="+id, http.StatusSeeOther)

	}

}
