package main

import (
	"flag"
	"github.com/Aneg/calendar/internal/config"
	"github.com/Aneg/calendar/internal/repositories"
	"github.com/Aneg/calendar/internal/web"
	log2 "github.com/Aneg/calendar/pkg/log"
	calendar "github.com/Aneg/calendar/proto"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

func main() {
	conf, err := config.GetConfigFromFile(getConfigDir())
	if err != nil {
		log.Fatal(err)
	}

	log2.Logger, err = getLogger(conf.LogFile, conf.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	web.Init(conf)

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log2.Logger.Fatal(err.Error())
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	calendarServer := &web.CalendarServer{CalendarRepository: repositories.NewCalendarMap()}

	calendar.RegisterCalendarCheckerServer(grpcServer, calendarServer)
	grpcServer.Serve(lis)

	if err := http.ListenAndServe(conf.HttpListen, web.Router); err != nil {
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