package home

import (
	"fmt"
	"hexerent/backend/config"
	"hexerent/backend/models"
	"hexerent/backend/session"
	"net/http"
)

// SessionCart group some session
type SessionCart struct {
	Sess1      //sessionVar1 string
	Sess2      //sessionVar2 bool
	PostsLists //[]models.PublicHomePost
}

// PostsLists to be used as interface
type PostsLists interface{}

// Sess1 to be used as interface
type Sess1 interface{}

// Sess2 to be used as interface
type Sess2 interface{}

// HomePage renders the home.html file
func HomePage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet || r.Method == "GET" {

		var postsLists = models.FindAllHomeFeedPosts()

		var v string //models.PublicHomePost

		//var listHomePostFeed []models.PublicHomePost

		var listHomePostFeed = []string{}

		for _, v = range postsLists {

			listHomePostFeed = append(listHomePostFeed, v)

		}

		// Get a session. We're ignoring the error resulted from decoding an
		// existing session: Get() always returns a session, even if empty.

		//sesString, _ := login.TestSession.Values["user"]
		sesString, _ := session.GlobalSession.Values["user"]
		sesString2, _ := session.GlobalSession.Values["firstTimeUser"]
		fmt.Println(sesString, " : THIS IS sesString")
		fmt.Println(sesString2, " : THIS IS sesString2")
		fmt.Println(listHomePostFeed, " : THIS IS sesString2")

		passedSession := SessionCart{
			sesString,
			sesString2,
			listHomePostFeed,
		}

		config.Tpl.ExecuteTemplate(w, "home.html", passedSession)
	} else if r.Method == http.MethodPost || r.Method == "POST" {

		comment := r.FormValue("homeFeedPost")

		// Function call
		postInserter := models.InsertNewHomePost(comment)

		fmt.Println(postInserter)

		//Redirects to home page
		http.Redirect(w, r, "/user/home", http.StatusSeeOther)

	}
}
