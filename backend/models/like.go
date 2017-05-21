package models

import (
	"fmt"
	"hexerent/backend/database"
)

// Like userd to store information on a single like
type Like struct {
	LikeID     uint64
	UserID     uint64
	PostID     uint64
	isClikcked bool
}

// NewLike keeps information on a new like
func NewLike(likeid, userid, postid uint64, isclikcked bool) *Like {
	like := new(Like)
	like.LikeID = likeid
	like.UserID = userid
	like.PostID = postid
	like.isClikcked = isclikcked

	return like
}

// FindAllLikesPerPost finds all likes per post
func FindAllLikesPerPost(postID uint64) uint64 {

	var result uint64
	DB, err := database.NewOpen()

	countedLikesResult, err := DB.Query("SELECT * FROM like WHERE PostID=?", postID)
	if err != nil {
		fmt.Println(err)
	}

	for countedLikesResult.Next() {
		result = result + 1
	}

	DB.Close()

	fmt.Println("Number of likes for u:", result)

	return result
}

// CreateLike creates a like
func CreateLike(like *Like, postID, userID uint64) Like {

	DB, err := database.NewOpen()

	var insertStatement = "INSERT INTO likes SET UserID=?,PostID=?"
	stmt, err := DB.Prepare(insertStatement)

	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(like.UserID, like.PostID)

	if err != nil {
		fmt.Println(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println(err)
	}

	// converst int64 to uint64 (unsigned)
	like.LikeID = uint64(id)

	// Updates after

	DB.Close()

	DB, err = database.NewOpen()

	var updateStatement = "UPDATE post SET Likes= Likes + 1 WHERE postID=?"
	stmt, err = DB.Prepare(updateStatement)

	if err != nil {
		fmt.Println(err)
	}

	_, err = stmt.Exec(like.PostID)

	if err != nil {
		fmt.Println(err)
	}

	DB.Close()

	return *like

}


