package main

import (
    // "encoding/json"
    // "log"
    // maelstrom "github.com/jepsen-io/maelstrom/demo/go"
	"math/rand"
	"fmt"

)

func generateUniqueId() int {
	return rand.Intn(1000000)
}


func main() {
	for i:=1; i <= 10; i++ {
		fmt.Println(generateUniqueId())
	}

}