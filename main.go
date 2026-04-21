package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/david-galdamez/smile-routine-backend/appointment"
	"github.com/david-galdamez/smile-routine-backend/database"
	"github.com/david-galdamez/smile-routine-backend/habits"
	mealtimes "github.com/david-galdamez/smile-routine-backend/meal_times"
	"github.com/david-galdamez/smile-routine-backend/user_settings"
	"github.com/david-galdamez/smile-routine-backend/users"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load(".env")

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

	userSettingRepo := user_settings.NewUserSettingsRepository(db)
	mealTimeRepo := mealtimes.NewMealTimeRepository(db)
	mealTimeService := mealtimes.NewMealTimeService(mealTimeRepo, userSettingRepo)
	appointmentRepo := appointment.NewAppointmentRepository(db)
	appointmentService := appointment.NewAppointmentService(appointmentRepo)
	habitsRepo := habits.NewHabitsRepository(db)
	habitsService := habits.NewHabitsService(habitsRepo)
	userRepo := users.NewUserRepository(db)
	userService := users.NewUserService(userRepo, mealTimeRepo, userSettingRepo)

	router.Handle("/api/users/", users.NewUserHandler(userService))
	router.Handle("/api/meal-time/", mealtimes.NewMealTimeHandler(mealTimeService))
	router.Handle("/api/habits/", habits.NewHabitsHandler(habitsService))
	router.Handle("/api/appointment/", appointment.NewAppointmentHandler(appointmentService))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(server.ListenAndServe())
}
