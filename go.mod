module github.com/eleboucher/coalibot

go 1.14

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/hako/durafmt v0.0.0-20191009132224-3f39dc1ed9f4
	github.com/joho/godotenv v1.3.0
	github.com/lib/pq v1.5.2 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/slack-go/slack v0.6.4
	gitlab.com/clafoutis/api v0.0.0-00010101000000-000000000000 // indirect
	gitlab.com/clafoutis/api42 v0.0.0-00010101000000-000000000000
	golang.org/x/text v0.3.3
)

replace gitlab.com/clafoutis/api42 => ../../../gitlab.com/clafoutis/api42

replace gitlab.com/clafoutis/api => ../../../gitlab.com/clafoutis/api
