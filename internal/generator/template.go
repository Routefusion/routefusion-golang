package generator

const fnTmpl = `
		{{$i := separator ", "}}
		{{$o := separator ", "}}
		{{$bodyEmpty := isBodyEmpty .Body}}
		{{$length := len .OutputParams}}
		func (r *Routefusion) {{.MethodName}} ({{range .InputParams}}{{call $i}}{{.Name}} {{.Type}}{{end}}) ( 
		{{range .OutputParams}} {{call $o}} {{.Name}} {{.Type}} {{end}}){
			op := client.Operation{
				HTTPMethod: http.Method{{.Verb}},
				HTTPPath:   r.baseURL + "/{{.Path}}",
			}

			var response {{(index .OutputParams 0).Type}}

			{{if eq $bodyEmpty false}}
			bdy , err := json.Marshal({{.Body.Name}})
			if err != nil {
				return nil, err 
			}
			req, err := r.cl.NewRequest(op, &response, bdy)
			{{ else }}
			req, err := r.cl.NewRequest(op, &response, nil)
			{{ end }}

			if err != nil {
				{{if eq $length 1}} return err {{ else }} return nil, err {{ end }} 
			}
			if err := req.Send(); err != nil {
				{{if eq $length 1}} return err {{ else }} return nil, err {{ end }} 
			}

			{{if eq $length 1}} return nil {{ else }} return response, nil {{ end }} 
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

func isBodyEmpty(body Parameter) bool {
	return body == Parameter{}
}
