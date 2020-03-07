# Genmethods

The api stubs in this repository are code generated. 

## How it works: 

The generator reads through the interface declarations in api.go in root and 
creates stub code based on the following rules.

1. A method has to be annotated with the endpoint and http verb. Where we add 
these endpoint and verbs are still up for debate. They could be on the method
or in a config.json.

2. The input struct gets parsed and sent as part of the request, and the output
struct and error are marshalled from the response.

3. This relies heavily on the client implementation to keep the generated code
clean and readable.

This is a sample of how the generated code will be.

```go 

// interface definition in api.go. *UserDetails needs to be defined as well.
// GetUser
// Get <specify version if not v1>
// /users/me
GetUser() (*UserDetails, error)

// Generated code
func (r *Routefusion) GetUser() (*UserDetails, error) {
	var output = &UserDetails{}
	op := client.Operation{
		HTTPPath:   r.baseURL + v1 + "/users/me",
		HTTPMethod: http.MethodGet,
	}

	req, err := r.cl.NewRequest(op, output, nil)
	if err != nil {
		return nil, err
	}

	if err := req.Send(); err != nil {
		return nil, err
	}

	return output, nil
}
```


## Genmethods has two components:

1. parser.go: Read through a single file(api.go)'s AST and create an intermediate struct for processing.
2. generator.go: Look through the intermediate struct and create method handles for all the interfaces listed in api.go.
