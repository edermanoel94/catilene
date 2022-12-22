package scan

import "fmt"

func Run() {

}

func report(extensionStats map[string]int) {
	total := 0
	for _, v := range extensionStats {
		total += v
	}
	fmt.Printf("Total de arquivos é de: %d \n", total)
}
