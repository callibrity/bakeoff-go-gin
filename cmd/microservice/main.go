package main

import (
	"fmt"
	"github.com/callibrity/bakeoff-go-gin/api/artists"
	"github.com/callibrity/bakeoff-go-gin/cmd/microservice/config"
	"github.com/callibrity/bakeoff-go-gin/internal/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

// ...

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err.Error())
	}

	pg, err := database.NewPostgres(cfg.Db.Host, cfg.Db.Name, cfg.Db.User, cfg.Db.Pass)
	if err != nil {
		log.Fatal(err.Error())
	}
	pg.DB.SetMaxIdleConns(5)
	pg.DB.SetMaxOpenConns(25)

	driver, err := postgres.WithInstance(pg.DB, &postgres.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Loading migrations from files...\n")
	migration, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("Executing migrations...\n")
	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err.Error())
	}

	// Instantiates the author service
	queries := database.New(pg.DB)
	artistsService := artists.NewService(queries)

	// Register our service handlers to the router
	router := gin.Default()
	fmt.Printf("Registering handlers...\n")
	artistsService.RegisterHandlers(router)

	// Start the server
	fmt.Printf("Starting Gin server...\n")
	router.Run()
}
