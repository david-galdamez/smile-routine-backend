package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/david-galdamez/smile-routine-backend/database"
	"github.com/david-galdamez/smile-routine-backend/users"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error getting .env file: %v", err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Port env variable not found")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL env variable not found")
	}

	conn, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Cant connect to database")
	}

	db := database.New(conn)

	router := http.NewServeMux()

	router.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// user endpoints
	userRepo := users.NewUserRepository(db)
	userService := users.NewUserService(userRepo)

	router.Handle("/api/users/", users.NewUserHandler(userService))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(server.ListenAndServe())
}
