package home

import (
	"bytes"
	"fmt"
	"hexerent/backend/config"
	"hexerent/backend/models"
	"hexerent/backend/session"
	"net/http"
	"time"
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

// GetUserInformation stores user data
func GetUserInformation() uint64 {
	sesString, _ := session.GlobalSession.Values["userID"]
	naswer := sesString.(uint64)
	return naswer
}

// HomePage renders the home.html file
func HomePage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet || r.Method == "GET" {

		var postsLists = models.FindAllHomeFeedPosts()
		var v models.PublicHomePost
		var listHomePostFeed []models.PublicHomePost
		var newList []models.PublicHomePost

		var postCounter int64

		for _, v = range postsLists {

			listHomePostFeed = append(listHomePostFeed, v)
			postCounter++
			fmt.Println("postCounter is: ", postCounter)
		}

		postCounter = (postCounter - 1)

		var k int64
		k = 0

		for i := postCounter; i > -1; i-- {
			newList = append(newList, listHomePostFeed[i])
			k++
		}

		sesString, _ := session.GlobalSession.Values["user"]
		sesString2, _ := session.GlobalSession.Values["firstTimeUser"]

		passedSession := SessionCart{
			sesString,
			sesString2,
			newList,
		}

		config.Tpl.ExecuteTemplate(w, "home.html", passedSession)
	} else if r.Method == http.MethodPost || r.Method == "POST" {

		author := GetUserInformation()
		topic := "Random"           //r.FormValue("homeFeedPost")
		category := "Entertainment" //r.FormValue("homeFeedPost")
		content := r.FormValue("homeFeedPost")
		datePosted := time.Now()  //r.FormValue("homeFeedPost")
		var likes uint64 = 134    //r.FormValue("homeFeedPost")
		var upvotes uint64 = 134  //r.FormValue("homeFeedPost")
		var downvotes uint64 = 12 //r.FormValue("homeFeedPost")

		createdTime := FormatDateTime(datePosted)

		postInserter := models.InsertNewHomePost(author, topic, category, content, createdTime, likes, upvotes, downvotes)
		fmt.Println(postInserter)

		//Redirects to home page
		http.Redirect(w, r, "/user/home", http.StatusSeeOther)

	}
}

// FormatDateTime used to format date and time
func FormatDateTime(t time.Time) string {
	var buffer bytes.Buffer
	buffer.WriteString(t.Month().String()[:3])
	buffer.WriteString(fmt.Sprintf(" %2d '%2d at %2d:%2d", t.Day(), t.Year()%100, t.Hour(), t.Minute()))
	return buffer.String()
}
