package models

import (
	"fmt"
	"hexerent/backend/database"
	"time"
)

// Todo struct
type Todo struct {
	TodoID    int64
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// NewTodo : Acts as a constructor
func NewTodo(todoid int64, name string, completed bool, due time.Time) *Todo {
	todo := new(Todo)
	todo.TodoID = todoid
	todo.Name = name
	todo.Completed = completed
	todo.Due = due
	return todo
}

// RepoFindTodo do
func RepoFindTodo(id int64) Todo {

	const shortForm = "2006-Jan-02"
	enteredTime, _ := time.Parse(shortForm, "2011-01-19")
	var name string

	DB, err := database.NewOpen()

	rows, err := DB.Prepare("SELECT Name FROM todo WHERE ToDoID = ?")
	if err != nil {
		fmt.Println(err)
	}

	rows.QueryRow(id).Scan(&name)

	if err != nil {
		fmt.Println(err)
	}

	todo := NewTodo(id, name, true, enteredTime)

	fmt.Println(todo.Name)
	/*
		for _, t := range todos {
			if t.Id == id {
				return t
			}
		}
		// return empty Todo if not found
		return Todo{}
	*/

	DB.Close()
	return *todo
}

// RepoFindAllTodos stuff
func RepoFindAllTodos() []Todo {

	// empty list of Todos
	todoLists := []Todo{}

	const shortForm = "2006-Jan-02"
	enteredTime, _ := time.Parse(shortForm, "2011-01-19")

	todo := NewTodo(0, "", true, enteredTime)

	DB, err := database.NewOpen()

	rows, err := DB.Query("SELECT * FROM todo")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var todoid int64
		var name string
		var due string //time.Time
		var completed bool

		var tinyInt string

		//const shortForm = "2006-Jan-02"
		//enteredTime, _ := time.Parse(shortForm, "2011-01-19")
		//fmt.Println(enteredTime, " : this is the time entered")

		//xt := reflect.TypeOf(enteredTime).Kind()

		//fmt.Println("time is of type name: ", xt)

		err = rows.Scan(&todoid, &name, &tinyInt, &due)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(todoid, " this is id")
		fmt.Println(name, " this is name")
		completed = false
		if len(tinyInt) == 1 && tinyInt[0] == 1 {
			completed = true
		}

		fmt.Println(completed, " this is completed")
		fmt.Println(due, " this is due")

		enteredTime2, _ := time.Parse(shortForm, due)

		todo.TodoID = todoid
		todo.Name = name
		todo.Completed = completed
		todo.Due = enteredTime2

		todoLists = append(todoLists, Todo{todo.TodoID,
			todo.Name,
			todo.Completed,
			todo.Due})

		//fmt.Printf("/n/nthis is todoLists: ", todoLists)

	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("this is todoLists final: ", todoLists)

	DB.Close()

	return todoLists
}

// RepoCreateTodo stuff
func RepoCreateTodo(td *Todo) Todo {

	DB, err := database.NewOpen()

	var insertStatement = "INSERT todo SET Name=?,Completed=?,Due=?"
	stmt, err := DB.Prepare(insertStatement)
	//checkErr(err)

	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(td.Name, td.Completed, td.Due)

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
	fmt.Println(td.Name, " this is todo.name")
	fmt.Println(td, " this is todo")
	td.TodoID = id

	DB.Close()

	return *td
}

func RepoUpdateTodo(id int64, td *Todo) Todo {

	DB, err := database.NewOpen()

	var updateStatement = "UPDATE todo SET Name=?,Completed=?,Due=? WHERE ToDoID = ?"
	stmt, err := DB.Prepare(updateStatement)
	//checkErr(err)

	if err != nil {
		fmt.Println(err)
	}

	_, err = stmt.Exec(td.Name, td.Completed, td.Due, id)

	//checkErr(err)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(id)
	fmt.Println(td.Name, " this is todo.name")
	fmt.Println(td, " this is todo")
	td.TodoID = id

	DB.Close()

	return *td
}

// RepoDestroyTodo do stuff
func RepoDestroyTodo(id int64) {

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
