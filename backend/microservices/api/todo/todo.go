package todo

import (
	"encoding/json"
	"fmt"
	"hexerent/backend/models"
	"io"
	"io/ioutil"
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
		//return
	}

	//fmt.Fprintf(res, string(payload))
	return string(payload)
}

// Index renders the home.html file
func Index(w http.ResponseWriter, r *http.Request) []string { //[]models.Todo {

	var todoLists = models.RepoFindAllTodos()

	var v string // models.Todo

	var listTodos []string //[]models.Todo

	for _, v = range todoLists {

		listTodos = append(listTodos, v)

	}

	//listOfTodos := jsonResponse(w, listTodos)
	return listTodos //listOfTodos

}

// Create stuff
func Create(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("name")
	completed := r.FormValue("completed")
	dueDate := r.FormValue("due")

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

	todo := models.NewTodo(0, name, completedTask, enteredTime)

	/*body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}*/

	models.RepoCreateTodo(todo)

}

// Show does stuff
func Show(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet || r.Method == "GET" {

		var newID int64

		vars := mux.Vars(r)
		todoID := vars["todoId"]
		fmt.Println("Todo show:", todoID)

		newID, _ = strconv.ParseInt(todoID, 10, 0)

		todoData := models.RepoFindTodo(newID)

		jsonResponse(w, todoData) /*

			b, err := json.Marshal(todoData.Name)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(b))

			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(string(b)))
		*/

	}

}

// Update does stuff
func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet || r.Method == "GET" {
		var newID int64

		vars := mux.Vars(r)
		todoID := vars["todoId"]
		fmt.Println("Todo show:", todoID)

		newID, _ = strconv.ParseInt(todoID, 10, 0)

		name := "Updated Kemar"  //r.FormValue("name")
		completed := ""          //r.FormValue("completed")
		dueDate := "2020-Jan-02" //r.FormValue("due")

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

		todo := models.NewTodo(newID, name, completedTask, enteredTime)

		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}
		if err := r.Body.Close(); err != nil {
			panic(err)
		}
		if err := json.Unmarshal(body, &todo); err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(422) // unprocessable entity
			if err := json.NewEncoder(w).Encode(err); err != nil {
				panic(err)
			}
		}

		t := models.RepoUpdateTodo(newID, todo)

		jsonResponse(w, t)
	}

}

// Delete does stuff
func Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet || r.Method == "GET" {

		var newID int64

		vars := mux.Vars(r)
		todoID := vars["todoId"]
		fmt.Println("Todo show:", todoID)

		newID, _ = strconv.ParseInt(todoID, 10, 0)

		models.RepoDestroyTodo(newID)

		jsonResponse(w, newID)

	}
}
