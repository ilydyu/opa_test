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

func TestReaderCannotPost(t *testing.T) {
	err := auth.InitOPA("../../policy/auth.rego")

	if err != nil {
		t.Fatal(err)
	}

	m := middlware.NewMiddlware()
	controller := controller.NewController()
	handler := m.AuthMiddleware(http.HandlerFunc(controller.CreateResource))

	t.Run("Reader cannot create resource", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/resource", nil)

		rr := httptest.NewRecorder()

		claims := &auth.Claims{
			Roles: []string{"reader"},
		}
		ctx := context.WithValue(req.Context(), auth.ClaimsKey, claims)
		req = req.WithContext(ctx)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusForbidden {
			t.Errorf("expected 403, got %d", rr.Code)
		}
	})

	t.Run("Admin can create resource", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/resource", nil)

		rr := httptest.NewRecorder()

		claims := &auth.Claims{
			Roles: []string{"admin"},
		}
		ctx := context.WithValue(req.Context(), auth.ClaimsKey, claims)
		req = req.WithContext(ctx)
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected 204, got %d", rr.Code)
		}
	})
}
