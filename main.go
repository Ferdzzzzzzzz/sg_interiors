package main

import (
	"context"
	"embed"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sg/alert"
	"sg/mux"
	"syscall"

	"github.com/Rockup-Consulting/std/core/web"
)

var (
	//go:embed assets
	assets embed.FS
)

func main() {
	l := log.New(os.Stdout, "| ", log.Ldate|log.Ltime|log.LUTC|log.Lmsgprefix)
	l.Println("starting server")

	if err := run(l); err != nil {
		l.Printf("ERROR: %s", err)
		os.Exit(1)
	} else {
		l.Println("done.")
		os.Exit(0)
	}
}

func run(l *log.Logger) error {
	// ====================================================================
	// Conf

	dev := flag.Bool("dev", false, "run application in dev mode")
	port := flag.Uint("port", 4003, "overwrite the default port")

	flag.Parse()

	// ====================================================================
	// Telegram Bot

	var alertBot alert.Alerter

	if *dev {
		alertBot = alert.LogOnly{}

	} else {
		telegramApiToken := os.Getenv("TELEGRAM_BOT_API_TOKEN")
		if telegramApiToken == "" {
			return errors.New("TELEGRAM_BOT_API_TOKEN key has not been set")
		}

		if bot, err := alert.NewTelegramAlertBot(telegramApiToken, l); err != nil {
			return err
		} else {
			alertBot = bot
		}
	}

	// ====================================================================
	// Server

	m := mux.Init(l, alertBot, assets)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: m,
	}

	// ====================================================================
	// Lifecycle channels

	shutdown := make(chan os.Signal, 1)
	serverErr := make(chan error)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		serverErr <- server.ListenAndServe()
	}()

	if *dev {
		l.Printf("listening at http://localhost:%d\n", *port)
	} else {
		l.Printf("listening at :%d\n", *port)
	}

	select {
	case err := <-serverErr:
		return fmt.Errorf("server error: %s", err)
	case sig := <-shutdown:
		l.Printf("starting server shutdown: %s", sig)
		defer l.Println("server shutdown complete")

		ctx, cancel := context.WithTimeout(context.Background(), web.DefaultShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			server.Close()
			return fmt.Errorf("could not gracefully shutdown server: %s", err)
		}
	}

	return nil
}
