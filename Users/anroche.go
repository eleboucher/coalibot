package Users

import (
	"math/rand"
	"time"

	"github.com/genesixx/coalibot/Struct"
)

const (
	Decisecond = 100 * time.Millisecond
	Day        = 24 * time.Hour
)

func Anroche(option string, event *Struct.Message) bool {
	letItGo := []string{"https://www.youtube.com/watch?v=L0MK7qz13bU", "https://www.youtube.com/watch?v=wQP9XZc2Y_c", "https://www.youtube.com/watch?v=bgq5nlCYzaE", "https://www.youtube.com/watch?v=BbOLik9Esqo", "https://www.youtube.com/watch?v=BjwDV1Is34U", "https://www.youtube.com/watch?v=DpJYhF1M_-o", "https://www.youtube.com/watch?v=vaZYGX6BimI", "https://www.youtube.com/watch?v=70M-mSxB2bU", "https://www.youtube.com/watch?v=W-66bxpzkQw", "https://www.youtube.com/watch?v=R-sJk6iIEAA"}
	event.API.PostMessage(event.Channel, letItGo[rand.Int()%len(letItGo)-1], Struct.SlackParams)
	return true
}
