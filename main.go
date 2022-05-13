package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func closure(ch chan int) {
	inc := func() func() int {
		i := 0
		return func() int {
			i++
			return i
		}
	}

	c := inc()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	ch <- 9
}

func main() {
	var res = fmt.Sprintf("test-%s", "go")
	fmt.Println(res)

	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(getSquareRoot(9))

	var c1 Circle
	c1.radius = 10.00
	fmt.Println(c1.area())

	ch := make(chan int)
	go closure(ch)
	<-ch
}
