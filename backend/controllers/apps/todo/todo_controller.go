package todo

import (
	"hexerent/backend/config"
	"hexerent/backend/microservices/api/todo"
	"hexerent/backend/session"

	"net/http"
)

// SessionCart group some session
type SessionCart struct {
	Sess1
	Todo
	TodoCounter uint64
}

// Todo to be used as interface
type Todo interface{}

// Sess1 to be used as interface
type Sess1 interface{}

// IndexTodoPage renders the apps-page.html file
func IndexTodoPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet || r.Method == "GET" {
		listOfTodos := todo.Index(w, r)
		numberOfTodos := todo.CountTodos(w, r)

		sesString, _ := session.GlobalSession.Values["user"]

		passedSession := SessionCart{
			sesString,
			listOfTodos,
			numberOfTodos,
		}
		config.Tpl.ExecuteTemplate(w, "todo-app.html", passedSession)
	} else if r.Method == http.MethodPost {

		todo.Create(w, r)

		listOfTodos := todo.Index(w, r)

		sesString, _ := session.GlobalSession.Values["user"]

		passedSession := SessionCart{
			sesString,
			listOfTodos,
			1,
		}
		config.Tpl.ExecuteTemplate(w, "todo-app.html", passedSession)
	}
}

// UpdateTodoPage updates Todo list
func UpdateTodoPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet || r.Method == "GET" {

		sesString, _ := session.GlobalSession.Values["user"]

		todo := todo.Show(w, r)

		passedSession := SessionCart{
			sesString,
			todo,
			1,
		}

		config.Tpl.ExecuteTemplate(w, "todo-app-update.html", passedSession)

	} else if r.Method == http.MethodPost || r.Method == "POST" {
		todo.Update(w, r)
		http.Redirect(w, r, "/user/profile/apps/todo", http.StatusSeeOther)
	}

}

// DeleteTodoPage deletes a Todo
func DeleteTodoPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet || r.Method == "GET" {
		todo.Delete(w, r)
		http.Redirect(w, r, "/user/profile/apps/todo", http.StatusSeeOther)
	}

}
