package broker

import (
	"errors"
	"testing"
	"time"

	mock_broker "github.com/b00lduck/rest-websocket-tester/internal/broker/mocks"
	"github.com/b00lduck/rest-websocket-tester/internal/dto"
	mock_log "github.com/b00lduck/rest-websocket-tester/internal/log/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewBroker(t *testing.T) {

	// given
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	broadcasterMock := mock_broker.NewMockBroadcaster(ctrl)
	loggerMock := mock_log.NewMockSugaredLogger(ctrl)

	// when
	testSubject := NewBroker(loggerMock, broadcasterMock)

	// then
	a.NotNil(testSubject.marshaller)
	a.NotNil(testSubject.broadcaster)
	a.NotNil(testSubject.messageQueue)
}

func TestRun(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	testPayload := []byte("msgMessage")
	jsonMessage := []byte("jsonMessage")
	testMessage := dto.Message{
		Text: jsonMessage}

	// and
	marshallerMock := mock_broker.NewMockMarshaller(ctrl)
	marshallerMock.EXPECT().Marshal(testPayload).Return(jsonMessage, nil)

	// and
	messageQueueMock := make(chan queuedMessage)

	// and
	broadcasterMock := mock_broker.NewMockBroadcaster(ctrl)
	broadcasterMock.EXPECT().Broadcast(testMessage)

	// and
	loggerMock := mock_log.NewMockSugaredLogger(ctrl)

	// and the test subject
	testSubject := broker{
		logger:       loggerMock,
		broadcaster:  broadcasterMock,
		messageQueue: messageQueueMock,
		marshaller:   marshallerMock}

	// when
	go testSubject.Run()

	// then message is sent
	messageQueueMock <- queuedMessage{
		message: testPayload}

	time.Sleep(500 * time.Millisecond)

}

func TestRunFaultyJson(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	testPayload := []byte("msgMessage")

	// and
	marshallerMock := mock_broker.NewMockMarshaller(ctrl)
	marshallerMock.EXPECT().Marshal(testPayload).Return(nil, errors.New("some error"))

	// and
	messageQueueMock := make(chan queuedMessage)

	// and
	broadcasterMock := mock_broker.NewMockBroadcaster(ctrl)
	broadcasterMock.EXPECT().Broadcast(gomock.Any()).Times(0)

	// and
	loggerMock := mock_log.NewMockSugaredLogger(ctrl)
	loggerMock.EXPECT().Info("error marshalling json")

	// and the test subject
	testSubject := broker{
		logger:       loggerMock,
		broadcaster:  broadcasterMock,
		messageQueue: messageQueueMock,
		marshaller:   marshallerMock}

	// when
	go testSubject.Run()

	// and message is sent
	messageQueueMock <- queuedMessage{
		message: testPayload}
}

func TestMessage(t *testing.T) {

	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	messageQueueMock := make(chan queuedMessage, 1)

	// and
	loggerMock := mock_log.NewMockSugaredLogger(ctrl)

	// and the test subject
	testSubject := broker{
		logger:       loggerMock,
		messageQueue: messageQueueMock}

	// when
	testSubject.Message([]byte("someMsg"))
	res := <-messageQueueMock

	// then
	a.Equal(queuedMessage{message: []byte("someMsg")}, res)
}
