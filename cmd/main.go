package main

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"library/internal/config"
	"library/internal/handlers"
	"library/internal/hash/passwordHasher"
	database "library/internal/store/db"
	"library/internal/store/repo"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	m "library/internal/middleware"

	"github.com/go-chi/chi/v5/middleware"
)

// Environment
// Set to production at build time
// used to determine what assets to load
var Environment = "development"

func init() {
	err := os.Setenv("env", Environment)
	if err != nil {
		panic(err)
	}
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	r := chi.NewRouter()

	cfg := config.MustLoadConfig()

	database.MustOpen(cfg.DSN)
	pwHasher := passwordHasher.NewHPasswordHasher()

	userStore := repo.NewUserRepo(
		repo.NewUserRepoParams{
			PasswordHasher: pwHasher,
		},
	)

	sessionStore := repo.NewSessionRepo()

	fileServer := http.FileServer(http.Dir("./static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	authMiddleware := m.NewAuthMiddleware(sessionStore, cfg.SessionCookieName)

	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Logger,
			m.TextHTMLMiddleware,
			m.CSPMiddleware,
			authMiddleware.AddUserToContext,
		)

		r.NotFound(handlers.NewNotFoundHandler().ServeHTTP)

		r.Get("/", handlers.NewHomeHandler().ServeHTTP)

		r.Get("/about", handlers.NewAboutHandler().ServeHTTP)

		r.Get("/register", handlers.NewGetRegisterHandler().ServeHTTP)

		r.Post("/register", handlers.NewPostRegisterHandler(handlers.PostRegisterHandlerParams{
			UserStore: userStore,
		}).ServeHTTP)

		r.Get("/login", handlers.NewGetLoginHandler().ServeHTTP)

		r.Post("/login", handlers.NewPostLoginHandler(handlers.PostLoginHandlerParams{
			UserStore:         userStore,
			SessionStore:      sessionStore,
			PasswordHash:      pwHasher,
			SessionCookieName: cfg.SessionCookieName,
		}).ServeHTTP)

		r.Post("/logout", handlers.NewPostLogoutHandler(handlers.PostLogoutHandlerParams{
			SessionCookieName: cfg.SessionCookieName,
		}).ServeHTTP)
	})

	killSig := make(chan os.Signal, 1)

	signal.Notify(killSig, os.Interrupt, syscall.SIGTERM)

	srv := &http.Server{
		Addr:    cfg.Port,
		Handler: r,
	}

	go func() {
		err := srv.ListenAndServe()

		if errors.Is(err, http.ErrServerClosed) {
			logger.Info("Server shutdown complete")
		} else if err != nil {
			logger.Error("Server error", slog.Any("err", err))
			os.Exit(1)
		}
	}()

	logger.Info("Server started", slog.String("port", cfg.Port), slog.String("env", Environment))
	<-killSig

	logger.Info("Shutting down server")

	// Create a context with a timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error("Server shutdown failed", slog.Any("err", err))
		os.Exit(1)
	}

	logger.Info("Server shutdown complete")
}
