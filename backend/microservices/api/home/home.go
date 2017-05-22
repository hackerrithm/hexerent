package home

import (
	"bytes"
	"fmt"
	"hexerent/backend/models"
	"hexerent/backend/session"
	"net/http"
	"time"
)

// FormatDateTime used to format date and time
func FormatDateTime(t time.Time) string {
	var buffer bytes.Buffer
	buffer.WriteString(t.Month().String()[:3])
	buffer.WriteString(fmt.Sprintf(" %2d '%2d at %2d:%2d", t.Day(), t.Year()%100, t.Hour(), t.Minute()))
	return buffer.String()
}

// GetUserIdentification stores user data
func GetUserIdentification() uint64 {
	sesString, _ := session.GlobalSession.Values["userID"]
	userID := sesString.(uint64)
	return userID
}

// GetUserInformation stores user data
func GetUserInformation() models.Student {
	sesString, _ := session.GlobalSession.Values["user"]
	userName := sesString.(models.Student)
	return userName
}

// GetAllPostInformation processes all the information to be rendered on user's
// home page.
func GetAllPostInformation() ([]models.PublicHomePost, uint64, []models.Comment) {
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
	}

	for _, q = range commentLists {

		listComment = append(listComment, q)
		commentCounter++
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

	for i := commentCounter; i > -1; i-- {
		newCommentList = append(newCommentList, listComment[i])
		k++
	}

	return newList, totalLikesPerPost, newCommentList
}

// CreateNewPost processes all data to be posted to home page
func CreateNewPost(w http.ResponseWriter, r *http.Request) {
	author := GetUserIdentification()
	userName := GetUserInformation()
	topic := "Random"           //r.FormValue("homeFeedPost")
	category := "Entertainment" //r.FormValue("homeFeedPost")
	content := r.FormValue("homeFeedPost")
	datePosted := time.Now()
	var likes uint64
	var upvotes uint64
	var downvotes uint64
	var comments uint64

	createdTime := FormatDateTime(datePosted)

	userUserName := userName.UserName

	postInserter := models.InsertNewHomePost(author, userUserName, topic, category, content, createdTime, likes, upvotes, downvotes, comments)
	fmt.Println(postInserter)
}
