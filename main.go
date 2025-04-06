package main 

import (
	"IAC-Scanner/scanner"
	"fmt"
	"os"
)

func main (){
	if len(os.Args) < 2 {
		fmt.Println("Usage: IAC-Scanner <path-to-scan>")
		os.Exit(1)
	}
	path := os.Args[1]
	scanner.RunScanner(path)
}