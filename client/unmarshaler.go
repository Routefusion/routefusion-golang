package client

import (
	"encoding/json"
	"io"
)

var unmarshalers = map[string]Unmarshaler{
	"application/json": &jsonUnmarshaler{},
}

// AddUnmarshaler can add any unmarshaler to the list of accepted unmarshalers
// by client.
func AddUnmarshaler(header string, unmarshaler Unmarshaler) {
	unmarshalers[header] = unmarshaler
}

// Unmarshaler is the specification every accept header type needs to implement
// to be able to unmarshal the registered type.
type Unmarshaler interface {
	Unmarshal(r io.Reader, v interface{}) error
}

type jsonUnmarshaler struct{}

func (j *jsonUnmarshaler) Unmarshal(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}
