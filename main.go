package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"context"
	"github.com/gorilla/handlers"

	r "github.com/Maximo-Miranda/example-api-rest/internal/routes"
	migratations "github.com/Maximo-Miranda/example-api-rest/migrations/functions"
)

// main ...
func main(){

	var wait time.Duration

	port := os.Getenv("EXAMPLE_API_REST_APP_PORT")

	if port == "" {
		log.Println("$PORT must be set")
		port = "8090"
	}

	if err := migratations.UpMigrations(); err != nil {
		fmt.Println(err)
	}

	routes := r.Router()

	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Content-Type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "PATCH", "POST", "PUT", "DELETE", "OPTIONS"})

	fmt.Println("Example APIRest Service Running on port: " + port)

	srv := &http.Server{
		Addr:         "0.0.0.0:"+port,
		// Timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(routes), // gorilla/mux instance.

	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)

	log.Println("Example APIRest Service shutting down")
	os.Exit(0)

}
