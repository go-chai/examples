package router

import (
	"errors"
	"net/http"

	"github.com/go-chai/chai"
	"github.com/go-chai/examples/pkg/controller"
	"github.com/go-chai/examples/pkg/httputil"
	"github.com/go-chi/chi/v5"
)

func GetRoutes() *chi.Mux {
	r := chi.NewRouter()

	c := controller.NewController()

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/accounts", func(r chi.Router) {
			chai.GetG(r, "/{id}", c.ShowAccount)
			chai.GetG(r, "/", c.ListAccounts)
			chai.PostG(r, "/", c.AddAccount)
			chai.Delete(r, "/{id}", c.DeleteAccount)
			chai.Patch(r, "/{id}", c.UpdateAccount)
			chai.Post(r, "/{id}/images", c.UploadAccountImage)
		})

		r.Route("/bottles", func(r chi.Router) {
			chai.GetG(r, "/{id}", c.ShowBottle)
			chai.GetG(r, "/", c.ListBottles)
		})

		r.Route("/admin", func(r chi.Router) {
			r.Use(auth)

			chai.PostG(r, "/auth", c.Auth)
		})

		r.Route("/examples", func(r chi.Router) {
			chai.GetG(r, "/ping", c.PingExample)
			chai.GetG(r, "/calc", c.CalcExample)
			// chai.GetG(r, "/group{s/{gro}up_id}/accounts/{account_id}", c.PathParamsExample)
			chai.GetG(r, "/groups/{group_id}/accounts/{account_id}", c.PathParamsExample)
			chai.GetG(r, "/header", c.HeaderExample)
			chai.GetG(r, "/securities", c.SecuritiesExample)
			chai.GetG(r, "/attribute", c.AttributeExample)
			chai.PostG(r, "/attribute", c.PostExample)
		})
	})

	return r
}
func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("Authorization")) == 0 {
			httputil.NewError(w, http.StatusUnauthorized, errors.New("Authorization is required Header"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
