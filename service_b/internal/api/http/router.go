package api

import (
	"github.com/go-chi/chi"
	chimware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/fx"
)

var Module = fx.Module("router",
	fx.Provide(initRouter),
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
		chimware.RealIP,
	)

	return router
}
