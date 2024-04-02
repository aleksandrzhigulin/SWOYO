package main

import (
	"SWOYO/db"
	"SWOYO/handler"
	"SWOYO/services"
	"SWOYO/storage"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Если есть аргумент -d, то сохранение данных идёт в postgresql
	// иначе - в hashmap чтоб узнавать есть ли там url за O(1)
	argument := ""
	err := godotenv.Load()
	if err != nil {
		return
	}
	if len(os.Args[1:]) > 0 {
		argument = os.Args[1]
	}
	address := ":8080" // Порт, на котором будет работать сервис
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var httpHandler http.Handler

	// Подключение к базе данных если есть -d, в любом случае управление сохранением данных
	// делегируется StorageService
	if argument == "-d" {
		dbUser, dbPassword, dbName, dbHost, dbPort :=
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PORT")
		database, err := db.Init(dbUser, dbPassword, dbName, dbHost, dbPort)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer database.Connection.Close()
		httpHandler = handler.NewHandler(database)
		services.StorageServiceInit(database, true)
	} else {
		storage.StorageInit()
		httpHandler = handler.NewHandlerWithoutDb()
		services.StorageServiceInit(db.Database{}, false)
	}

	server := &http.Server{
		Handler: httpHandler,
	}

	go func() {
		server.Serve(listener)
	}()
	defer stop(server)
	fmt.Println("Server started on port ", address)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping")

}

func stop(s *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		fmt.Println("Can't shutdown correctly")
		os.Exit(1)
	}
}
