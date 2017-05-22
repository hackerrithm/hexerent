package todo

import (
	"encoding/json"
	"fmt"
	"hexerent/backend/microservices/api/home"
	"hexerent/backend/models"
	"net/http"
	"strconv"

	"time"

	"github.com/gorilla/mux"
)

func errorCheck(r http.ResponseWriter, err error) bool {
	if err != nil {
		http.Error(r, err.Error(), http.StatusInternalServerError)
		return true
	}
	return false
}

func jsonResponse(res http.ResponseWriter, data interface{}) string {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")

	payload, err := json.Marshal(data)
	if errorCheck(res, err) {
		fmt.Println(err)
	}

	//fmt.Fprintf(res, string(payload))
	return string(payload)
}

// Index renders the home.html file
func Index(w http.ResponseWriter, r *http.Request) []models.Todo {

	userID := home.GetUserIdentification()

	var todoLists = models.RepoFindAllTodos(userID)
	var v models.Todo
	var listTodos []models.Todo

	for _, v = range todoLists {
		listTodos = append(listTodos, v)
	}
	return listTodos
}

// Create stuff
func Create(w http.ResponseWriter, r *http.Request) {

	title := r.FormValue("title")
	name := r.FormValue("name")
	category := "general" //r.FormValue("category")
	status := "active"    //r.FormValue("status")
	completed := ""       //r.FormValue("completed")
	dueDate := r.FormValue("due")
	dateCreated := time.Now().Local()
	userID := home.GetUserIdentification()

	fmt.Println("********* ************** this is userID: ", userID)

	var a = completed
	var completedTask bool
	isCompletedTicked := true
	if len(a) == 0 {
		fmt.Println(!isCompletedTicked)
		completedTask = !isCompletedTicked
	} else {
		completedTask = isCompletedTicked
	}

	const shortForm = "2006-01-02"
	enteredTime, _ := time.Parse(shortForm, dueDate)
	fmt.Println(enteredTime, " : this is the time entered")
	createdTime, _ := time.Parse(shortForm, dateCreated.Format("2006-01-02"))

	todo := models.NewTodo(0, title, name, category, status, completedTask, enteredTime, createdTime, userID)

	models.RepoCreateTodo(todo, userID)
}

// Show does stuff
func Show(w http.ResponseWriter, r *http.Request) models.Todo {

	var newID uint64

	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Println("Todo show:", todoID)

	newID, _ = strconv.ParseUint(todoID, 10, 0)
	userID := home.GetUserIdentification()

	todo := models.RepoFindTodo(newID, userID)

	//fmt.Println(todo.Title)

	return todo

}

// CountTodos returns the reults after counting the number of
// Todos per user
func CountTodos(w http.ResponseWriter, r *http.Request) uint64 {
	userID := home.GetUserIdentification()
	return models.RepoFindAllTodosPerUser(userID)
}

// Update does stuff
func Update(w http.ResponseWriter, r *http.Request) {
	var newID uint64

	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Println("Todo show:", todoID)

	newID, _ = strconv.ParseUint(todoID, 10, 0)

	title := r.FormValue("title")
	name := r.FormValue("name")
	category := "general" //r.FormValue("category")
	status := "active"    //r.FormValue("status")
	completed := r.FormValue("completed")
	dueDate := r.FormValue("due")
	dateCreated := time.Now().Local()
	userID := home.GetUserIdentification()

	var a = completed
	var completedTask bool
	isCompletedTicked := true
	if len(a) == 0 {
		fmt.Println(!isCompletedTicked)
		completedTask = !isCompletedTicked
	} else {
		completedTask = isCompletedTicked
	}

	const shortForm = "2006-Jan-02"
	enteredTime, _ := time.Parse(shortForm, dueDate)
	fmt.Println(enteredTime, " : this is the time entered")
	createdTime, _ := time.Parse(shortForm, dateCreated.Format("2006-01-02"))

	todo := models.NewTodo(newID, title, name, category, status, completedTask, enteredTime, createdTime, userID)

	models.RepoUpdateTodo(newID, userID, todo)

}

// Delete does stuff
func Delete(w http.ResponseWriter, r *http.Request) {

	var newID uint64

	vars := mux.Vars(r)
	todoID := vars["todoId"]
	fmt.Println("Todo show:", todoID)

	newID, _ = strconv.ParseUint(todoID, 10, 0)

	models.RepoDestroyTodo(newID)
}
