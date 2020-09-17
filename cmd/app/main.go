package main

import (
	"net/http"
	"os"

	"github.com/b00lduck/rest-websocket-tester/internal/broker"
	"github.com/b00lduck/rest-websocket-tester/internal/ws"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {

	rawlogger, _ := zap.NewDevelopment()

	logger := rawlogger.Sugar()

	logger.Info("Starting websocket test server")

	listen, found := os.LookupEnv("LISTEN")
	if !found {
		listen = ":8080"
	}

	logfile, found := os.LookupEnv("LOGFILE")
	if !found {
		logfile = ""
	}

	hub := ws.NewHub(logger)
	go hub.Run()

	b := broker.NewBroker(logger, hub)
	go b.Run()

	r := mux.NewRouter()

	// Root websocket handler
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws.Handler(logger, hub, b, w, r, logfile)
	}).Methods(http.MethodGet)

	// Post new message handler
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		broker.MessageHandler(logger, b, w, r)
	}).Methods(http.MethodPost)

	logger.Infow("Start to listen",
		"listenAddress", listen)

	err := http.ListenAndServe(listen, r)
	if err != nil {
		logger.Fatal("Error listening",
			"listenAddress", listen,
			"error", err)
	}

	logger.Info("Exiting websocket test server")
}
