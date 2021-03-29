package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const usage = `terrassist [options] <pkg> <type expression>

Given a type "Foo", <type expression> can be one of:

  - Foo
  - *Foo
  - map[string]Foo
  - map[string]*Foo
  - *map[string]Foo
  - *map[string]*Foo
  - []Foo
  - []*Foo
  - *[]Foo
  - *[]*Foo

Options:
`

func main() {
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

	ctx, err := NewCtx(CtxOptions{
		Dir:      ".",
		PkgName:  pkgName,
		TypeExpr: typeName,
	}, flags{honorJSONIgnore: *honorJSONIgnore})
	if err != nil {
		log.Fatal(err)
	}
	f := ctx.run()

	if err := f.Render(os.Stdout); err != nil {
		log.Fatalf("failed to render: %v", err)
	}
}
