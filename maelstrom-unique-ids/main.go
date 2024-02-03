package main

import (
	// "encoding/json"
	// "log"
	"encoding/json"
	"fmt"
	"math/rand"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var backupUniqueId = -1

func generateUniqueId() int {
	backupUniqueId = rand.Intn(1000000)
	return backupUniqueId
}


func main() {
	// for i:=1; i <= 10; i++ {
	// 	fmt.Println(generateUniqueId())
	// }

	n := maelstrom.NewNode()

	n.Handle("unique-ids", func(msg maelstrom.Message) error {
		var requestBody map[string]any
	
		if err := json.Unmarshal(msg.Body, &requestBody); err != nil {
			return err
		}	

		requestBody["type"] = "generate_ok"
		requestBody["id"] = generateUniqueId()

		return n.Reply(msg, requestBody)
	})

}