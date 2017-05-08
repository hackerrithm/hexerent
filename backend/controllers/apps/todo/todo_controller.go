package todo

import (
	"fmt"
	"hexerent/backend/config"
	"hexerent/backend/microservices/api/todo"
	"hexerent/backend/session"

	"net/http"
)

// SessionCart group some session
type SessionCart struct {
	Sess1 //sessionVar1 string
	Todo  //[]models.PublicHomePost
}

// Todo to be used as interface
type Todo interface{}

// Sess1 to be used as interface
type Sess1 interface{}

// IndexTodoPage renders the apps-page.html file
func IndexTodoPage(w http.ResponseWriter, r *http.Request) {

	listOfTodos := todo.Index(w, r)

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.

	//sesString, _ := login.TestSession.Values["user"]
	sesString, _ := session.GlobalSession.Values["user"]
	//sesString2, _ := session.GlobalSession.Values["firstTimeUser"]
	fmt.Println(sesString, " : THIS IS sesString")
	//fmt.Println(sesString2, " : THIS IS sesString2")
	fmt.Println(listOfTodos, " : THIS IS list of todos")

	passedSession := SessionCart{
		sesString,
		//sesString2,
		listOfTodos,
	}

	config.Tpl.ExecuteTemplate(w, "todo-app.html", passedSession)
}

// CreateTodoPage renders the apps-page.html file
func CreateTodoPage(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet || r.Method == "GET" {

		sesString, _ := session.GlobalSession.Values["user"]

		config.Tpl.ExecuteTemplate(w, "todo-app-create.html", sesString)

	} else {
		todo.Create(w, r)

		// Get a session. We're ignoring the error resulted from decoding an
		// existing session: Get() always returns a session, even if empty.

		//sesString, _ := login.TestSession.Values["user"]
		sesString, _ := session.GlobalSession.Values["user"]
		//sesString2, _ := session.GlobalSession.Values["firstTimeUser"]
		fmt.Println(sesString, " : THIS IS sesString")

		config.Tpl.ExecuteTemplate(w, "todo-app-create.html", sesString)
	}

}
