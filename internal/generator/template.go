package generator

const fnTmpl = `
		func (%method_signature%) {{.MethodName}} ( {{range .}} {{.Name}} {{.Type}} {{end}})
`
