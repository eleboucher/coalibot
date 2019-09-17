package twitch

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/genesixx/coalibot/utils"
	"github.com/nlopes/slack"
)

type gbStruct struct {
	ID   int    `json:"id"`
	Code string `json:"code"`
}

type sbStruct struct {
	Emotes []gbStruct `json:"emotes"`
}

func download(url string) io.ReadCloser {
	response, e := http.Get(url)
	if e != nil {
		return nil
	}

	return response.Body
}

func getGbEmotes(emote string) *gbStruct {
	resp, err := http.Get("https://twitchemotes.com/api_cache/v3/global.json")
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	var result map[string]*gbStruct
	json.Unmarshal(data, &result)
	if result[emote] != nil {
		return result[emote]
	}
	return nil
}

func getSbEmotes(emoteName string) *gbStruct {
	file, _ := ioutil.ReadFile("subscriber.json")

	var result map[string]sbStruct
	json.Unmarshal([]byte(file), &result)
	for _, data := range result {
		for _, emote := range data.Emotes {
			if emote.Code == emoteName {
				return &emote
			}
		}
	}
	return nil
}

func Emotes(option string, event *utils.Message) bool {
	emote := getGbEmotes(option)
	if emote == nil {
		emote = getSbEmotes(option)
	}
	if emote == nil {
		return false
	}
	url := "https://static-cdn.jtvnw.net/emoticons/v1/" + strconv.Itoa(emote.ID) + "/2.0"
	downloaded := download(url)
	defer downloaded.Close()
	if downloaded == nil {
		return false
	}
	params := slack.FileUploadParameters{
		Reader:   downloaded,
		Channels: []string{event.Channel},
		Filename: option + ".png",
	}
	event.API.UploadFile(params)
	return true
}
