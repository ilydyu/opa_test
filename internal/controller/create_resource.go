package controller

import (
	"net/http"

	"github.com/ilydyu/opa_test/internal/auth"
)

func (c *Controller) CreateResource(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(auth.ClaimsKey).(*auth.Claims)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "Hello, ` + claims.Name + `", "resource": "secret data"}`))
}
