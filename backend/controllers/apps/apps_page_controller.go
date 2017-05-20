package apps

import (
	"hexerent/backend/config"
	"hexerent/backend/session"
	"net/http"
)

// SessionCart group some session
type SessionCart struct {
	Sess1 //sessionVar1 string
	Sess2 //sessionVar2 bool
}

type Sess1 interface{}
type Sess2 interface{}

// AppsPage renders the apps-page.html file
func AppsPage(w http.ResponseWriter, r *http.Request) {

	sesString, _ := session.GlobalSession.Values["user"]
	sesString2, _ := session.GlobalSession.Values["firstTimeUser"]

	passedSession := SessionCart{
		sesString,
		sesString2,
	}

	config.Tpl.ExecuteTemplate(w, "apps-page.html", passedSession)
}
