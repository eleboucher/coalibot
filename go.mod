module github.com/eleboucher/coalibot

go 1.21.0

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/hako/durafmt v0.0.0-20191009132224-3f39dc1ed9f4
	github.com/joho/godotenv v1.3.0
	github.com/sirupsen/logrus v1.6.0
	github.com/slack-go/slack v0.12.3
	gitlab.com/clafoutis/api42 v0.0.0-00010101000000-000000000000
	golang.org/x/text v0.7.0
)

require (
	github.com/andybalholm/cascadia v1.1.0 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	gitlab.com/clafoutis/api v0.0.0-00010101000000-000000000000 // indirect
	golang.org/x/net v0.7.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
)

replace gitlab.com/clafoutis/api42 => ../../../gitlab.com/clafoutis/api42

replace gitlab.com/clafoutis/api => ../../../gitlab.com/clafoutis/api
