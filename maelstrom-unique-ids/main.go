package main

import (
	// "encoding/json"
	"encoding/json"
	// "fmt"
	"log"
	"math/rand"
	"slices"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

var backupUniqueId = -1
var generatedIds []int
var currNodeId any
var node_ids any

func generateUniqueId() int {
	backupUniqueId = rand.Intn(1000000)
	for slices.Contains(generatedIds, backupUniqueId) {
		backupUniqueId = rand.Intn(1000000)
	}
	generatedIds = append(generatedIds, backupUniqueId)
	return backupUniqueId
}


func main() {

	//Algorithm 1:
	//Sender calculates "node_ids" and "node_id" by calling n.NodeIds() and n.ID()
	//and set otherNodes slice = "node_ids"-"node_id" 
	//Sender generates a unique id
	//send a message body of type "share", msg_id of node_id number + uniqueIdShared, new_id of the newly generated unique id
	//Receiver has a message handler that listen for messages of type "share"
	//and deconstructs the message for the property "new_id"
	//add the value mapped to "new_id" to its generatedIds slice

	n := maelstrom.NewNode()


	n.Handle("init", func(msg maelstrom.Message) error {
		var requestBody map[string]any
	
		if err := json.Unmarshal(msg.Body, &requestBody); err != nil {
			return err
		}	

		currNodeId = requestBody["node_id"]

		// a, _ := json.Marshal(map[string]any{"currNodeId": currNodeId})
		// fmt.Println(string(a)) // {"bar":2,"baz":3,"foo":1}


		requestBody["type"] = "init_ok"
		requestBody["in_reply_to"] = requestBody["msg_id"]

		return n.Reply(msg, requestBody)
	})


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