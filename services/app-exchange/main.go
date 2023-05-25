package main

import (
	"github.com/AppsLab-KE/backend-everyshilling/services/app-exchange/internal/routes/handlers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {

	router := gin.Default()

	// instance of the server handlers
	serverHandlers := handlers.NewServerHandlers()

	//accountRepo := storage.NewAccountRepository()
	//accountService := service.NewAccountService(accountRepo)

	// Register the handlers with the router
	handlers.RegisterHandlers(router, serverHandlers)

	port := "8080"

	// Starting the server
	log.Printf("Server listening on port %s", port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
