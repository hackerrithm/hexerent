package models

import (
	"fmt"
	"hexerent/backend/database"
)

// Downvote userd to store information on a single like
type Downvote struct {
	DownvoteID uint64
	UserID     uint64
	PostID     uint64
	isClikcked bool
}

// NewDownvote keeps information on a new like
func NewDownvote(downvoteid, userid, postid uint64, isclikcked bool) *Downvote {
	downvote := new(Downvote)
	downvote.DownvoteID = downvoteid
	downvote.UserID = userid
	downvote.PostID = postid
	downvote.isClikcked = isclikcked

	return downvote
}

// FindAllDownvotePerPost finds all likes per post
func FindAllDownvotePerPost(postID uint64) uint64 {

	var result uint64
	DB, err := database.NewOpen()

	countedDownvotesResult, err := DB.Query("SELECT * FROM downvote WHERE PostID=?", postID)
	if err != nil {
		fmt.Println(err)
	}

	for countedDownvotesResult.Next() {
		result = result + 1
	}

	DB.Close()

	fmt.Println("Number of Downvotes for u:", result)

	return result
}

// CreateDownvote creates a Downvotes
func CreateDownvote(downvote *Downvote, postID, userID uint64) Downvote {

	DB, err := database.NewOpen()

	var insertStatement = "INSERT INTO downvote SET UserID=?,PostID=?"
	stmt, err := DB.Prepare(insertStatement)

	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(downvote.UserID, downvote.PostID)

	if err != nil {
		fmt.Println(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println(err)
	}

	// converst int64 to uint64 (unsigned)
	downvote.DownvoteID = uint64(id)

	// Updates after

	DB.Close()

	DB, err = database.NewOpen()

	var updateStatement = "UPDATE post SET Downvotes= Downvotes + 1 WHERE postID=?"
	stmt, err = DB.Prepare(updateStatement)

	if err != nil {
		fmt.Println(err)
	}

	_, err = stmt.Exec(downvote.PostID)

	if err != nil {
		fmt.Println(err)
	}

	DB.Close()

	return *downvote

}
