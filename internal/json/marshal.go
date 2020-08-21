package json

import (
	"bytes"
	"text/template"

	"github.com/b00lduck/rest-websocket-tester/internal/log"
)

type JsonData struct {
	Topic     string
	Timestamp int64
	Message   string
}

type marshaller struct {
	logger       log.SugaredLogger
	jsonTemplate TemplateI
}

func NewMarshaller(logger log.SugaredLogger) *marshaller {
	return &marshaller{
		logger:       logger,
		jsonTemplate: template.Must(template.New("jsonMessage").Parse(jsonTemplate))}
}

func (m *marshaller) Marshal(message []byte) ([]byte, error) {
	data := JsonData{
		Message: string(message),
	}

	var buf bytes.Buffer
	err := m.jsonTemplate.Execute(&buf, data)
	if err != nil {
		m.logger.Error("error creating JSON", err)
		return nil, err
	}
	return buf.Bytes(), nil
}

const jsonTemplate = `{
    "type": "BROKER/BLAH",
    "payload": {
      "message": {{.Message}}
    }
}`
