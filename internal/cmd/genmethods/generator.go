package main

import (
	"fmt"
	"os"

	"github.com/routefusion/routefusion-golang/internal/generator"
)

func main() {
	f, err := os.Open("api.go")
	fmt.Println(err)
	g := generator.NewAST(f, os.Stdout)
	w := &generator.ASTWriter{}
	r := &generator.ASTReader{}
	g.Generate(r, w)
}
