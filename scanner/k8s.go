package scanner 

import(
	"fmt"
	"IAC-Scanner/rules"
)

func scanKubernetes(path string){
	fmt.Printf("Scanning Kubernetes file: %s\n", path)
	// rules.CheckKubernetes(path)
	rules.CheckKubernetesYAML(path)
}


