package main

import (
	app "GoModDemo/app"
	"fmt"
)

func main() {
	hello()
	c := app.NewComplex(1.0, 2.0)
	fmt.Println(c.Real, c.Imaginary)
}
