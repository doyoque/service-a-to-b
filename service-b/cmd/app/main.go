package main

import (
	"fmt"
	"net/http"

	api "github.com/doyoque/service-b/internal/api/http"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		api.Module,
		fx.Invoke(func(apiRouter chi.Router) {
			mainRouter := chi.NewRouter()
			mainRouter.Mount("/", apiRouter)

			logrus.Infof("listening port: %d\n", 8100)
			http.ListenAndServe(fmt.Sprintf("%s:%d", "", 8100), mainRouter)
		}),
	)
}
