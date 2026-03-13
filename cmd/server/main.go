package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/erazorrr/go-link-shortener/internal/delivery/http/handlers"
	"github.com/erazorrr/go-link-shortener/internal/delivery/http/routes"
	"github.com/erazorrr/go-link-shortener/internal/repository"
	"github.com/erazorrr/go-link-shortener/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("unable to create dbpool: %v", err)
	}
	defer dbPool.Close()

	linksRepository := repository.NewLinkRepository(dbPool)
	linksCommandService := usecase.NewLinkCommandService(linksRepository)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	routes.RegisterLinkRoutes(router, handlers.NewLinkHandler(linksCommandService))

	addr := fmt.Sprintf("%s:%s", os.Getenv("ADDR"), os.Getenv("PORT"))
	log.Printf("Started listening on %s\n", addr)
	http.ListenAndServe(addr, router)
}
