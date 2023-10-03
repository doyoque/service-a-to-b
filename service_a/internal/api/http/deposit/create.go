package deposithttp

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Create(
	router chi.Router,
) {
	router.Post("/deposits", func(w http.ResponseWriter, r *http.Request) {

	})
}
