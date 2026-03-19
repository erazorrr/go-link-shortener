package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/erazorrr/go-link-shortener/internal/delivery/http/handlers/cleanup"
	linkHandler "github.com/erazorrr/go-link-shortener/internal/delivery/http/handlers/link"
	"github.com/erazorrr/go-link-shortener/internal/delivery/http/routes"
	linkRepository "github.com/erazorrr/go-link-shortener/internal/repository/link"
	linkService "github.com/erazorrr/go-link-shortener/internal/usecase/link"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func main() {
	internalApiKey := os.Getenv("INTERNAL_API_KEY")
	if internalApiKey == "" {
		log.Fatal("INTERNAL_API_KEY env variable not specified")
	}

	dbPool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatalf("unable to create dbpool: %v", err)
	}
	defer dbPool.Close()

	opts, err := redis.ParseURL(os.Getenv("CACHE_URL"))
	if err != nil {
		log.Fatalf("unable to connect to cache: %v", err)
	}
	rdb := redis.NewClient(opts)

	maxConcurrentCacheWrites, err := strconv.Atoi(os.Getenv("MAX_CONCURRENT_CACHE_WRITES"))
	if err != nil || maxConcurrentCacheWrites < 1 {
		log.Printf("could not parse MAX_CONCURRENT_CACHE_WRITES, using 100 by default")
		maxConcurrentCacheWrites = 100
	}
	linksDBRepository := linkRepository.NewDBLinkRepository(dbPool)
	linksRepository := linkRepository.NewCachedDBLinkRepository(
		linksDBRepository,
		linkRepository.NewCacheLinkRepository(rdb),
		maxConcurrentCacheWrites,
	)
	linksQueryService := linkService.NewLinkQueryService(linksRepository)
	linksCommandService := linkService.NewLinkCommandService(linksRepository)
	linksCleanupService := linkService.NewLinkCleanupService(linksDBRepository)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	cleanupHandler := cleanup.NewCleanupHandler(linksCleanupService)
	routes.RegisterCleanupRoutes(router, internalApiKey, cleanupHandler)

	linkHandler := linkHandler.NewLinkHandler(linksQueryService, linksCommandService)
	routes.RegisterLinkRoutes(router, linkHandler)
	routes.RegisterRedirectRoute(router, linkHandler)

	addr := fmt.Sprintf("%s:%s", os.Getenv("ADDR"), os.Getenv("PORT"))
	log.Printf("Started listening on %s\n", addr)
	err = http.ListenAndServe(addr, router)
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("unable to listen and serve: %v", err)
	}
}
