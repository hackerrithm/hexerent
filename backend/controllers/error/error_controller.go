package error

import (
	"hexerent/backend/config"
	"net/http"
)

// ErrorHandler renders 404 page
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		config.Tpl.ExecuteTemplate(w, "login.html", nil)
	}

}
