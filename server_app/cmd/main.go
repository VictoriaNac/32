package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	user_app "finish"
	"finish/server_app/interial/hendler"
	"finish/server_app/interial/repository"
	"finish/server_app/interial/service"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func main() {

	port := flag.String("port", "8080", "string")
	flag.Parse()

	initLogging()

	if err := initConfig(); err != nil {
		log.Err(err).Msg("Configuration file not loaded")
	}

	log.Info().Msg("Connecting to the database")
	db, err := repository.NewDataBase()
	if err != nil {
		log.Error().Msg("Connection error")
	}
	log.Info().Msg("Connected to the database")

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	hendler := hendler.NewHendler(services)

	log.Info().Msg("Starting server")
	server := new(user_app.Server)
	go func() {
		if err := server.Run(*port, hendler.InitRouters()); err != nil {
			log.Err(err).Msg("Server is not running")
		}
	}()
	log.Info().Msg("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-quit

	log.Warn().Msg("Server shutting down...")

	log.Info().Msg("Disconnecting the database connection")
	if err := db.Close(); err != nil {
		log.Err(err).Msg("Server did not shut down correctly")
	}
	log.Info().Msg("Database disconnected")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Err(err).Msg("Server did not shut down correctly")
	}
	log.Info().Msg("Bye!")
}

// initConfig() initializes configuration files
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

// initLogging() sets the logging settings
func initLogging() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "02.01.2006 15:04:05"}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Logger()
}