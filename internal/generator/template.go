package generator

const fnTmpl = `
		{{$i := separator ", "}}
		{{$o := separator ", "}}
		{{$length := len .OutputParams}}
		func (r *Routefusion) {{.MethodName}} ({{range .InputParams}}{{call $i}}{{.Name}} {{.Type}}{{end}}) ( 
		{{range .OutputParams}} {{call $o}} {{.Name}} {{.Type}} {{end}}){
			op := client.Operation{
				HTTPMethod: http.Method{{.Verb}},
				HTTPPath:   endpoint + "/{{.Path}}",
			}

			{{.Body}}

			var response {{(index .OutputParams 0).Type}}
			req, err := r.cl.NewRequest(op, &response, nil)
			if err != nil {
				{{if eq $length 1}}
					return nil
				{{ end }} 
				{{if eq $length 2}}
					return response, nil
				{{ end }} 
			}
			if err := req.Send(); err != nil {
				{{if eq $length 1}}
					return nil
				{{ end }} 
				{{if eq $length 2}}
					return response, nil
				{{ end }} 
			}


		{{if eq $length 1}}
			return nil
		{{ end }} 

		{{if eq $length 2}}
			return response, nil
		{{ end }} 

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
