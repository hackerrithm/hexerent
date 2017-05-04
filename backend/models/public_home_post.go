package models

import (
	"fmt"
	"hexerent/backend/database"
)

// PublicHomePost exported
type PublicHomePost struct {
	PostID  int64
	Comment string `json:"homeFeedPost"`
}

// NewPublicHomePost : Acts as a constructor
func NewPublicHomePost(postid int64, comment string) *PublicHomePost {
	publicHomePost := new(PublicHomePost)
	publicHomePost.Comment = comment

	return publicHomePost
}

// InsertNewHomePost inserts a new user in the sql database
func InsertNewHomePost(comment string) PublicHomePost {
	DB, err := database.NewOpen()

	var insertStatement = "INSERT post SET Comment=?"
	stmt, err := DB.Prepare(insertStatement)
	//checkErr(err)

	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(comment)

	//checkErr(err)
	if err != nil {
		fmt.Println(err)
	}

	id, err := res.LastInsertId()
	//checkErr(err)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id)

	homeFeedPost := NewPublicHomePost(id, comment)

	return *homeFeedPost

}

// FindAllHomeFeedPosts stuff
func FindAllHomeFeedPosts() []string /*[]PublicHomePost.Comment*/ {

	//x := make(map[int][]string)

	// empty list of Posts
	//postLists := []PublicHomePost{}

	var postLists = []string{}

	homePost := NewPublicHomePost(0, "")

	DB, err := database.NewOpen()

	rows, err := DB.Query("SELECT * FROM post")
	if err != nil {
		fmt.Println(err)
	}

	//i := 0

	for rows.Next() {
		var postid int64
		var comment string

		err = rows.Scan(&postid, &comment)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(postid, " this is id")
		fmt.Println(comment, " this is comment")

		homePost.PostID = postid
		homePost.Comment = comment

		postLists = append(postLists, homePost.Comment /*PublicHomePost{homePost.PostID, homePost.Comment}*/)
		//x[i] = append(x[i], homePost.Comment)
		//postLists = append(postLists, x[i])

		//i++

	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("this is postLists final: ", postLists)

	DB.Close()

	return postLists
}
