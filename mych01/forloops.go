package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i * i)
	}

	fmt.Println("******************")
	i := 0
	for {
		if i == 10 {
			break
		}
		fmt.Println(i * i)
		i++
	}
}
