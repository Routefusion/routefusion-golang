package generator

import "io"

// Parameter is a parameter attached to a function.
type Parameter struct {
	Name string
	Type string
}

// API is a set of indicatives that we will loop through
// to generate code.
type API struct {
	InterfaceName string
	Methods       []Method
}

// Method represents a set of metadata levels details that belong
// to an api method.
type Method struct {
	Endpoint     string
	MethodName   string
	Path         string
	Verb         string
	Body         string
	InputParams  []Parameter
	OutputParams []Parameter
}

// Reader is a general wrapper on io.Reader to read an input to obtain
// the apis.
type Reader interface {
	ReadAPI(io.Reader) ([]API, error)
}

// Writer is a general wrapper on io.Writer to write the processed file.
type Writer interface {
	WriteAPI(io.Writer, []API) error
}

// Generator is an interface that specifies how a generation would occur.
type Generator interface {
	Generate(Reader, Writer) error
}
