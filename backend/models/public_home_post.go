package models

import (
	"fmt"
	"hexerent/backend/database"
	"time"
)

// PublicHomePost exported
type PublicHomePost struct {
	PostID     uint64
	Author     uint64
	Topic      string `json:"topic"`
	Category   string `json:"category"`
	Content    string `json:"content"`
	DatePosted string `json:"datePosted"`
	Likes      uint64
	Upvotes    uint64
	Downvotes  uint64
	Comments   uint64
}

// NewPublicHomePost : Acts as a constructor
func NewPublicHomePost(postid, author uint64, topic, category, content, datePosted string, likes, upvotes, downvotes, comments uint64) *PublicHomePost {
	publicHomePost := new(PublicHomePost)
	publicHomePost.PostID = postid
	publicHomePost.Author = author
	publicHomePost.Topic = topic
	publicHomePost.Category = category
	publicHomePost.Content = content
	publicHomePost.DatePosted = datePosted
	publicHomePost.Likes = likes
	publicHomePost.Upvotes = upvotes
	publicHomePost.Downvotes = downvotes
	publicHomePost.Comments = comments

	return publicHomePost
}

// InsertNewHomePost inserts a new post in the database
func InsertNewHomePost(author uint64, topic, category, content, datePosted string, likes, upvotes, downvotes, comments uint64) PublicHomePost {

	DB, err := database.NewOpen()

	var insertStatement = "INSERT post SET AuthorID=?,Topic=?,Category=?,Content=?,DatePosted=?,Likes=?,Upvotes=?,Downvotes=?, Comments=?"
	stmt, err := DB.Prepare(insertStatement)

	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(author, topic, category, content, datePosted, likes, upvotes, downvotes, comments)

	if err != nil {
		fmt.Println(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id)

	identoficationNumber := uint64(id)

	homeFeedPost := NewPublicHomePost(identoficationNumber, author, topic, category, content, datePosted, likes, upvotes, downvotes, comments)

	return *homeFeedPost

}

// FindHomeFeedPost finds a single Post information
func FindHomeFeedPost(id uint64) PublicHomePost {

	var authorid, likes, upvotes, downvotes, comments uint64
	var topic, category, content, datePosted string

	DB, err := database.NewOpen()

	rows, err := DB.Prepare("SELECT * FROM post WHERE PostID = ?")
	if err != nil {
		fmt.Println(err)
	}

	rows.QueryRow(id).Scan(&authorid, &topic, &category, &content, &datePosted, &likes, &upvotes, &downvotes, &comments)

	if err != nil {
		fmt.Println(err)
	}

	post := NewPublicHomePost(id, authorid, topic, category, content, datePosted, likes, upvotes, downvotes, comments)

	DB.Close()

	return *post
}

// FindAllHomeFeedPosts stuff
func FindAllHomeFeedPosts() []PublicHomePost {

	// empty list of Posts
	postLists := []PublicHomePost{}

	const shortForm = "2006-Jan-02"
	enteredTime, _ := time.Parse(shortForm, "2016-May-22")
	fmt.Println(enteredTime, " : this is the time entered")

	homePost := NewPublicHomePost(1, 1, "", "", "", "", 1, 1, 1, 1)

	DB, err := database.NewOpen()

	rows, err := DB.Query("SELECT * FROM post")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var postid, authorid, likes, upvotes, downvotes, comments uint64
		var topic, category, content, datePosted string

		err = rows.Scan(&postid, &authorid, &topic, &category, &content, &datePosted, &likes, &upvotes, &downvotes, &comments)
		if err != nil {
			fmt.Println(err)
		}

		homePost.PostID = postid
		homePost.Author = authorid
		homePost.Topic = topic
		homePost.Category = category
		homePost.Content = content
		homePost.DatePosted = datePosted
		homePost.Likes = likes
		homePost.Upvotes = upvotes
		homePost.Downvotes = downvotes
		homePost.Comments = comments

		postLists = append(postLists, PublicHomePost{
			homePost.PostID,
			homePost.Author,
			homePost.Topic,
			homePost.Category,
			homePost.Content,
			homePost.DatePosted,
			homePost.Likes,
			homePost.Upvotes,
			homePost.Downvotes,
			homePost.Comments,
		})

		/*fmt.Println(homePost.PostID, "--")
		fmt.Println(authorid, "--")
		fmt.Println(topic, "--")
		fmt.Println(category, "--")
		fmt.Println(content, "--")
		fmt.Println(datePosted, "--")
		fmt.Println(likes, "--")
		fmt.Println(upvotes, "--")
		fmt.Println(downvotes, "--")*/

	}

	DB.Close()

	return postLists
}
