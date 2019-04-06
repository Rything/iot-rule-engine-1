package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/nattaponra/iot-rule-engine/engine"
	api "github.com/nattaponra/iot-rule-engine/engine/api/http"
	"github.com/nattaponra/iot-rule-engine/engine/postgres"
	"github.com/nattaponra/iot-rule-engine/logger"
)

type Config struct {
	LogLevel   string
	DBConfig   postgres.Config
	HTTPPort   string
	ServerCert string
	ServerKey  string
}

func loadConfig() Config {
	return Config{
		LogLevel: "error",
		DBConfig: postgres.Config{
			Host:        "127.0.0.1",
			Port:        "5432",
			User:        "my_user",
			Pass:        "my_password",
			Name:        "engine",
			SSLMode:     "disable",
			SSLCert:     "",
			SSLKey:      "",
			SSLRootCert: "",
		},
		HTTPPort:   "8787",
		ServerCert: "",
		ServerKey:  "",
	}
}

func main() {

	//Load Config
	cfg := loadConfig()

	//Initial Log
	logger, err := logger.New(os.Stdout, cfg.LogLevel)
	if err != nil {
		log.Fatalf(err.Error())
	}

	//Connect DB
	db := connectToDB(cfg.DBConfig, logger)
	defer db.Close()

	//New Service
	repo := postgres.New(db)
	svc := engine.New(repo)

	//Declare error channel to recsive error from goroutine calling
	errs := make(chan error, 2)

	//Start serve HTTP
	go ServeHTTP(svc, cfg.HTTPPort, logger, errs)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	err = <-errs
	logger.Error(fmt.Sprintf("Users service terminated: %s", err))
}

func connectToDB(dbConfig postgres.Config, logger logger.Logger) *sql.DB {
	db, err := postgres.Connect(dbConfig)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to connect to postgres: %s", err))
		os.Exit(1)
	}
	return db
}

func ServeHTTP(svc engine.Service, port string, logger logger.Logger, errs chan error) {
	errs <- http.ListenAndServe(fmt.Sprintf(":%s", port), api.MakeHandler(svc, logger))
}
