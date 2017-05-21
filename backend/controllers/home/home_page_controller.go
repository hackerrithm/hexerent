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
		var commentLists = models.FindAllComments()

		var v models.PublicHomePost
		var q models.Comment

		var listHomePostFeed []models.PublicHomePost
		var listComment []models.Comment

		var newList []models.PublicHomePost
		var newCommentList []models.Comment

		var postCounter, commentCounter int64

		for _, v = range postsLists {

			listHomePostFeed = append(listHomePostFeed, v)
			postCounter++
			fmt.Println("postCounter is: ", postCounter)
		}

		for _, q = range commentLists {

			listComment = append(listComment, q)
			commentCounter++
			fmt.Println("commentCounter is: ", commentCounter)
		}

		postCounter = (postCounter - 1)
		commentCounter = (commentCounter - 1)

		var k int64
		k = 0
		var totalLikesPerPost, totalUpvotesPerPost, totalDownvotesPerPost, totalCommentsPerPost uint64
		for i := postCounter; i > -1; i-- {
			newList = append(newList, listHomePostFeed[i])
			totalLikesPerPost = listHomePostFeed[i].Likes
			totalUpvotesPerPost = listHomePostFeed[i].Upvotes
			totalDownvotesPerPost = listHomePostFeed[i].Downvotes
			totalCommentsPerPost = listHomePostFeed[i].Comments

			fmt.Println("Number of likes: ", totalLikesPerPost,
				"Number of upvotes: ", totalUpvotesPerPost,
				"Number of downvotes: ", totalDownvotesPerPost,
				"Number of comments: ", totalCommentsPerPost)
			k++
		}

		//var userComment []string
		//userComment := make([]string, 0)
		for i := commentCounter; i > -1; i-- {
			newCommentList = append(newCommentList, listComment[i])
			fmt.Println("This is the comment: ", listComment[i].CommentText)
			//userComment[i] = listComment[i].CommentText
			k++
		}

		/*var index uint64
		var toatalLikesPerPost uint64
		for index = 0; index < uint64(len(newList)); index++ {
			listHomePostFeed[index] = models.FindHomeFeedPost(index + 1)
			toatalLikesPerPost = listHomePostFeed[index].Likes
			fmt.Println("Number of likes: ", toatalLikesPerPost)
		}*/

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

		author := GetUserInformation()
		topic := "Random"           //r.FormValue("homeFeedPost")
		category := "Entertainment" //r.FormValue("homeFeedPost")
		content := r.FormValue("homeFeedPost")
		datePosted := time.Now()
		var likes uint64
		var upvotes uint64
		var downvotes uint64
		var comments uint64

		createdTime := FormatDateTime(datePosted)

		postInserter := models.InsertNewHomePost(author, topic, category, content, createdTime, likes, upvotes, downvotes, comments)
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
