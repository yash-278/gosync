package routes

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/yash-278/gosync/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func (apiCfg *apiConfig) router() *chi.Mux {
	r := chi.NewRouter()
	apiRouter := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(cors.AllowAll().Handler)

	r.Mount("/v1", apiRouter)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})
	apiRouter.Get("/readiness", apiCfg.readinessSuccessHandler)
	apiRouter.Get("/err", apiCfg.readinessErrHandler)

	return r
}

func StartServer() {
	port := os.Getenv("PORT")
	dbURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err.Error())
		log.Fatal("Error connecting to DB")
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	log.Printf("Serving on port: %s\n", port)

	err = http.ListenAndServe(fmt.Sprintf(":%s", port), apiCfg.router())

	log.Fatal(err)

	if err != nil {
		panic(err)
	}
}
