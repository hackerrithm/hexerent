package server

import (
	"fmt"
	//"github.com/codegangsta/negroni"
	//jwt "github.com/dgrijalva/jwt-go"

	"hexerent/backend/database"
	"hexerent/backend/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	//"hexerent/backend/config"
)

// DBInitializerCaller starts MySQL db
func DBInitializerCaller() {
	DB, err := database.NewOpen()

	if err != nil {
		panic(err.Error())
	}

	defer DB.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = DB.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("Connected to MySQL server")

}

// Router is the router
func Router() {

	// Caller for DBInitializerCaller
	DBInitializerCaller()

	//clientMachineInformer.MemoryData()

	router := routes.NewRouter()

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:9000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  120 * time.Minute,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Shutting down server")
		os.Exit(1)
	}()

	log.Fatal(srv.ListenAndServe())

}
