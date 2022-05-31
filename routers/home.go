package routers

import (
	"log"
	"net/http"
	"strings"

	"github.com/BreCkver/st-codeChallenge/internal"
	"github.com/BreCkver/st-codeChallenge/models"
)

const (
	templateHome = "./templates/home.html"
	uploadFolder = "./uploads/"
)

/*Index Load initial page*/
func Index(w http.ResponseWriter, r *http.Request) {
	Render(w, templateHome, nil)
}

/*LoadFile handler when the user clic in the buttom submit */
func LoadFile(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Sorry, something went wrong 01 "+err.Error(), http.StatusInternalServerError)
		return
	}
	file, handler, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "Sorry, something went wrong 02"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	trxFile := &internal.Request{
		Account:  models.Account{UserName: r.PostFormValue("accountName"), Email: r.PostFormValue("email")},
		FileName: uploadFolder + handler.Filename,
	}

	if !trxFile.Validate() {
		Render(w, templateHome, trxFile)
	}

	if err := trxFile.LoadFile(file); err != nil {
		trxFile.Errors["File"] = err.Error()
		Render(w, templateHome, trxFile)
		return
	}

	if trxList, errList := trxFile.ReadFile(); trxList == nil {
		log.Printf("Errors: %v", strings.Join(errList, " - "))
		trxFile.Errors["File"] = strings.Join(errList, " - ")
		Render(w, templateHome, trxFile)
		return
	} else {

		identifier, err := internal.TransactionSave(trxList, &trxFile.Account, errList)

		if err != nil {
			trxFile.Errors["File"] = err.Error()
			Render(w, templateHome, trxFile)
			return
		}

		http.Redirect(w, r, "/confirmation?id="+identifier, http.StatusSeeOther)
	}
}
