/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size        2017
 */

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type JsonResp struct {
	Name []string `json:"Name"`
}

func main() {

	//// JSON
	///
	///

	os.Setenv("LANGJSON", `{"name":["C","C++","PHP","Python", "Ruby"]}`)

	Ljson := []byte(os.Getenv("LANGJSON"))

	Ijson := JsonResp{}

	err := json.Unmarshal(Ljson, &Ijson)

	if err != nil {
		fmt.Println("Error", err)
	}

	// Only one field
	fmt.Println(Ijson.Name[1])

	// Listing all
	for key, value := range Ijson.Name {

		fmt.Println("Json Name, Key: ", key, "Value: ", value)
	}
}
