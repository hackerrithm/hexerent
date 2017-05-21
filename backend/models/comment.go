package models

import (
	"fmt"
	"hexerent/backend/database"
)

// Comment stores relevant datatypes for a comment
type Comment struct {
	CommentID   uint64
	UserID      uint64
	PostID      uint64
	CommentText string
}

// NewComment keeps information on a new comment
func NewComment(commentid, userid, postid uint64, commenttext string) *Comment {
	comment := new(Comment)
	comment.CommentID = commentid
	comment.UserID = userid
	comment.PostID = postid
	comment.CommentText = commenttext

	return comment
}

// InsertNewComment inserts a new comment for post  in the database
func InsertNewComment(comment *Comment, userid, postid uint64, commenttext string) Comment {

	DB, err := database.NewOpen()

	var insertStatement = "INSERT comment SET UserID=?,PostID=?,CommentText=?"
	stmt, err := DB.Prepare(insertStatement)

	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(userid, postid, commenttext)

	if err != nil {
		fmt.Println(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id)

	comment.CommentID = uint64(id)

	DB.Close()

	//Updates done here
	DB, err = database.NewOpen()

	var updateStatement = "UPDATE post SET Comments= Comments + 1 WHERE postID=?"
	stmt, err = DB.Prepare(updateStatement)

	if err != nil {
		fmt.Println(err)
	}

	_, err = stmt.Exec(comment.PostID)

	if err != nil {
		fmt.Println(err)
	}

	DB.Close()

	return *comment

}

// FindComment finds a single Comment information
func FindComment(id uint64) Comment {

	var userid, postid uint64
	var commenttext string

	DB, err := database.NewOpen()

	rows, err := DB.Prepare("SELECT * FROM comment WHERE CommentID = ?")
	if err != nil {
		fmt.Println(err)
	}

	rows.QueryRow(id).Scan(&userid, &postid, &commenttext)

	if err != nil {
		fmt.Println(err)
	}

	comment := NewComment(id, userid, postid, commenttext)

	DB.Close()

	return *comment
}

// FindAllComments stuff
func FindAllComments() []Comment {

	// empty list of Comments
	commentLists := []Comment{}

	commentPost := NewComment(1, 1, 1, "")

	DB, err := database.NewOpen()

	rows, err := DB.Query("SELECT * FROM comment")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var commentid, userid, postid uint64
		var commenttext string

		err = rows.Scan(&commentid, &userid, &postid, &commenttext)
		if err != nil {
			fmt.Println(err)
		}

		commentPost.CommentID = commentid
		commentPost.UserID = userid
		commentPost.PostID = postid
		commentPost.CommentText = commenttext

		commentLists = append(commentLists, Comment{
			commentPost.CommentID,
			commentPost.UserID,
			commentPost.PostID,
			commentPost.CommentText,
		})
	}

	DB.Close()

	return commentLists
}

// CountAllCommentsPerPost finds all likes per post
func CountAllCommentsPerPost(postID uint64) uint64 {

	var result uint64
	DB, err := database.NewOpen()

	countedCommentsResult, err := DB.Query("SELECT * FROM comment WHERE PostID=?", postID)
	if err != nil {
		fmt.Println(err)
	}

	for countedCommentsResult.Next() {
		result = result + 1
	}

	DB.Close()

	fmt.Println("Number of comments for u:", result)

	return result
}
