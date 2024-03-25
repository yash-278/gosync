package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
)

// type apiConfig struct {
// 	DB *database.Queries
// }

func router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(cors.AllowAll().Handler)
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	return r
}

func StartServer() {
	port := os.Getenv("PORT")

	fmt.Printf("Server starting on port %s", port)

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router())
	if err != nil {
		panic(err)
	}
}
