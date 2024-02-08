package main

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

// TODO: GPT generated, take with grain of salt, verify when refactoring
func TokenAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the Authorization header is present
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		// Extract the token from the Authorization header
		authToken := strings.Split(authHeader, " ")
		if len(authToken) != 2 || authToken[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}
		token := authToken[1]

        if err := verifyToken(token); err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
func (app *application) routes() *chi.Mux {
	mux := chi.NewRouter()

	// TODO: how to actually setup CORS for frontend????
    corsMW := cors.AllowAll()
	mux.Use(corsMW.Handler)
   //mux.Use(TokenAuthMiddleware)
	mux.Post("/user/login", app.loginUser)
	mux.Post("/user/signup", app.signupUser)

    mux.Post("/auction/create", app.createAuction)
	return mux
}
