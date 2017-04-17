/*
* Golang presentation
*
* @package     main
* @author      @jeffotoni
* @size		   2017
 */

package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		str := recover()
		fmt.Println("Recover: ", str)
	}()

	_, errf := os.Stat("file.txt")

	if errf == nil {

		os.Remove("file.txt")
	}

	var y, x int64

	f, err := os.OpenFile("file.txt", os.O_RDWR|os.O_CREATE, 0755)
	defer f.Close()

	if err != nil {

		panic(err)
	}

	size, errx := f.Write([]byte("test text!!!"))

	if errx != nil {

		panic(errx)
	}

	fmt.Println(size, errx, err)

	x = 10
	y = 0

	fmt.Println(x, y)

	fmt.Println("Fatorial: ", factorial(4))

}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
