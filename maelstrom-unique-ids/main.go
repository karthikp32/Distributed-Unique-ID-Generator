package main

import (
	// "encoding/json"
	"encoding/json"
	"log"
	"math/rand"
	"slices"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var backupUniqueId = -1
var generatedIds []int

func generateUniqueId() int {
	backupUniqueId = rand.Intn(1000000)
	for !slices.Contains(generatedIds, backupUniqueId) {

	}
	generatedIds = append(generatedIds, backupUniqueId)
	return backupUniqueId
}


func main() {
	// for i:=1; i <= 10; i++ {
	// 	fmt.Println(generateUniqueId())
	// }

	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {
		var requestBody map[string]any
	
		if err := json.Unmarshal(msg.Body, &requestBody); err != nil {
			return err
		}	

		requestBody["type"] = "generate_ok"
		requestBody["id"] = generateUniqueId()

		return n.Reply(msg, requestBody)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}