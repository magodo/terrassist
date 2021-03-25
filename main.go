package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const usage = "terrassist <pkg> <type>"

func main() {
	forPointer := flag.Bool("p", false, `Whether to generate for the pointer of the specified type?`)
	honorJSONIgnore := flag.Bool("j", false, `Ignore struct field that has json tag "-" specified`)

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s\n", usage)
		flag.PrintDefaults()
	}
	flag.Parse()

	if len(flag.Args()) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	pkgName, typeName := flag.Args()[0], flag.Args()[1]

	ctx := Ctx{
		existFuncs: map[string]bool{},
		options: options{
			honorJSONIgnore: *honorJSONIgnore,
			forPointer:      *forPointer,
		},
	}
	f := ctx.run(".", pkgName, typeName)

	if err := f.Render(os.Stdout); err != nil {
		log.Fatalf("failed to render: %v", err)
	}
}
