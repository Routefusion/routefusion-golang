package main

import (
	"os"

	"github.com/routefusion/routefusion-golang/internal/generator"
)

func main() {
	f, err := os.Open("api.go")
	if err != nil {
		panic(err)
	}

	o, err := os.Create("methods.go")
	if err != nil {
		panic(err)
	}
	g := generator.NewAST(f, o)
	w := &generator.ASTWriter{PackageName: "routefusion"}
	r := &generator.ASTReader{}
	g.Generate(r, w)
}
