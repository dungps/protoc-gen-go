package main

import (
	"flag"
	"fmt"
	"github.com/dungps/protoc-gen-go/generator"
	"github.com/dungps/protoc-gen-go/internal/version"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	showVersion = flag.Bool("version", false, "show version")
)

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-gof %v\n", version.String())
		return
	}

	protogen.Options{
		ParamFunc: flag.CommandLine.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if f.Generate {
				generator.GenerateFile(gen, f)
			}
		}

		return nil
	})
}
