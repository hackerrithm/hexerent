package profile

import (
	"fmt"
	"hexerent/backend/config"
	"hexerent/backend/microservices/api/profile"
	"hexerent/backend/session"
	"net/http"
)

// SessionCart group some session
type SessionCart struct {
	Sess1
	Sess2
	Profile
}

// Sess1 to be used as interface
type Sess1 interface{}

// Sess2 to be used as interface
type Sess2 interface{}

// Profile to be used as interface
type Profile interface{}

// ProfilePage renders the profile.html file
func ProfilePage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet || r.Method == "GET" {
		// Get a session. We're ignoring the error resulted from decoding an
		// existing session: Get() always returns a session, even if empty.
		sesString, _ := session.GlobalSession.Values["user"]
		sesString2, _ := session.GlobalSession.Values["firstTimeUser"]
		fmt.Println(sesString, " : THIS IS sesString")
		fmt.Println(sesString2, " : THIS IS sesString2")
		profileData := profile.GetProfileInformation()
		fmt.Println("this is profile struct: ", profileData)

		passedSession := SessionCart{
			sesString,
			sesString2,
			profileData,
		}

		fmt.Println(passedSession)

		config.Tpl.ExecuteTemplate(w, "profile.html", passedSession)

	} else if r.Method == http.MethodPost || r.Method == "POST" {

	}

}
