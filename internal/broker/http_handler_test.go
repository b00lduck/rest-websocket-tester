package broker

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	mock_broker "github.com/b00lduck/rest-websocket-tester/internal/broker/mocks"
	mock_log "github.com/b00lduck/rest-websocket-tester/internal/log/mocks"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestMessageHandler(t *testing.T) {

	// given
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// and
	brokerMock := mock_broker.NewMockBroker(ctrl)
	brokerMock.EXPECT().Message([]byte("testMessage"))

	// and
	loggerMock := mock_log.NewMockSugaredLogger(ctrl)
	loggerMock.EXPECT().Info("Received new message", "messageLength", 11)

	// and custom test router
	r := mux.NewRouter()
	r.HandleFunc("/{topic}", func(w http.ResponseWriter, r *http.Request) {
		MessageHandler(loggerMock, brokerMock, w, r)
	}).Methods(http.MethodPost)
	req, _ := http.NewRequest("POST", "/testTopic", strings.NewReader("testMessage"))
	rec := httptest.NewRecorder()

	// when
	r.ServeHTTP(rec, req)
	res, _ := ioutil.ReadAll(rec.Body)

	// then
	a.Equal([]byte{}, res)
	a.Equal(201, rec.Code)
}

func TestMessageHandlerMissingMessage(t *testing.T) {

	// given
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// and
	brokerMock := mock_broker.NewMockBroker(ctrl)

	// and
	loggerMock := mock_log.NewMockSugaredLogger(ctrl)
	loggerMock.EXPECT().Error("missing message")

	// and custom test router
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		MessageHandler(loggerMock, brokerMock, w, r)
	}).Methods(http.MethodPost)
	req, _ := http.NewRequest("POST", "/", strings.NewReader(""))
	rec := httptest.NewRecorder()

	// when
	r.ServeHTTP(rec, req)
	res, _ := ioutil.ReadAll(rec.Body)

	// then
	a.Equal([]byte{}, res)
	a.Equal(400, rec.Code)
}

func TestMessageHandlerReadError(t *testing.T) {

	// given
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// and
	loggerMock := mock_log.NewMockSugaredLogger(ctrl)
	loggerMock.EXPECT().Error("error reading http request body", errors.New("some error"))

	// and custom test router
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		MessageHandler(loggerMock, nil, w, r)
	}).Methods(http.MethodPost)
	req, _ := http.NewRequest("POST", "/", ErroneousReader{})
	rec := httptest.NewRecorder()

	// when
	r.ServeHTTP(rec, req)
	res, _ := ioutil.ReadAll(rec.Body)

	// then
	a.Equal([]byte{}, res)
	a.Equal(500, rec.Code)
}

type ErroneousReader struct{}

func (e ErroneousReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("some error")
}
