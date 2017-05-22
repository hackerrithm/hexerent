package home

import (
	"hexerent/backend/config"
	"hexerent/backend/microservices/api/home"
	"hexerent/backend/session"
	"net/http"
)

// SessionCart group some session
type SessionCart struct {
	Sess1
	Sess2
	PostsLists
	PostInformationLikes
	CommentLists
}

// PostsLists to be used as interface
type PostsLists interface{}

// Sess1 to be used as interface
type Sess1 interface{}

// Sess2 to be used as interface
type Sess2 interface{}

// PostInformationLikes to be used as interface
type PostInformationLikes interface{}

// CommentLists to be used as interface
type CommentLists interface{}

// HomePage renders the home.html file
func HomePage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet || r.Method == "GET" {

		newList, totalLikesPerPost, newCommentList := home.GetAllPostInformation()

		sesString, _ := session.GlobalSession.Values["user"]
		sesString2, _ := session.GlobalSession.Values["firstTimeUser"]

		passedSession := SessionCart{
			sesString,
			sesString2,
			newList,
			totalLikesPerPost,
			newCommentList,
		}

		config.Tpl.ExecuteTemplate(w, "home.html", passedSession)
	} else if r.Method == http.MethodPost || r.Method == "POST" {

		home.CreateNewPost(w, r)

		http.Redirect(w, r, "/user/home", http.StatusSeeOther)

	}
}
