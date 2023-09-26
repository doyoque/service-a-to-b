package main

import (
	"fmt"
	"net/http"

	"github.com/doyoque/service_a/conf"
	api "github.com/doyoque/service_a/internal/api/http"
	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)

func main() {
	conf.InitEnv()
	fx.New(
		api.Module,
		fx.Invoke(func(apiRouter chi.Router) {
			mainRouter := chi.NewRouter()
			mainRouter.Mount("/", apiRouter)

			logrus.Infof("listening port: %d\n", conf.AppPort)
			http.ListenAndServe(fmt.Sprintf("%s:%d", "", conf.AppPort), mainRouter)
		}),
	)
}
