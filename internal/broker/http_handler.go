package broker

import (
	"io/ioutil"
	"net/http"

	"github.com/b00lduck/rest-websocket-tester/internal/log"
)

func MessageHandler(logger log.SugaredLogger, broker Broker, w http.ResponseWriter, r *http.Request) {

	message, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error("error reading http request body", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(message) < 1 {
		logger.Error("missing message")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	logger.Info("Received new message",
		"messageLength", len(message))
	broker.Message(message)

	w.WriteHeader(http.StatusCreated)
}
