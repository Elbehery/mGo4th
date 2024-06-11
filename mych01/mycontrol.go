package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("too less args")
		return
	}

	arg := os.Args[1]

	switch arg {
	case "0":
		fmt.Println("Zero!")
	case "1":
		fmt.Println("One!")
	case "2", "3", "4":
		fmt.Println("2 or 3 or 4")
		fallthrough
	default:
		fmt.Println("value ", arg)
	}

	val, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("invalid int as string")
		return
	}

	switch {
	case val == 0:
		fmt.Println("Zero!")
	case val > 0:
		fmt.Println("> Zero")
	case val < 0:
		fmt.Println("< Zero")
	default:
		fmt.Println("impossible !!!")
	}
}
