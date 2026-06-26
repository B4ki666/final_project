package main

import (
	"log"
	"os"

	service "final_project/internal/Service"
	"final_project/internal/config"
	"final_project/internal/db"
	"final_project/internal/repository"
	"final_project/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	cfg, err := config.Load()
	if err != nil {
		logger.Fatal(err)
	}

	db, err := db.InitDB(cfg.DBFile)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("db is ready")
	defer db.Close()

	repo := repository.NewRepository(db)
	svc := service.NewService(repo, logger, cfg.Password)

	srv := server.NewServer(logger, cfg.Port, svc, cfg.Password)

	err = srv.Start()
	if err != nil {
		logger.Fatal(err)
	}
}
