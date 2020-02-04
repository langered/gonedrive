package main

import (
	"fmt"
	"github.com/langered/gonedrive/cmd"
	"github.com/spf13/cobra/doc"
)

func main() {
	fmt.Println("Generating new documentation")
	gonedriveCLI := cmd.NewRootCmd()
	err := doc.GenMarkdownTree(gonedriveCLI, "./doc")
	if err != nil {
		fmt.Println(err)
	}
}
