package main

import "fmt"

func sum(s []int, chnl chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}

	chnl <- sum
}

func provider(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	s := []int{-9, 10, 11, 7, 9, 22}
	ch := make(chan int)
	anCh := make(chan int)
	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)

	x, y := <-ch, <-ch
	fmt.Printf("x = %d", x, "y = %d", y, "sum = %d]\n", (x + y))

	go provider(anCh)

	for i := range anCh {
		fmt.Println(i)
	}
}
