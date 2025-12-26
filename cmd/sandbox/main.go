package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1107-adishjain/sandbox/app"
	"github.com/1107-adishjain/sandbox/config"
	"github.com/1107-adishjain/sandbox/models"
	"github.com/1107-adishjain/sandbox/routes"
	"github.com/1107-adishjain/sandbox/storage"
)

func main(){
	// load the configuration variables.
	cfg, err := config.LoadConfig()
	if err !=nil{
		fmt.Printf("Error loading config: %v\n",err)
	}
	// initialize database connection
	db, err:= storage.NewConnection(cfg)
	if err != nil{
		fmt.Printf("Error connecting to database: %v\n",err)
	}
	// #closing database connection when main go routine ends.
	defer func(){
		err:= storage.CloseConnection(db)
		if err!= nil{
			fmt.Printf("Error closing database connection: %v\n",err)
		}
	}()
	
	err= models.MigrateBooks(db)
	if err!=nil{
		fmt.Printf("Error migrating database: %v\n",err)
		return
	}

	app:= &app.Application{
		Cfg: cfg,
		DB: db,
	}

	srv:= &http.Server{
		Addr: ":"+cfg.Port,
		Handler: route.Routes(app),
		IdleTimeout: 120 * time.Second,
	}

	shutdown:= make(chan error) 
	go func(){
		// this will listen for shutdown signals and gracefully shutdown the server
		quit:=  make(chan os.Signal,1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
		s:=<-quit
		fmt.Printf("Received shutdown signal: %v\n",s)
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		shutdown <- srv.Shutdown(ctx)
	}()
	fmt.Printf("Starting server on port %v\n", cfg.Port)
	err = srv.ListenAndServe()
	if err != nil{
		fmt.Printf("Error starting server: %v\n",err)
	}	

	err =<-shutdown
	if err != nil{
		fmt.Printf("Error during server shutdown: %v\n",err)
	}
	fmt.Println("Server gracefully shhutted down")

}