package main

import (
	"fmt"
	"os"

	"github.com/routefusion/routefusion-golang/internal/generator"
)

func main() {
	f, err := os.Open("api.go")
	fmt.Println(err)

	o, err := os.Create("methods.go")
	fmt.Println(err)
	g := generator.NewAST(f, o)
	w := &generator.ASTWriter{PackageName: "routefusion"}
	r := &generator.ASTReader{}
	g.Generate(r, w)
}
