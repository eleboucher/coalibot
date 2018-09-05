package main

import (
	"fmt"
	"strings"
)

func indexOf(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}

func reply()

// m := map[string] fn {
//     "f": foo,
//     "b": bar,
//   }

func handleCommand(message *Message) {
	var isCommand = false
	var option = ""
	var command = ""
	// remove extra whitespace
	message.Message = strings.Join(strings.Fields(message.Message), " ")

	fmt.Printf("<#%s> @%s: %s\n", message.Channel, message.User, message.Message)
	splited := strings.Split(message.Message, " ")
	if indexOf(splited[0], []string{"coalibot", "bc", "cb"}) > -1 && len(splited) > 1 {
		fmt.Printf("Coalibot command\n")
	} else if splited[0][0] == '!' && len(splited[0]) > 1 {
		fmt.Printf("Coalibot command\n")
	}

}
