package register

import (
	"fmt"
	"hexerent/backend/config"
	"hexerent/backend/models"
	"hexerent/backend/session"
	"log"
	"net/http"
	"regexp"
)

// RegisterValidator validates Register
func RegisterValidator(w http.ResponseWriter, r *http.Request, isValidInformation bool) bool {
	r.ParseForm()

	// Ensures all form fields are filled out
	if len(r.Form.Get("firstname")) == 0 || len(r.Form.Get("lastname")) == 0 || len(r.Form.Get("username")) == 0 || len(r.Form.Get("password")) == 0 || len(r.Form.Get("email")) == 0 {
		fmt.Println("form not filled out.... too short")
		isValidInformation = false
		fmt.Println(isValidInformation, " for everything")
	}

	// Checks to ensure username has at least three (3) characters
	if len(r.Form.Get("username")) < 3 {
		fmt.Println("name too short")
		isValidInformation = false
		fmt.Println(isValidInformation, " for name")
	}

	// checks email to ensure validity
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.FormValue("email")); !m {
		fmt.Println("incorrect email address format")
		isValidInformation = false
		fmt.Println(isValidInformation, " for email")
	}

	var a = r.Form["subscribe"]
	isSubscribedTicked := true
	if len(a) == 0 {
		fmt.Println(!isSubscribedTicked)
	}

	slice := []string{"1", "2", "3"}

	for _, v := range slice {
		if v == r.Form.Get("usertype") {
			fmt.Println(slice)
		}
	}

	return isValidInformation
}

// Register used to registers a new user
func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) // r.Method used for getting the request method
	if r.Method == http.MethodGet || r.Method == "GET" {
		// Just display the page for registration
		config.Tpl2.ExecuteTemplate(w, "index.html", nil)

	} else if r.Method == http.MethodPost || r.Method == "POST" {

		var validatedUserData = true

		firstName := r.FormValue("firstname")
		lastName := r.FormValue("lastname")
		userName := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")

		//Validator
		goAhead := RegisterValidator(w, r, validatedUserData)

		fmt.Println("this is goAhead: ", goAhead)

		if goAhead == true {
			// Function call
			userInserter := models.InsertNewUser(userName, firstName, lastName, password, email)

			// Get a session. We're ignoring the error resulted from decoding an
			// existing session: Get() always returns a session, even if empty.
			session.GlobalSession, _ = session.UniversalSessionStore.Get(r, "user-session")
			// Set some session values.
			session.GlobalSession.Values["user"] = userInserter //.UserName
			session.GlobalSession.Values["firstTimeUser"] = true
			// Save it before we write to the response/return from the handler.
			session.GlobalSession.Save(r, w)

			//Redirects to home page
			http.Redirect(w, r, "/user/home", http.StatusSeeOther)
		}
		//goAhead = false
		config.Tpl2.ExecuteTemplate(w, "index.html", nil)

	} else {
		log.Fatalln("dangerous things are going on")
	}

}
