package main

import (
	"fmt"
	"io"
	"os"

	"github.com/killthebuddh4/gecko/core"
	"github.com/killthebuddh4/gecko/eval"
	"github.com/killthebuddh4/gecko/lex"
	"github.com/killthebuddh4/gecko/parse"
)

func main() {

	args := os.Args

	if len(args) != 2 {
		fmt.Println("Usage: gecko [script]")
		fmt.Println(len(args))
	} else {
		exec(args[1])
	}
}

func exec(pathToFile string) {
	source := ""

	_, excludeLib := os.LookupEnv("GECKO_EXCLUDE_LIB")

	lib := []string{"lib/math.g", "lib/array.g", "lib/map.g", "lib/schema.g"}

	files := []string{pathToFile}

	if !excludeLib {
		files = append(lib, files...)
	}

	for _, file := range files {
		f, err := os.Open(file)

		if err != nil {
			fmt.Println("Error opening file: ", err)
			return
		}

		defer f.Close()

		data, err := io.ReadAll(f)

		if err != nil {
			fmt.Println("Error reading file: ", err)
			return
		}

		source += "\n"
		source += "\n"
		source += string(data)
	}

	lexemes, err := lex.Lex(source)

	if err != nil {
		fmt.Println("Error lexing: ", err)
		return
	}

	rootOperator := core.Operator{
		Type:  "program",
		Value: "program",
	}

	if err != nil {
		fmt.Println("Error creating root operator: ", err)
		return
	}

	rootExp := core.Expression{
		Parent:   nil,
		Operator: rootOperator,
		Signature: &core.Signature{
			Parameters: make(map[string]string),
			Returns:    []*core.Expression{},
		},
		Arguments:    map[string][]*core.Expression{},
		Catches:      []core.Block{},
		Trajectories: []*core.Trajectory{},
	}

	parseErr := parse.Parse(&rootExp, lexemes)

	root := core.NewTrajectory(nil, &rootExp)

	sString, err := eval.SchemaString(&root)

	if err != nil {
		fmt.Println("Error creating string schema: ", err)
		return
	}

	sNumber, err := eval.SchemaNumber(&root)

	if err != nil {
		fmt.Println("Error creating number schema: ", err)
		return
	}

	sBoolean, err := eval.SchemaBoolean(&root)

	if err != nil {
		fmt.Println("Error creating boolean schema: ", err)
		return
	}

	sArray, err := eval.SchemaArray(&root)

	if err != nil {
		fmt.Println("Error creating array schema: ", err)
		return
	}

	sHash, err := eval.SchemaHash(&root)

	if err != nil {
		fmt.Println("Error creating hash schema: ", err)
		return
	}

	sFunction, err := eval.SchemaFunction(&root)

	if err != nil {
		fmt.Println("Error creating function schema: ", err)
		return
	}

	sIdentity, err := eval.SchemaIdentity(&root)

	if err != nil {
		fmt.Println("Error creating identity schema: ", err)
		return
	}

	core.DefineName(&root, "String", sString)
	core.DefineName(&root, "Number", sNumber)
	core.DefineName(&root, "Boolean", sBoolean)
	core.DefineName(&root, "Array", sArray)
	core.DefineName(&root, "Hash", sHash)
	core.DefineName(&root, "Function", sFunction)
	core.DefineName(&root, "Identity", sIdentity)

	if parseErr != nil {
		fmt.Println("Error parsing: ", parseErr)
		return
	}

	_, err = eval.Exec(nil, &root, root.Expression)

	if err != nil {
		fmt.Println("Gecko :: Error evaluating: ", err)
		return
	}

}
