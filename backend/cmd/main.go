package main

import (
	"backend/config"
	"backend/logger"
	"context"
	"flag"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

// run makes it easy to test. For testing purposes, the args could changes so can the function getenv
func run(ctx context.Context, args []string, getenv func(string) string, stdin io.Reader, stdout, stderr io.Writer) error {
	// ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	// defer cancel()

	fs := flag.NewFlagSet("logfile", flag.ExitOnError)
	fileName := fs.String("logfile", "na", "file location for logfile")
	if err := fs.Parse(args); err != nil {
		return err
	}

	var writer io.Writer
	if *fileName == "na" {
		writer = stdout
	} else {
		file, err := os.Open(*fileName)
		defer func() error {
			err := file.Close()
			if err != nil {
				return err
			}
			return nil
		}()
		if err != nil {
			return err
		}
		writer = file
	}

	config := &config.Config{
		Port:     getenv("PORT"),
		Domain:   getenv("DOMAIN"),
		LogLevel: logger.ToLevel(getenv("LOG_LEVEL")),
		LogFile:  &writer,
	}

	logger := logger.New(*config.LogFile)

	server := NewServer(logger, config)
	httpServer := http.Server{
		Addr:    net.JoinHostPort(config.Domain, config.Port),
		Handler: server,
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		logger.Info().Str("addr", httpServer.Addr).Msg("http server listening..")
		if err := httpServer.ListenAndServe(); err != nil {
			logger.
				Panic().
				AnErr("err", err).
				Str("port", config.Port).
				Str("domain", config.Domain).
				Msg("failed to start server")

		}
	}()
	wg.Wait()

	return nil
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv, os.Stdin, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

}
