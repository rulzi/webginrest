package server

import (
	"context"
	"fmt"
	"runtime/debug"
	"strings"
	"time"
	"webginrest/server/config"
	"webginrest/server/providers"
	"webginrest/server/services"

	"github.com/mattn/go-colorable"
	"github.com/snowzach/rotatefilehook"

	log "github.com/sirupsen/logrus"
)

func configureLogging() {
	lLevel := config.Get("server.log.level")
	fmt.Println("Setting log level to ", lLevel)
	switch strings.ToUpper(lLevel) {
	default:
		fmt.Println("Unknown level [", lLevel, "]. Log level set to ERROR")
		log.SetLevel(log.ErrorLevel)
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO":
		log.SetLevel(log.InfoLevel)
	case "WARN":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	}

	currentTime := time.Now()

	rotateFileHook, err := rotatefilehook.NewRotateFileHook(rotatefilehook.RotateFileConfig{
		Filename:   "logs/webginrest-" + currentTime.Format("2006-01-02") + ".log",
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     7, //days
		Level:      log.GetLevel(),
		Formatter: &log.JSONFormatter{
			TimestampFormat: time.RFC3339,
		},
	})

	if err != nil {
		log.Fatalf("Failed to initialize file rotate hook: %v", err)
	}

	log.SetReportCaller(true)
	log.SetLevel(log.GetLevel())
	log.SetOutput(colorable.NewColorableStdout())
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	log.AddHook(rotateFileHook)
}

func recoverPanic() {
	if r := recover(); r != nil {
		log.WithField("panic", r).WithField("stack trace", string(debug.Stack())).Error("we panicked!")
	}
}

// Start start server
func Start() {
	defer recoverPanic()
	configureLogging()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbHost := config.Get("database.host")
	dbPort := config.Get(`database.port`)
	dbUser := config.Get(`database.user`)
	dbPass := config.Get(`database.pass`)
	dbName := config.Get(`database.name`)
	services.OpenDBConnection(ctx, dbUser, dbPass, dbHost, dbPort, dbName)
	defer services.CloseDBConnection()

	redisHost := config.Get(`redis.host`)
	redisPort := config.Get(`redis.port`)

	services.OpenRedisConnection(redisHost, redisPort)

	defer services.CloseDBConnection()

	mailHost := config.Get(`mail.host`)
	mailPort := config.GetInt(`mail.port`)
	mailUser := config.Get(`mail.name`)
	mailPass := config.Get(`mail.pass`)

	services.OpenEmailClient(ctx, mailHost, mailPort, mailUser, mailPass)

	route := providers.Route(ctx)

	route.Run(config.Get(`server.address`))
}
