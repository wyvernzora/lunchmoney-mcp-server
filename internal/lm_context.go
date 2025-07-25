package internal

import (
	"context"
	"net/http"

	lm "github.com/icco/lunchmoney"
	"github.com/mark3labs/mcp-go/server"
)

type lmClientKeyType struct{}

var lmClientKey = lmClientKeyType{}

// WithLunchMoneyClient returns a StreamableHTTPOption that injects a
// LunchMoney API client (configured with the given token) into the HTTP request context.
func WithLunchMoneyClient(token string) server.StreamableHTTPOption {
	client, _ := lm.NewClient(token)
	return server.WithHTTPContextFunc(func(ctx context.Context, r *http.Request) context.Context {
		return context.WithValue(ctx, lmClientKey, client)
	})
}

// LunchMoneyClientFromContext retrieves the LunchMoney client stored in the context.
// It panics if no client has been added.
func LunchMoneyClientFromContext(ctx context.Context) *lm.Client {
	c, ok := ctx.Value(lmClientKey).(*lm.Client)
	if !ok || c == nil {
		panic("LunchMoney client not found in context")
	}
	return c
}
