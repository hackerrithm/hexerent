package apps

import (
	"hexerent/backend/config"
	"hexerent/backend/session"
	"net/http"
)

// AppsPage renders the apps-page.html file
func AppsPage(w http.ResponseWriter, r *http.Request) {

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.

	sesString, _ := session.GlobalSession.Values["user"]

	config.Tpl.ExecuteTemplate(w, "apps-page.html", sesString)
}
