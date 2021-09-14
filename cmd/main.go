package main

import (
	"context"
	"fmt"
	"gazuberlandia/handler"
	"gazuberlandia/postgres"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v4"
	_ "github.com/joho/godotenv/autoload"
)

const (
	urlConn = "postgres://docker:docker@localhost:5432/guberlandia?sslmode=disable"
)

func main() {
	conn, err := postgres.Open(urlConn)

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	defer conn.Close()

	server := handler.NewServer(
		handler.NewServices(conn),
		handler.NewRouter(),
	)

	http.ListenAndServe(":5000", server)

}

func run() error {
	conn, err := postgres.Open(urlConn)

	if err != nil {
		return err
	}
	defer conn.Close()

	h := handler.NewHandler(conn)

	srv := http.Server{
		Handler:  h,
		Addr:     ":5000",
		ErrorLog: log.New(os.Stderr, "LoggerError: ", log.Lshortfile),
	}

	go func() {
		err = srv.ListenAndServe()

		if err != nil {
			log.Fatal("Error to setup server. Error: ", err)
			os.Exit(1)
		}
	}()

	log.Printf(" ==> Server started <== Port %s", srv.Addr)

	downServerChan := make(chan os.Signal, 1)

	signal.Notify(downServerChan, os.Interrupt, syscall.SIGTERM)

	sig := <-downServerChan
	log.Println("Server down, receive signal ==> ", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	srv.Shutdown(ctx)
	fmt.Println()
	log.Println(" <=== Shutting down Server ===>")
	os.Exit(0)

	return nil
}

// if err := run(); err != nil {
// 	log.Println("Error running server..", err)
// 	os.Exit(1)
// }

// conn, err := postgres.Open(urlConn)

// if err != nil {
// 	fmt.Println("Error", err)
// 	os.Exit(1)
// }
// defer conn.Close()

// userRepository := postgres.NewUserRepository(conn)

// server, err := handler.NewServer(
// 	handler.NewConfigUserHandler(userRepository),
// )

// if err != nil {
// 	fmt.Println("Error", err)
// }
