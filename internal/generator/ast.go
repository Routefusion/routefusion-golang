package generator

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"regexp"
	"strings"
	"text/template"
)

var verbEndpointRegex = regexp.MustCompile(`Get|Put|Post|Delete|v1\/.*`)

type AST struct {
	reader io.Reader
	writer io.Writer
}

func NewAST(r io.Reader, w io.Writer) *AST {
	return &AST{
		reader: r,
		writer: w,
	}
}

func (ar *AST) Generate(r Reader, w Writer) error {
	api, err := r.ReadAPI(ar.reader)
	if err != nil {
		return err
	}

	if err := w.WriteAPI(ar.writer, api); err != nil {
		return err
	}

	return nil
}

type ASTReader struct{}

func (ar *ASTReader) ReadAPI(r io.Reader) ([]API, error) {
	fs := token.NewFileSet()
	p, err := parser.ParseFile(fs, "", r, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var apis = make([]API, 0)
	// loop through all the top level declarations of the AST.
	for _, decl := range p.Decls {
		// We don't care about functional declarations, so lets pick
		// only the generic ones.
		decl, ok := decl.(*ast.GenDecl)
		// We don't really need to process token.TYPE.
		if !ok || decl.Tok != token.TYPE {
			continue
		}

		var api API
		for _, spec := range decl.Specs {
			//We only want to read through type declarations.
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			api.InterfaceName = typeSpec.Name.String()

			// Again, we are only interested in Interface definitions.
			interfaceType, ok := typeSpec.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}

			api.Methods = getMethods(interfaceType.Methods.List)
			apis = append(apis, api)
		}
	}

	return apis, nil
}

func getMethods(fields []*ast.Field) []Method {
	var methods = make([]Method, 0)
	for _, method := range fields {
		if len(method.Names) == 0 {
			continue
		}

		var ignore bool
		var matches [][]string
		for _, d := range method.Doc.List {
			comment := extractComment(d.Text)
			if comment == "Ignore" {
				ignore = true
			} else {
				matches = verbEndpointRegex.FindAllStringSubmatch(comment, -1)
			}
		}
		if ignore {
			continue
		}

		fn, ok := method.Type.(*ast.FuncType)
		if !ok {
			continue
		}

		var methodData Method
		methodData.Verb = matches[0][0]
		methodData.Path = matches[1][0]
		methodData.MethodName = method.Names[0].String()
		methodData.InputParams = getFields(fn.Params.List)
		methodData.OutputParams = getFields(fn.Results.List)

		methods = append(methods, methodData)
	}
	return methods
}

func extractComment(comment string) string {
	slashRemoved := strings.Trim(comment, "//")
	return strings.Trim(slashRemoved, " ")
}

func getFields(fields []*ast.Field) []Parameter {
	params := make([]Parameter, 0)
	for _, field := range fields {
		var name string
		if len(field.Names) > 0 {
			name = field.Names[0].Name
		}
		switch t := field.Type.(type) {
		case *ast.SelectorExpr:
			p := Parameter{
				Name: name,
				Type: selectorToString(t),
			}
			params = append(params, p)
		case *ast.StarExpr:
			p := Parameter{
				Name: name,
				Type: starToString(t),
			}
			params = append(params, p)
		case *ast.Ident:
			p := Parameter{
				Name: name,
				Type: t.Name,
			}
			params = append(params, p)
		}

	}
	return params
}

func selectorToString(t *ast.SelectorExpr) string {
	var param bytes.Buffer
	if ident, ok := t.X.(*ast.Ident); ok {
		param.WriteString(ident.Name)
	}
	param.WriteString(".")
	param.WriteString(t.Sel.Name)
	return param.String()
}

func starToString(t *ast.StarExpr) string {
	var param bytes.Buffer
	param.WriteString("*")
	if ident, ok := t.X.(*ast.Ident); ok {
		param.WriteString(ident.Name)
	}
	return param.String()
}

type ASTWriter struct {
	PackageName string
}

func (aw *ASTWriter) WriteAPI(w io.Writer, apis []API) error {
	var buf bytes.Buffer
	//s := "rf *routefusion"
	packageName := fmt.Sprintf("package %s", aw.PackageName)
	buf.WriteString(packageName)
	buf.WriteString("\n\n// Auto-generated by internal/cmd/genmethods/generator.go. DO NOT EDIT!\n")
	buf.WriteString("const endpoint = \"https://api.routefusion.co\"")

	for _, api := range apis {
		for _, method := range api.Methods {
			t := template.Must(template.New("func").
				Funcs(template.FuncMap{"separator": separator}).
				Parse(fnTmpl))
			t.Execute(&buf, method)
		}
	}

	//	formattedData, err := format.Source(buf.Bytes())
	//	if err != nil {
	//		return err
	//	}

	formattedData := buf.Bytes()

	if _, err := w.Write(formattedData); err != nil {
		return err
	}

	return nil
}
