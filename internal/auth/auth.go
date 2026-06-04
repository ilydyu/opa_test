package auth

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync/atomic"
	"time"

	"github.com/open-policy-agent/opa/v1/rego"
)

type Claims struct {
	Sub   string   `json:"sub"`
	Name  string   `json:"name"`
	Email string   `json:"email"`
	Roles []string `json:"roles"`
}

func NewMockClaims() *Claims {
	return &Claims{
		Sub:   "user-123",
		Name:  "Test User",
		Email: "test@example.com",
		Roles: []string{"reader"},
	}
}

type ContextKey string

const ClaimsKey ContextKey = "claims"
const interval = 5 * time.Second

var prepared atomic.Value

func InitOPA(policyPath string) error {
	query, err := initOPA(policyPath)

	if err != nil {
		return err
	}

	prepared.Store(query)

	// Наивная реализация релоада, лучшим решением было бы сделать подписку, и ждать событие по дескриптору.
	go func() {
		for {
			time.Sleep(interval)
			q, err := initOPA(policyPath)
			if err != nil {
				log.Printf("reload policy: %v", err)
				continue
			}
			prepared.Store(q)
			log.Println("policy reloaded")
		}
	}()

	return nil
}

func GetPreparedQuery() *rego.PreparedEvalQuery {
	return prepared.Load().(*rego.PreparedEvalQuery)
}

func initOPA(policyPath string) (*rego.PreparedEvalQuery, error) {
	policy, err := os.ReadFile(policyPath)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %w", err)
	}

	query, err := rego.New(
		rego.Query("data.app.auth.allow"),
		rego.Module("auth.rego", string(policy)),
	).PrepareForEval(context.Background())

	if err != nil {
		return nil, fmt.Errorf("rego.New: %w", err)
	}

	return &query, nil
}
