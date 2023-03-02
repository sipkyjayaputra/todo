package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"todo/app"
	"todo/config"
)

func main() {
	// init config
	config.InitConfig()

	// init db connection
	gormDB, sqlDB, err := config.Connect()
	if err != nil {
		log.Println("Error on create connection", err.Error())
	}
	defer sqlDB.Close()

	//run server here
	router := app.InitRouter(gormDB)
	port := config.CONFIG["PORT"]
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// goroutine for listen and serve
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// channel to shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown server...")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %s\n", err)
	}
	log.Println("Server exiting")
}
