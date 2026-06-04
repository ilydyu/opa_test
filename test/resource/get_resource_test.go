package resource

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ilydyu/opa_test/internal/auth"
	"github.com/ilydyu/opa_test/internal/controller"
	"github.com/ilydyu/opa_test/internal/middlware"
)

func TestReader(t *testing.T) {
	err := auth.InitOPA("../../policy/auth.rego")
	if err != nil {
		t.Fatal(err)
	}

	m := middlware.NewMiddlware()
	controller := controller.NewController()
	handler := m.AuthMiddleware(http.HandlerFunc(controller.GetResource))

	t.Run("Reader can GET", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/resource", nil)
		claims := &auth.Claims{
			Sub:   "user-123",
			Name:  "Test Reader",
			Roles: []string{"reader"},
		}
		ctx := context.WithValue(req.Context(), auth.ClaimsKey, claims)
		req = req.WithContext(ctx)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("expected 200, got %d", rr.Code)
		}
	})

	t.Run("No roles denied", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/resource", nil)
		claims := &auth.Claims{
			Roles: []string{},
		}
		ctx := context.WithValue(req.Context(), auth.ClaimsKey, claims)
		req = req.WithContext(ctx)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusForbidden {
			t.Errorf("expected 403, got %d", rr.Code)
		}
	})
}
