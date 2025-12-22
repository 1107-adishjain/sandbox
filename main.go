package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/1107-adishjain/sandbox/config"
	"github.com/1107-adishjain/sandbox/routes"
	"github.com/1107-adishjain/sandbox/storage"
	"gorm.io/gorm"
	"os/signal"
	"syscall"
	"context"
	"time"
)

type Application struct{
	Cfg *config.Config
	DB *gorm.DB
}
func main(){

	cfg, err := config.LoadConfig()
	if err !=nil{
		fmt.Printf("Error loading config: %v\n",err)
	}

	db, err:= storage.NewConnection(cfg)
	if err != nil{
		fmt.Printf("Error connecting to database: %v\n",err)
	}
	app:= &Application{
		Cfg: cfg,
		DB: db,
	}
	srv:= &http.Server{
		Addr: ":"+cfg.Port,
		Handler: route.Routes(),
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