package mayhem

import (
	"github.com/JakeStrang1/mayhem/config"
	j "github.com/dave/jennifer/jen"
)

func Generate(t config.T) error {
	// What are we generating?
	// An API, and the resulting logic
	// - sign up / login
	// - connect with other users
	// - user profile

	GenerateMain(t)

	return nil
}

func GenerateMain(t config.T) error {
	f := j.NewFile("main")
	f.Func().Id("main").Params().Block(
		j.Qual("fmt", "Println").Call(j.Lit("Hello, world")),
	)
	return f.Save(t.ProjectPath + "/main.go")
}
