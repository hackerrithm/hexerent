package models

import (
	"fmt"
	"hexerent/backend/database"
	"time"
)

// Todo struct
type Todo struct {
	TodoID      uint64
	Title       string    `json:"title"`
	Name        string    `json:"name"`
	Category    string    `json:"category"`
	Status      string    `json:"status"`
	Completed   bool      `json:"completed"`
	Due         time.Time `json:"due"`
	DateCreated time.Time `json:"dateCreated"`
	UserID      uint64
}

// NewTodo : Acts as a constructor
func NewTodo(todoid uint64, title, name, category, status string, completed bool, due, dateCreated time.Time, userID uint64) *Todo {
	todo := new(Todo)
	todo.TodoID = todoid
	todo.Title = title
	todo.Name = name
	todo.Category = category
	todo.Status = status
	todo.Completed = completed
	todo.Due = due
	todo.DateCreated = dateCreated
	todo.UserID = userID
	return todo
}

// RepoFindTodo do
func RepoFindTodo(id, userID uint64) Todo {

	const shortForm = "2006-Jan-02"
	enteredTime, _ := time.Parse(shortForm, "2011-01-19")
	var title, name, category, status string
	var completed bool
	dateCreated, _ := time.Parse(shortForm, "2011-01-19")
	due, _ := time.Parse(shortForm, "2011-01-19")

	todo := NewTodo(id, "", "", "", "", true, enteredTime, enteredTime, userID)

	DB, err := database.NewOpen()

	rows, err := DB.Prepare("SELECT Title, Name, Category, Status, Completed, Due, DateCreated FROM todo WHERE ToDoID = ?")
	if err != nil {
		fmt.Println(err)
	}

	rows.QueryRow(id).Scan(&title, &name, &category, &status, &completed, &due, &dateCreated, &userID)

	if err != nil {
		fmt.Println(err)
	}

	//var todoItem Todo
	/*
		todo.TodoID = id
		todo.Title = title
		todo.Name = name
		todo.Category = category
		todo.Status = status
		todo.Completed = completed
		todo.Due = due
		todo.DateCreated = dateCreated
		todo.UserID = userID

		todoItem := Todo{
			todo.TodoID,
			todo.Title,
			todo.Name,
			todo.Category,
			todo.Status,
			todo.Completed,
			todo.Due,
			todo.DateCreated,
			todo.UserID,
		}
	*/

	todo = NewTodo(id, title, name, category, status, completed, due, dateCreated, userID)

	/*todoItem = Todo{
		id, title, name, category, status, completed, due, dateCreated, userID,
	}*/

	fmt.Println("CHECK OUT TODO NAME: ", name, " Please tell me if working")

	DB.Close()
	return *todo
}

// RepoFindAllTodos stuff
func RepoFindAllTodos(userID uint64) []Todo {

	todoLists := []Todo{}

	const shortForm = "2006-Jan-02"
	enteredTime, _ := time.Parse(shortForm, "2011-01-19")

	todo := NewTodo(0, "", "", "", "", true, enteredTime, enteredTime, userID)

	DB, err := database.NewOpen()

	rows, err := DB.Query("SELECT * FROM todo WHERE UserID=?", userID)
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var todoid, userID uint64
		var title, name, category, status, due, dateCreated string
		var completed bool
		var tinyInt string

		err = rows.Scan(&todoid, &title, &name, &category, &status, &tinyInt, &due, &dateCreated, &userID)
		if err != nil {
			fmt.Println(err)
		}

		completed = false
		if len(tinyInt) == 1 && tinyInt[0] == 1 {
			completed = true
		}

		enteredTime2, _ := time.Parse(shortForm, due)

		todo.TodoID = todoid
		todo.Title = title
		todo.Name = name
		todo.Category = category
		todo.Status = status
		todo.Completed = completed
		todo.Due = enteredTime2
		todo.DateCreated = enteredTime2
		todo.UserID = userID

		todoLists = append(todoLists, Todo{todo.TodoID,
			todo.Title,
			todo.Name,
			todo.Category,
			todo.Status,
			todo.Completed,
			todo.Due,
			todo.DateCreated,
			todo.UserID,
		})
	}

	DB.Close()

	return todoLists
}

// RepoFindAllTodosPerUser returns the number of todos for a single user
func RepoFindAllTodosPerUser(userID uint64) uint64 {
	var result uint64
	DB, err := database.NewOpen()

	countedTodosResult, err := DB.Query("SELECT * FROM todo WHERE UserID=?", userID)
	if err != nil {
		fmt.Println(err)
	}

	for countedTodosResult.Next() {
		result = result + 1
	}

	DB.Close()

	fmt.Println("Number of todos for u:", result)

	return result
}

// RepoCreateTodo stuff
func RepoCreateTodo(td *Todo, userID uint64) Todo {

	DB, err := database.NewOpen()

	var insertStatement = "INSERT todo SET Title=?,Name=?,Category=?,Status=?,Completed=?,Due=?,DateCreated=?,UserID=?"
	stmt, err := DB.Prepare(insertStatement)

	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(td.Title, td.Name, td.Category, td.Status, td.Completed, td.Due, td.DateCreated, td.UserID)

	if err != nil {
		fmt.Println(err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println(err)
	}

	// converst int64 to uint64 (unsigned)
	td.TodoID = uint64(id)

	DB.Close()

	return *td
}

// RepoUpdateTodo updates Todo list
func RepoUpdateTodo(id, userID uint64, td *Todo) Todo {

	DB, err := database.NewOpen()

	var updateStatement = "UPDATE todo SET Title=?,Name=?,Category=?,Status=?,Completed=?,Due=?,DateCreated=? WHERE ToDoID = ? AND userID=?"
	stmt, err := DB.Prepare(updateStatement)

	if err != nil {
		fmt.Println(err)
	}

	_, err = stmt.Exec(td.Title, td.Name, td.Category, td.Status, td.Completed, td.Due, td.DateCreated, id, userID)

	if err != nil {
		fmt.Println(err)
	}

	td.TodoID = id

	DB.Close()

	return *td
}

// RepoDestroyTodo do stuff
func RepoDestroyTodo(id uint64) {

	DB, err := database.NewOpen()

	rows, err := DB.Prepare("DELETE FROM todo WHERE ToDoID = ?")
	if err != nil {
		fmt.Println(err)
	}

	res, err := rows.Exec(id)

	if err != nil {
		fmt.Println(err)
	}

	affect, err := res.RowsAffected()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(affect)

	/*
		for i, t := range todos {
			if t.Id == id {
				todos = append(todos[:i], todos[i+1:]...)
				return nil
			}
		}
		return fmt.Errorf("Could not find Todo with id of %d to delete", id)
	*/

	DB.Close()
}
