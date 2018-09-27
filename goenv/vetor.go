/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// semicolon
	//
	//
	os.Setenv("LANG", "C,C++,PHP,Python, Ruby")

	LANG := os.Getenv("LANG")

	Vetor := strings.Split(LANG, ",")

	fmt.Println("LANG:", len(Vetor), Vetor[3])
	fmt.Println("------------------------------------")

	// Listing all
	for key, value := range Vetor {

		fmt.Println("Vetor, Key: ", key, "Value: ", value)
	}
}
