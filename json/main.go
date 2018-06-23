package main

import (
	"encoding/json"
	"fmt"
)

// Reported message
type Reported struct {
	InterfaceAddrs interface{}
}

// State state message
type State struct {
	reported Reported
}

// Message AWS IoT message
type Message struct {
	state State
}

func main() {
	message := Message{
		state: State{
			reported: Reported{
				InterfaceAddrs: "test",
			},
		},
	}

	payload, err := json.Marshal(message)

	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("payload:", payload)
	}
}
