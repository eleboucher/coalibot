package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/nlopes/slack"
)

func indexOf(word string, data []string) int {
	for k, v := range data {
		if word == v {
			return k
		}
	}
	return -1
}

func reply(command string, event *Message) bool {
	// Open our jsonFile
	jsonFile, err := os.Open("reply.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// a map container to decode the JSON structure into
	c := make(map[string]interface{})

	// unmarschal JSON
	e := json.Unmarshal(byteValue, &c)
	if e != nil || c[command] == nil {
		return false
	}

	// output result to STDOUT
	fmt.Printf("reply %s\n", c[command].(string))
	event.API.PostMessage(event.Channel, c[command].(string), slack.PostMessageParameters{})
	return true
}

func hello(option string, event *Message) bool {
	event.API.PostMessage(event.Channel, "Hello <@"+event.User+"> ! powered by go", slack.PostMessageParameters{})
	return true
}

var commands = map[string]func(string, *Message) bool{
	"hello": hello,
}

func handleCommand(event *Message) {
	var isCommand = false
	var option = ""
	var command = ""

	event.Message = strings.Join(strings.Fields(event.Message), " ")
	fmt.Printf("<#%s> @%s: %s\n", event.Channel, event.User, event.Message)
	splited := strings.Split(event.Message, " ")
	if indexOf(splited[0], []string{"coalibot", "bc", "cb"}) > -1 && len(splited) > 1 {
		command = splited[1]
		option = strings.Join(splited[1:], " ")
		isCommand = reply(command, event)
		if !isCommand {
			if isCommand, err := commands[command](option, event); ok {

			}
		}
	} else if splited[0][0] == '!' && len(splited[0]) > 1 {

		fmt.Printf("Coalibot command\n")
	}

}
