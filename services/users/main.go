package main

import (
	// "database/sql"

	"github.com/gorilla/mux"

	// "scooter/users/internal/model"
	"scooter/users/internal/server"
	// "scooter/users/internal/middleware"
)

// função principal
func main() {
	// db, err := sql.Open("postgresql", "connection_string")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()

	// Create a database object using your database package
	// databaseObj := model.NewUserDB(db)

	// Create the router
	router := mux.NewRouter()

	// Register the routes
	server.RegisterRoutes(router, 3000)

	// Apply the dbMiddleware to all route handlers
	// router.Use(middleware.dbMiddleware(databaseObj))
}
