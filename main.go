package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/david-galdamez/smile-routine-backend/database"
	mealtimes "github.com/david-galdamez/smile-routine-backend/meal_times"
	"github.com/david-galdamez/smile-routine-backend/user_settings"
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

	mealTimeRepo := mealtimes.NewMealTimeRepository(db)
	mealTimeService := mealtimes.NewMealTimeService(mealTimeRepo)
	userSettingRepo := user_settings.NewUserSettingsRepository(db)
	userSettingService := user_settings.NewUserSettingService(userSettingRepo)
	userRepo := users.NewUserRepository(db)
	userService := users.NewUserService(userRepo, mealTimeRepo, userSettingRepo)

	router.Handle("/api/users/", users.NewUserHandler(userService))
	router.Handle("/api/meal-time/", mealtimes.NewMealTimeHandler(mealTimeService))
	router.Handle("/api/user-settings/", user_settings.NewUserSettingsHandler(userSettingService))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(server.ListenAndServe())
}
