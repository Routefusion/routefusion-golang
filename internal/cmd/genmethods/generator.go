package main

import (
	"bytes"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
)

func main() {
	if err := _main(); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}

// Parameter is a parameter attached to a function.
type Parameter struct {
	Name string
	Type string
}

// Action is a set of indicatives that we will loop through
// to generate code.
type Action struct {
	interfaceName string
	methods       []Method
}

// Method represents a set of metadata levels details that belong
// to an api method.
type Method struct {
	endpoint     string
	methodName   string
	inputParams  []Parameter
	outputParams []Parameter
}

func _main() error {
	actions, err := getActionsFromAST()
	if err != nil {
		return err
	}

	f, err := os.Create("methods.go")
	if err != nil {
		return err
	}
	defer f.Close()

	if err := generateAPIMethods(f, actions); err != nil {
		return err
	}
	return nil
}

// This function is spaceship code to do a specific action. Look
// through type Interface declarations, for methods attached to them,
// extract their params and create the Action type.
func getActionsFromAST() ([]Action, error) {
	fs := token.NewFileSet()
	p, err := parser.ParseFile(fs, "api.go", nil, parser.AllErrors)
	if err != nil {
		return nil, err
	}

	var actions = make([]Action, 0)
	// loop through all the top level declarations of the AST.
	for _, decl := range p.Decls {
		// We don't care about functional declarations, so lets pick
		// only the generic ones.
		decl, ok := decl.(*ast.GenDecl)
		// We don't really need to process token.TYPE.
		if !ok || decl.Tok != token.TYPE {
			continue
		}

		var action Action
		for _, spec := range decl.Specs {
			//We only want to read through type declarations.
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			action.interfaceName = typeSpec.Name.String()

			// Again, we are only interested in Interface definitions.
			interfaceType, ok := typeSpec.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}

			action.methods = getMethods(interfaceType.Methods.List)
			actions = append(actions, action)
		}
	}
	return actions, nil
}

func getMethods(fields []*ast.Field) []Method {
	var methods = make([]Method, 0)
	for _, method := range fields {
		if len(method.Names) == 0 {
			continue
		}

		for _, name := range method.Names {
			var methodData Method
			fn, ok := method.Type.(*ast.FuncType)
			if !ok {
				continue
			}

			methodData.methodName = name.String()
			methodData.inputParams = getFields(fn.Params.List)
			methodData.outputParams = getFields(fn.Results.List)

			methods = append(methods, methodData)
		}
	}
	return methods
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

func generateAPIMethods(w io.Writer, actions []Action) error {
	var buf bytes.Buffer
	buf.WriteString("package routefusion\n")
	for _, action := range actions {
		for _, method := range action.methods {
			writeFunction(&buf, method)
		}
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	if _, err := w.Write(formatted); err != nil {
		return err
	}

	return nil
}

func writeFunction(buf *bytes.Buffer, method Method) {
	buf.WriteString("func ")
	buf.WriteString(method.methodName)
	buf.WriteString("(")
	for i, inputParam := range method.inputParams {
		buf.WriteString(inputParam.Name + " " + inputParam.Type)
		if i < len(method.inputParams)-1 {
			buf.WriteString(",")
		}
	}
	buf.WriteString(")")

	if len(method.outputParams) > 1 {
		buf.WriteString("(")
	}

	for i, outputParam := range method.outputParams {
		buf.WriteString(outputParam.Name + " " + outputParam.Type)
		if i < len(method.outputParams)-1 {
			buf.WriteString(",")
		}
	}

	if len(method.outputParams) > 1 {
		buf.WriteString(")")
	}

	buf.WriteString("{\n")

	buf.WriteString("}\n\n")
}
