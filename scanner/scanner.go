package scanner 

import(
	"fmt"
	"os"
    // "IAC-Scanner/scanner/docker"
    // "IAC-Scanner/scanner/k8s"
    // "IAC-Scanner/scanner/terraform"
	"path/filepath"
)

func RunScanner(root string){
	err := filepath.Walk(root, func	(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		switch filepath.Ext(path){
		case ".tf":
			ScanTerraform(path)
			fmt.Println("This is a terraform file")
		
		case ".yaml", ".yml":
			scanKubernetes(path)
			fmt.Printf("This is a Kubernetes file")
		
		case "":
			if filepath.Base(path) == "Dockerfile"{
				fmt.Println("This is a DockerFile")
				scanDockerFile(path)
			}
		}
		return nilThis is a Kubernetes file
	})

	if err != nil {
		fmt.Printf("Error-scanning %s: %v\n", root, err)
	}
}