package main

import (
	"fmt"
	"github.com/vazha/wsDataProvider"
	"log"
)

func main() {
	fmt.Println("Main")
	err := wsDataProvider.Run()
	if err != nil {
		log.Println("wsDataProvider fail:", err)
	}
}