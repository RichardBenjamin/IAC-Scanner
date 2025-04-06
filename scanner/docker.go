package scanner

import (
    "fmt"
    "IAC-Scanner/rules"
)

func scanDockerFile(path string){
	fmt.Printf("Scanning Docker file: %s\n", path)
	rules. CheckDockerfile(path)
}