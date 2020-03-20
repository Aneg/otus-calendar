package main

import (
	"flag"
	"github.com/Aneg/calendar/internal/config"
	"github.com/Aneg/calendar/internal/web"
	log2 "github.com/Aneg/calendar/pkg/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
)

func main() {
	config, err := config.GetConfigFromFile(getConfigDir())
	if err != nil {
		log.Fatal(err)
	}

	log2.Logger, err = getLogger(config.LogFile, config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	web.Init(config)
	if err := http.ListenAndServe(config.HttpListen, web.Router); err != nil {
		log2.Logger.Fatal(err.Error())
	}
}

func getLogger(logFile, logLevel string) (logger *zap.Logger, err error) {
	var level zapcore.Level
	switch logLevel {
	case "debug":
		level = zapcore.DebugLevel
		break
	case "info":
		level = zapcore.InfoLevel
		break
	case "warn":
		level = zapcore.WarnLevel
		break
	case "error":
		level = zapcore.ErrorLevel
	}

	return zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(level),
		OutputPaths: []string{"stdout", logFile},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "message", // <--
		},
	}.Build()
}

func getConfigDir() string {
	var configDir string
	flag.StringVar(&configDir, "config", "configs/config.yaml", "path to config file")
	flag.Parse()
	return configDir
}
