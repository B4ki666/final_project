package main

import (
	"log"
	"os"

	"final_project/internal/server"

	"github.com/joho/godotenv"
)

func main() {
	logger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	port := getPort(logger)
	srv := server.NewServer(logger, port)

	err := srv.Start()
	if err != nil {
		logger.Fatal(err)
	}
}

func getPort(logger *log.Logger) string {

	err := godotenv.Load()
	if err != nil {
		logger.Println(err)
	}

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	return port
}
