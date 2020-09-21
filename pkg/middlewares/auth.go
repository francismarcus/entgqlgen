package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/francismarcus/eg/pkg/auth"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// User A stand-in for our database backed user object
type User struct {
	ID int
}

// AuthMiddleware decodes the share session cookie and packs the session into context
func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := r.Header.Get("Authorization")
			// Allow unauthenticated users in
			if t == "" {
				next.ServeHTTP(w, r)
				return
			}

			token, _ := auth.DecodeToken(t)
			userID, _ := strconv.Atoi(fmt.Sprintf("%.f", token["id"]))

			user := &User{
				ID: userID,
			}

			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// UserContext finds the user from the context. REQUIRES Middleware to have run.
func UserContext(ctx context.Context) *User {
	user := ctx.Value(userCtxKey).(*User)
	return user
}
