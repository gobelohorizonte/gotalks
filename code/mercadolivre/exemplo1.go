package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type test struct {
	Message string  `json:"message"`
	Error   string  `json:"error"`
	Status  int     `json:"status"`
	Cause   []cause `json:"cause"`
}

type cause struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func createProducts(access_token string) {

	url := fmt.Sprintf("https://api.mercadolibre.com/items?access_token=%s", access_token)

	fmt.Println(url)

	file, err := ioutil.ReadFile("./product.json")
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("GET", url, bytes.NewReader(file))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if resp != nil {
			defer resp.Body.Close()
		}
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Header:", resp.Header)
	fmt.Println("response Body:", resp.Body)

	var t test
	if err := json.NewDecoder(resp.Body).Decode(&t); err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Println(t)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(body)
}

func main() {

	createProducts("APP_")
}
