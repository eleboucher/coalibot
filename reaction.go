package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func InitReaction() []Struct.React {
	file, err := os.OpenFile("reaction.json", os.O_WRONLY|os.O_CREATE, 0660)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadFile("reaction.json")
	var reactions []Struct.React
	json.Unmarshal(byteValue, &reactions)
	for i := 0; i < len(reactions); i++ {
		reactions[i].Compiled, _ = regexp.Compile(fmt.Sprintf("(?i)(^|[^a-zA-Z0-9])%s($|[^a-zA-Z0-9])", reactions[i].Match))
	}
	return reactions
}

func reacts(event Struct.Message, reaction Struct.React, msgRef slack.ItemRef) {
	for i := 0; i < len(reaction.Reactions); i++ {
		event.API.AddReaction(reaction.Reactions[i], msgRef)
	}
}

func React(event Struct.Message, reactions []Struct.React) {
	msgRef := slack.NewRefToMessage(event.Channel, event.Timestamp)
	for i := 0; i < len(reactions); i++ {
		if reactions[i].Compiled.FindStringIndex(event.Message) != nil {
			if reactions[i].Reaction != "" {
				go event.API.AddReaction(reactions[i].Reaction, msgRef)
			} else {
				go reacts(event, reactions[i], msgRef)
			}
		}
	}
}
