package app

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type App struct {
	ctx        context.Context
	router     *chi.Mux
	logger     *zap.Logger
}


func (a *App) Init(ctx context.Context) error {
	a.ctx = ctx
	logger, err := zap.NewDevelopment()
	if err != nil {
		return err
	}
	defer func() { _ = logger.Sync() }()

	a.logger = logger
	// a.users = *users
	// a.links = *links
	// a.shortlinks = *shortlinks
	return nil
}

func (a *App) Serve() error {
	//Initialize Router and add Middleware	
	a.router = chi.NewRouter()
	a.router.Use(middleware.Logger)
	a.router.Use(middleware.Timeout(5 * time.Second))
	a.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	return http.ListenAndServe(":3333", a.router)
}
