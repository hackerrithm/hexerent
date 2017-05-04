package routes

import (
	"flag"
	//"github.com/codegangsta/negroni"
	//jwt "github.com/dgrijalva/jwt-go"

	"hexerent/backend/controllers/apps"
	"hexerent/backend/controllers/chat"
	"hexerent/backend/controllers/home"
	"hexerent/backend/controllers/index"
	"hexerent/backend/controllers/login"
	"hexerent/backend/controllers/profile"
	"hexerent/backend/controllers/register"
	"hexerent/backend/microservices"
	"hexerent/backend/microservices/api/todo"
	"net/http"

	"github.com/gorilla/mux"
	//"hexerent/backend/config"
)

//Route struct for supplying multiple data to routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes with emty Route struct
type Routes []Route

// DataHandler is an empty interface
type DataHandler interface {
}

// NewRouter returns mux router
func NewRouter() *mux.Router {

	var dir string

	flag.StringVar(&dir, "dir", ".", "/")
	flag.Parse()

	//n := negroni.New(ValidateTokenMiddleware)

	hub := chat.NewHub()
	go hub.Run()

	var routes = Routes{
		Route{"Index", "GET", "/", index.Index},
		Route{"Index", "POST", "/", index.Index},
		Route{"About", "GET", "/about", index.About},
		Route{"Register", "GET", "/register", register.Register},
		Route{"Register", "POST", "/register", register.Register},
		Route{"Login", "GET", "/login", login.Login},
		Route{"Login", "POST", "/login", login.Login},
		Route{"Home", "GET", "/user/home", home.HomePage},
		Route{"Home", "POST", "/user/home", home.HomePage},
		Route{"Profile", "GET", "/user/profile", profile.ProfilePage},
		Route{"Logout", "GET", "/user/logout", login.Logout},
		Route{"AppsPage", "GET", "/user/profile/apps/all", apps.AppsPage},
		Route{"ToDoApp", "GET", "/user/profile/apps/todo", todo.Index},
		Route{"ToDoApp", "GET", "/user/profile/apps/todo/{todoId}", todo.Show},
		Route{"ToDoApp", "GET", "/user/profile/apps/todo/create", todo.TodoAppCreateClicked},
		Route{"ToDoApp", "POST", "/user/profile/apps/todo/create", todo.Create},
		Route{"ToDoApp", "GET", "/user/profile/apps/todo/update/{todoId}", todo.Update},
		Route{"ToDoApp", "GET", "/user/profile/apps/todo/delete/{todoId}", todo.Delete},
		Route{"ChatApp", "GET", "/user/chat", chat.ServeHome},
		Route{"ChatApp", "GET", "/user/chat/ws", func(w http.ResponseWriter, r *http.Request) {
			chat.ServeWs(hub, w, r)
		}},
	}

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = microservices.Logger(handler, route.Name)
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(handler /*route.HandlerFunc*/)

	}

	// for static files
	router.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	return router

}

/*
func ThaMethod() {
	kemar := models.NewUser("Kemar", "Kemar", "Galloway", "password", "emailaddress@mail.com", "251-6646", "876", "+1", "Jamaica", "Kingston", "kng13")
	kemar.SayHi()

	mike := models.NewStudent(models.NewUser("Mike", "Mike", "Campbell", "password", "emailaddress@mail.com", "551-6646", "876", "+1", "Jamaica", "Kingston", "kng13"), "MIT", 0.00)

	paul := models.NewStudent(models.NewUser("Jerry", "Jerry", "Springer", "password", "emailaddress@mail.com", "841-6446", "576", "+1", "USA", "New York", "NY34"), "Harvard", 100)

	sam := models.NewEmployee(models.NewUser("Tom", "Tom", "Timothy", "password", "emailaddress@mail.com", "243-6646", "976", "+1", "Japan", "Tokyo", "tky0"), "Golang Inc.", 1000)

	tom := models.NewEmployee(models.NewUser("Selena", "Selena", "Gomez", "password", "emailaddress@mail.com", "971-6646", "176", "+1", "China", "Shangai", "sh3"), "Things Ltd.", 5000)

	// define interface i
	var i models.Men

	//i can store Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")

	//i can store Employee
	i = tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	// slice of Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]models.Men, 3)
	// these three elements are different types but they all implemented interface Men
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x {
		value.SayHi()
	}
}

*/
/*
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	//validate token
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return config.PrivateKey, nil
	})

	if err == nil {

		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorised access to this resource")
	}

}
*/
