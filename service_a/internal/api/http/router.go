package api

import (
	deposithttp "github.com/doyoque/service_a/internal/api/http/deposit"
	"github.com/go-chi/chi/v5"
	chimware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/fx"
)

var Module = fx.Module("router",
	fx.Provide(initRouter),
	deposithttp.Module,
)

func initRouter() chi.Router {
	router := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	router.Use(
		cors.Handler,
		chimware.RequestID,
		chimware.Recoverer,
		chimware.RealIP,
	)

	return router
}
