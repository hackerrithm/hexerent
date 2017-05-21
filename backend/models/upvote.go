package models

import (
	"fmt"
	"hexerent/backend/database"
)

// Upvote userd to store information on a single like
type Upvote struct {
	UpvoteID   uint64
	UserID     uint64
	PostID     uint64
	isClikcked bool
}

// NewUpvote keeps information on a new like
func NewUpvote(upvoteid, userid, postid uint64, isclikcked bool) *Upvote {
	upvote := new(Upvote)
	upvote.UpvoteID = upvoteid
	upvote.UserID = userid
	upvote.PostID = postid
	upvote.isClikcked = isclikcked

	return upvote
}

// FindAllUpvotePerPost finds all likes per post
func FindAllUpvotePerPost(postID uint64) uint64 {

	var result uint64
	DB, err := database.NewOpen()

	countedUpvotesResult, err := DB.Query("SELECT * FROM upvote WHERE PostID=?", postID)
	if err != nil {
		fmt.Println(err)
	}

	for countedUpvotesResult.Next() {
		result = result + 1
	}

	DB.Close()

	fmt.Println("Number of upvotes for u:", result)

	return result
}

// CreateUpvote creates a upvotes
func CreateUpvote(upvote *Upvote, postID, userID uint64) Upvote {

	DB, err := database.NewOpen()

	var insertStatement = "INSERT INTO upvote SET UserID=?,PostID=?"
	stmt, err := DB.Prepare(insertStatement)

	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(upvote.UserID, upvote.PostID)

	if err != nil {
		fmt.Println(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println(err)
	}

	// converst int64 to uint64 (unsigned)
	upvote.UpvoteID = uint64(id)

	// Updates after

	DB.Close()

	DB, err = database.NewOpen()

	var updateStatement = "UPDATE post SET Upvotes= Upvotes + 1 WHERE postID=?"
	stmt, err = DB.Prepare(updateStatement)

	if err != nil {
		fmt.Println(err)
	}

	_, err = stmt.Exec(upvote.PostID)

	if err != nil {
		fmt.Println(err)
	}

	DB.Close()

	return *upvote

}
