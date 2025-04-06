package scanner 

import(
	"fmt"
	"IAC-Scanner/rules"
)

func ScanTerraform(path string){
	fmt.Printf("Scanning Terraform file: %s\n", path)
	rules.CheckTerraform(path)
}