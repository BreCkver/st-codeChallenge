package routers

import "net/http"

func Confirmation(w http.ResponseWriter, r *http.Request) {
	Render(w, "./templates/confirmation.html", nil)
}
