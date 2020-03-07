package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	if err := _main(); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}

type Parameter struct {
	Name string
	Type string
}

type Action struct {
	interfaceName string
	methods       []Method
}

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
	fmt.Println(actions)
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
