package main

import (
	api "github.com/doyoque/service-a/internal/api/http"
	"github.com/go-chi/chi"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		api.Module,
		fx.Invoke(func(apiRouter chi.Router) {
			// mainRouter := chi.NewRouter()
			// mainRouter.Mount("/", apiRouter)

			// logrus.Infof("listening port: %d\n", 8000)
			// http.ListenAndServe(fmt.Sprintf("%s:%d", "", 8000), mainRouter)
		}),
	)
}
