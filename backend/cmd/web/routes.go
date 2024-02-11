package main

import (
	"context"
	"fmt"
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

		userinfo, err := verifyToken(token)

		if err != nil {
			http.Error(w, "Invalid token 2", http.StatusUnauthorized)
			return
		}

		contextReq := context.WithValue(r.Context(), "token info", userinfo)

		// Token is valid, proceed to the next handler
		fmt.Println(userinfo)
		next.ServeHTTP(w, r.WithContext(contextReq))
	})
}
func (app *application) routes() *chi.Mux {
	mux := chi.NewRouter()
	// TODO: how to actually setup CORS for frontend????
	corsMW := cors.AllowAll()
	mux.Use(corsMW.Handler)
	mux.Post("/auth/login", app.loginUser)
	mux.Post("/auth/signup", app.signupUser)

	mux.Route("/auth", func(mux chi.Router) {
		mux.Use(TokenAuthMiddleware)
		mux.Get("/", app.getMe)
	})

	mux.Get("/auctions/active", app.getActiveAuctions)
	mux.Get("/auctions/{id}/bets", app.getAuctionBets)
	mux.Get("/auctions/{id}", app.getAuctionById)

	mux.Route("/auctions", func(r chi.Router) {
		r.Use(TokenAuthMiddleware)
		r.Post("/create", app.createAuction)
		r.Post("/{id}/makebet", app.makebet)
		r.Patch("/{id}/update", app.updateAuction)
	})

	mux.Route("/users", func(r chi.Router) {
		r.Get("/{id}/auction/active", app.getActiveAuctionsByUser)
	})

	return mux
}
