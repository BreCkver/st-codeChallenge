package routers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/BreCkver/st-codeChallenge/internal"
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
		m := make(map[string]int)

		for _, tx := range transactionList {
			m[tx.Period] += 1
		}

		var amount float32
		var count int32

		for i := 0; i < len(transactionList); i++ {
			if transactionList[i].Type == "credit" {
				count += 1
				amount += transactionList[i].Amount
			}
		}

		log.Printf("Total tx %v", count)
		log.Printf("Monto tx %v", amount)

		internal.SendEmail()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(m)
	}

}
