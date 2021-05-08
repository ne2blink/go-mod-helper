package main

import (
	"fmt"
	"os"

	helper "github.com/ne2blink/go-mod-helper"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no link")
		os.Exit(1)
	}
	h := helper.New(os.Args[1])
	err := h.Get()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	h.Print()
}
