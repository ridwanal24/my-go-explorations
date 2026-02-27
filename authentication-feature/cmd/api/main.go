package main

import (
	"authentication-feature/internal/config"
	"authentication-feature/internal/handler"
	"authentication-feature/internal/repositories"
	"authentication-feature/internal/services"
	"log"
	"net/http"
)

func main() {
	// routes := routes.SetupRoutes()

	// db, err := config.InitDB()

	// defer db.Close()

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// fmt.Println("Serve to port 8000")
	// err := http.ListenAndServe(":8000", routes)

	// if err != nil {
	// 	log.Println(err.Error())
	// }

	db := config.InitDB()
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	authService := services.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/login", authHandler.Login)
	mux.HandleFunc("GET /user", userHandler.GetUser)

	log.Println("Server running on port 8000")
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}

}
