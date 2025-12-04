// Package main implements a code generation tool for the gounctional library.
//
// This command-line application generates generic binding functions for the gounctional package.
// Binding functions allow partial application of multi-argument functions by fixing some
// of their arguments, returning a new function with fewer parameters.
//
// The tool generates functions from Bind1OfN to BindNOfN for any number of arguments N,
// where each BindXOfY function binds the first X arguments of a Y-argument function.
// For example:
// - Bind1Of3 fixes the first argument of a 3-argument function, returning a 2-argument function
// - Bind2Of3 fixes the first 2 arguments of a 3-argument function, returning a 1-argument function
// - Bind3Of3 fixes all 3 arguments of a 3-argument function, returning a 0-argument function
//
// Usage:
//
//	go run generate_bind.go -s 1 -e 10 -o gounctional_bind.go
//
// Flags:
//
//	-s: Starting number of arguments (default 1)
//	-e: Ending number of arguments (default 1)
//	-o: Output file (required)
//	-p: Package name (default "gounctional")
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	outputFile  = flag.String("o", "", "Output file")
	startArgs   = flag.Int("s", 1, "Starting number of arguments")
	endArgs     = flag.Int("e", 1, "Ending number of arguments")
	packageName = flag.String("p", "gounctional", "Package name")
)

func main() {
	flag.Parse()

	if *outputFile == "" {
		fmt.Println("Output file is required")
		return
	}

	var file *os.File
	if *outputFile == "-" {
		file = os.Stdout
	} else {
		var err error
		file, err = os.Create(*outputFile)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer func() {
			if err := file.Close(); err != nil {
				fmt.Println("Error closing file:", err)
			}
		}()
	}

	writer := bufio.NewWriter(file)
	if _, err := writer.WriteString("package " + *packageName + "\n\n"); err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	for numArgs := *startArgs; numArgs <= *endArgs; numArgs++ {
		for numBindedArgs := 1; numBindedArgs <= numArgs; numBindedArgs++ {
			if _, err := writer.WriteString(generateBindFunctionFor(numArgs, numBindedArgs)); err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
	}
	if err := writer.Flush(); err != nil {
		fmt.Println("Error flushing file:", err)
		return
	}
}

func generateBindFunctionFor(funcArgs, argsToBind int) string {
	return fmt.Sprintf(`
// Bind%[1]dOf%[2]d binds the first %[1]d arguments of the function f.
func Bind%[1]dOf%[2]d[%[3]s, R any](f func(%[4]s) R, %[5]s) func(%[6]s) R {
	return func(%[7]s) R {
		return f(%[8]s)
	}
}
`,
		argsToBind,                                           //1
		funcArgs,                                             //2
		generateGenericArgsFor(funcArgs),                     //3
		generateTypesForBindedFunction(funcArgs),             //4
		generateArgsForBindFunction(argsToBind),              //5
		generateTypesForResultFunction(funcArgs, argsToBind), //6
		generateArgsForResultFunction(funcArgs, argsToBind),  //7
		generateCallingArgs(funcArgs),                        //8
	)

}

func generateGenericArgsFor(numArgs int) string {
	var genericArgs []string
	for i := 0; i < numArgs; i++ {
		genericArgs = append(genericArgs, fmt.Sprintf("T%d any", i))
	}
	return strings.Join(genericArgs, ", ")
}

func generateTypesForBindedFunction(numArgs int) string {
	var args []string
	for i := 0; i < numArgs; i++ {
		args = append(args, fmt.Sprintf("T%d", i))
	}
	return strings.Join(args, ", ")
}

func generateArgsForBindFunction(numArgs int) string {
	var args []string
	for i := 0; i < numArgs; i++ {
		args = append(args, fmt.Sprintf("arg%d T%d", i, i))
	}
	return strings.Join(args, ", ")
}

func generateTypesForResultFunction(funcArgs, argsToBind int) string {
	var args []string
	for i := argsToBind; i < funcArgs; i++ {
		args = append(args, fmt.Sprintf("T%d", i))
	}
	return strings.Join(args, ", ")
}

func generateArgsForResultFunction(funcArgs, argsToBind int) string {
	var args []string
	for i := argsToBind; i < funcArgs; i++ {
		args = append(args, fmt.Sprintf("arg%d T%d", i, i))
	}
	return strings.Join(args, ", ")
}

func generateCallingArgs(numArgs int) string {
	var args []string
	for i := 0; i < numArgs; i++ {
		args = append(args, fmt.Sprintf("arg%d", i))
	}
	return strings.Join(args, ", ")
}
