package routers

import (
	"net/http"

	"github.com/BreCkver/st-codeChallenge/internal"
)

/*Confirmation load information about specific identifier */
func Confirmation(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(w, "The Id is mandatory", http.StatusBadRequest)
		return
	}

	summary, err := internal.GetTrasaction(id)
	if err != nil {
		http.Error(w, "Sorry, something went wrong 03"+err.Error(), http.StatusInternalServerError)
		return
	}

	Render(w, "./templates/confirmation.html", summary)
}
