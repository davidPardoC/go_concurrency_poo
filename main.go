package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	var x int
	x = 8
	y := 7
	fmt.Println(x, y)
	entero, error := strconv.ParseInt("123", 10, 64)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(entero)
	m := make(map[string]int)
	m["uno"] = 1
	fmt.Println(m)
	s := []int{1, 2, 3}

	for index, value := range s {
		fmt.Println(index, value)
	}

	fmt.Println("Agregando un elemento al slice")

	s = append(s, 4)

	for index, value := range s {
		fmt.Println(index, value)
	}

	c := make(chan int)

	go doSomething(c)

	<-c

	g := 25

	h := &g

	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(*h)

}

func doSomething(c chan int) {
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
	c <- 1
}
