// Echo2 exibe seus argumentos de linha de comando

package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		fmt.Println(index)
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
