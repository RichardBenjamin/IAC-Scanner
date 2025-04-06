package scanner 

import(
	"fmt"
	"IAC-Scanner/rules"
)

func ScanTerraform(path string){
	fmt.Println("Scanning Terraform file: %s\n", path)
	rules.CheckTerraform(path)
}