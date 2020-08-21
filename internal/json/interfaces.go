package json

//go:generate mockgen -source=interfaces.go -destination=mocks/json.go -package=mock_json

import (
	"io"
)

/*
type Marshaller interface {
	Marshal(message []byte) ([]byte, error)
}
*/

type TemplateI interface {
	Execute(io.Writer, interface{}) error
}
