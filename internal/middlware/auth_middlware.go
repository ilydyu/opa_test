package middlware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ilydyu/opa_test/internal/auth"
	"github.com/open-policy-agent/opa/v1/rego"
)

func (m *Middlware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, ok := r.Context().Value(auth.ClaimsKey).(*auth.Claims)
		if !ok {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "missing token", http.StatusUnauthorized)
				return
			}

			// Для простоты использован мок, роль всегда - reader. В тестах есть кейс с ролью admin
			claims = auth.NewMockClaims()
			ctx := context.WithValue(r.Context(), auth.ClaimsKey, claims)
			r = r.WithContext(ctx)
		}

		input := map[string]any{
			"method": r.Method,
			"path":   r.URL.Path,
			"roles":  claims.Roles,
		}

		results, err := auth.GetPreparedQuery().Eval(r.Context(), rego.EvalInput(input))
		if err != nil {
			http.Error(w, "policy evaluation error", http.StatusInternalServerError)
			return
		}

		if len(results) == 0 || !results.Allowed() {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
