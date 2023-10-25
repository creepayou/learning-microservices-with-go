package main

import (
	"context"
	"example/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)


func main(){

	// pgdb.Dbmain()

	//creating new handlers
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hotateHandler := handlers.NewHotate(l)
	productHandler := handlers.NewProducts(l)
	personHandler := handlers.NewPersonHandler(l)

	//creating new serveMux and register the handlers
	serveMux := http.NewServeMux()
	serveMux.Handle("/", hotateHandler)
	serveMux.Handle("/products", productHandler)

	serveMux2 := mux.NewRouter()
	getRouter := serveMux2.Methods("GET").Subrouter()
	getRouter.HandleFunc("/getProducts", productHandler.GetProducts) 
	getRouter.HandleFunc("/getPerson", personHandler.GetPerson)

	serveMux2.Handle("/products", productHandler).Methods("GET")

	//create a new server
	server := &http.Server{
		Addr: ":8082",
		Handler: getRouter,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	//start server with error handler
	go func(){
		err := server.ListenAndServe()
		if err != nil{
			l.Fatal(err)
		}
	}()

	//trap signal and notify for graceful shutdown
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)
	signal.Notify(signalChannel, os.Kill)
	sig := <-signalChannel
	l.Println("Received Terminate, gracefully shutting down.", sig)

	timeoutContext, _ := context.WithTimeout(context.Background(), 10*time.Second)
	server.Shutdown(timeoutContext)
}