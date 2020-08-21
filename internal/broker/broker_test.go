package broker

import (
	"testing"
	"time"

	mock_broker "github.com/b00lduck/rest-websocket-tester/internal/broker/mocks"
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
	a.NotNil(testSubject.broadcaster)
	a.NotNil(testSubject.messageQueue)
}

func TestRun(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	testPayload := []byte("msgMessage")

	// and
	messageQueueMock := make(chan []byte)

	// and
	broadcasterMock := mock_broker.NewMockBroadcaster(ctrl)
	broadcasterMock.EXPECT().Broadcast(testPayload)

	// and
	loggerMock := mock_log.NewMockSugaredLogger(ctrl)

	// and the test subject
	testSubject := broker{
		logger:       loggerMock,
		broadcaster:  broadcasterMock,
		messageQueue: messageQueueMock}

	// when
	go testSubject.Run()

	// then message is sent
	messageQueueMock <- testPayload

	time.Sleep(500 * time.Millisecond)

}

func TestMessage(t *testing.T) {

	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	messageQueueMock := make(chan []byte, 1)

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
	a.Equal([]byte("someMsg"), res)
}
