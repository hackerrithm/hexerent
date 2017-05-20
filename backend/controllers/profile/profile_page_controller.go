package profile

import (
	"fmt"
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

// ProfilePage renders the profile.html file
func ProfilePage(w http.ResponseWriter, r *http.Request) {

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	sesString, _ := session.GlobalSession.Values["user"]
	sesString2, _ := session.GlobalSession.Values["firstTimeUser"]
	fmt.Println(sesString, " : THIS IS sesString")
	fmt.Println(sesString2, " : THIS IS sesString2")

	passedSession := SessionCart{
		sesString,
		sesString2,
	}

	fmt.Println(passedSession)

	config.Tpl.ExecuteTemplate(w, "profile.html", passedSession)
}
