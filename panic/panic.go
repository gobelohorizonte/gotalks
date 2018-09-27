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
	"runtime"
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

	hello()
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func GoSafely(fn func()) {

	go func() {

		defer func() {

			if err := recover(); err != nil {

				stack := make([]byte, 1024*8)
				stack = stack[:runtime.Stack(stack, false)]

				f := "PANIC: %s\n%s"
				//logger.Logger.Error().Printf(f, err, stack)
				fmt.Println("oikkkkk", f, err, stack)
			}
		}()

		fn()
	}()
}

func hello() {

	fmt.Println("estou aqui")

	GoSafely(func() {
		panic("hi")
	})
}
