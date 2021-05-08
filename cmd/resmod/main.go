package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ne2blink/go-mod-helper/pkg/mod"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("no link")
		os.Exit(1)
	}
	resolved, err := mod.Resolve(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resolved)
}
