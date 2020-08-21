package json

import (
	"errors"
	"io"
	"testing"

	mock_log "github.com/b00lduck/rest-websocket-tester/internal/log/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNewMarshaller(t *testing.T) {

	// given
	a := assert.New(t)

	// when
	testSubject := NewMarshaller(nil)

	// then
	a.NotNil(testSubject.jsonTemplate)
}

func TestMarshal(t *testing.T) {

	// given
	a := assert.New(t)
	testSubject := NewMarshaller(nil)

	// when
	res, _ := testSubject.Marshal([]byte("\"hello\""))

	// then
	a.Equal([]byte(expectedJson), res)
}

func TestMarshalBadJson(t *testing.T) {

	// given
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	loggerMock := mock_log.NewMockSugaredLogger(ctrl)
	loggerMock.EXPECT().Error("error creating JSON", errors.New("some error"))

	// and
	testSubject := NewMarshaller(loggerMock)
	testSubject.jsonTemplate = ErroneousTemplate{}

	// when
	res, err := testSubject.Marshal([]byte("\"hello\""))

	// then
	a.Nil(res)
	a.Error(err, "some error")
}

const expectedJson = `{
    "type": "BROKER/BLAH",
    "payload": {
      "message": "hello"
    }
}`

type ErroneousTemplate struct{}

func (e ErroneousTemplate) Execute(io.Writer, interface{}) error {
	return errors.New("some error")
}
