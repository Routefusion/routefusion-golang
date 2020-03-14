package generator

const fnTmpl = `
		{{$i := separator ", "}}
		{{$o := separator ", "}}
		func (rf *Routefusion) {{.MethodName}} ({{range .InputParams}}{{call $i}}{{.Name}} {{.Type}}{{end}}) ( 
		{{range .OutputParams}} {{call $o}} {{.Name}} {{.Type}} {{end}}){
			op := client.Operation{
				HTTPMethod: http.Method{{.Verb}},
				HTTPPath:  ""  {{.Endpoint}},
			}

			{{.Body}}

			var response {{(index .OutputParams 0).Type}}
			req, err := rf.NewRequest(op, &response, nil)
			if err != nil {
				return nil, err
			}
			if err := req.Send(); err != nil {
				return nil, err
			}

			return response, nil
		}
`

func separator(s string) func() string {
	i := -1
	return func() string {
		i++
		if i == 0 {
			return ""
		}
		return s
	}
}
