package login

import (
	"crypto/md5"
	"fmt"
	"hexerent/backend/config"
	"hexerent/backend/models"
	"hexerent/backend/session"
	"io"
	"net/http"
	"strconv"
	"time"
	//"golang.org/x/crypto/bcrypt"
)

// Login logs in a user
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) // r.Method used for getting the request method
	if r.Method == http.MethodGet || r.Method == "GET" {

		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		config.Tpl.ExecuteTemplate(w, "login.html", token)
	} else if r.Method == http.MethodPost || r.Method == "POST" {
		r.ParseForm()
		/*token := r.Form.Get("token")
		if token != "" {
			// check token validity
		} else {
			// give error if no token
		}*/

		fmt.Println("username: ", r.Form["username"])
		fmt.Println("password: ", r.Form["password"])

		userName := r.FormValue("username")
		password := r.FormValue("password")

		account, foundRecord := models.FindUser(userName, password)

		if foundRecord == true {
			// Get a session. We're ignoring the error resulted from decoding an
			// existing session: Get() always returns a session, even if empty.
			session.GlobalSession, _ = session.UniversalSessionStore.Get(r, "user-session")
			//TestSession, _ = UniversalSessionStore.Get(r, "user-session")
			// Set some session values.
			session.GlobalSession.Values["user"] = account //.UserName
			session.GlobalSession.Values["firstTimeUser"] = false
			session.GlobalSession.Values["isLoggedIn"] = false

			// Save it before we write to the response/return from the handler.
			session.GlobalSession.Save(r, w)

			http.Redirect(w, r, "/user/home", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}

	}
}

// Logout clears the session and logs the user out.
func Logout(w http.ResponseWriter, r *http.Request) {

	session.GlobalSession, _ = session.UniversalSessionStore.Get(r, "user-session")
	// If user is authenticated
	if session.GlobalSession.Values["isLoggedIn"] != false {
		session.GlobalSession.Values["isLoggedIn"] = false
	}

	session.GlobalSession.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
}
