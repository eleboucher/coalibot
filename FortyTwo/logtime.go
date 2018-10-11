package FortyTwo

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
	"github.com/nlopes/slack"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var now = time.Now()
var usage = "```" + `Usage: 
bc logtime [-i] [LOGIN]
bc logtime [-i] -d BEGIN END [LOGIN]
bc logtime [-i] -y YEAR [LOGIN]
bc logtime [-i] -m MONTH [YEAR] [LOGIN]
bc logtime [-i] -t TRIMESTER [YEAR] [LOGIN]
bc logtime [-i] -s SEMESTER [YEAR] [LOGIN]

logtime displays the total school presence logtime

-i                   displays cluster logtime instead of school presence

-d BEGIN END displays logtime between begin_date and end_date
									 date format: %dd/%MM/%yyyy
-y YEAR             displays logtime during the given year
-m MONTH [YEAR]      displays logtime during the given month
-t TRIMESTER [YEAR] displays logtime during the given trimester
-s SEMESTER [YEAR]  displays logtime during the given semester

NOTE: school presence of the month is only updated at the end of the month.
School presence is only stored since Nov 2017.` + "```"

type logopt struct {
	count     int
	intra     bool
	dateBegin time.Time
	dateEnd   time.Time
	login     string
	logtime   time.Duration
	error     bool
}

var monthL = map[string]int{
	"janvier":   1,
	"fevrier":   2,
	"mars":      3,
	"avril":     4,
	"mai":       5,
	"juin":      6,
	"juillet":   7,
	"aout":      8,
	"septembre": 9,
	"octobre":   10,
	"novembre":  11,
	"decembre":  12,
}

var zeroDate = time.Time{}

func Logtime(option string, event *Struct.Message) bool {
	var logtimeOpt = logopt{count: 0, intra: false, dateBegin: zeroDate, dateEnd: zeroDate, login: "", logtime: -1, error: false}
	var splited = strings.Split(option, " ")
	if splited[logtimeOpt.count] == "-i" || splited[logtimeOpt.count] == "--intra" {
		logtimeOpt.intra = true
		logtimeOpt.count++
	}

	if len(splited) > logtimeOpt.count {
		switch splited[logtimeOpt.count] {
		case "--date", "-d":
			handleDate(splited, &logtimeOpt)
		case "--year", "-y":
			handleYear(splited, &logtimeOpt)
		case "--month", "-m":
			handleMonth(splited, &logtimeOpt)
		case "--trimestre", "-t":
			handleQuarter(splited, &logtimeOpt)
		case "--semestre", "-s":
			handleQuarter(splited, &logtimeOpt)
		case "-h", "--help":
			logtimeOpt.error = true
		default:
			logtimeOpt.dateBegin = time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, now.Location())
			logtimeOpt.dateEnd = logtimeOpt.dateBegin.AddDate(1, 0, 0).Add(-time.Nanosecond)
		}
	} else {
		logtimeOpt.dateBegin = time.Date(time.Now().Year(), time.January, 1, 0, 0, 0, 0, now.Location())
		logtimeOpt.dateEnd = logtimeOpt.dateBegin.AddDate(1, 0, 0).Add(-time.Nanosecond)
	}

	if len(splited) > logtimeOpt.count && (splited[logtimeOpt.count] == "-i" || splited[logtimeOpt.count] == "--intra") {
		logtimeOpt.intra = true
		logtimeOpt.count++
	}
	if !logtimeOpt.error {
		if len(splited) > logtimeOpt.count {
			logtimeOpt.login, logtimeOpt.error = Utils.GetLogin(splited[logtimeOpt.count], event)
		} else {
			logtimeOpt.login, logtimeOpt.error = Utils.GetLogin("", event)
		}
	}
	if len(splited) > logtimeOpt.count && !logtimeOpt.error {
		logtimeOpt.count++
	}
	if len(splited) > logtimeOpt.count {
		logtimeOpt.error = true
	}
	if logtimeOpt.error {
		params := Struct.SlackParams
		params.ThreadTimestamp = event.Timestamp
		event.API.PostMessage(event.Channel, usage, params)
		return false
	}
	if logtimeOpt.dateBegin != zeroDate && logtimeOpt.dateEnd != zeroDate {
		switch logtimeOpt.intra {
		case false:
			logtimeOpt.logtime = Utils.Logtime(logtimeOpt.login, logtimeOpt.dateBegin, logtimeOpt.dateEnd, event.FortyTwo)
		case true:
			logtimeOpt.logtime = Utils.IntraLogtime(logtimeOpt.login, logtimeOpt.dateBegin, logtimeOpt.dateEnd, event.FortyTwo)
		}
	}

	if logtimeOpt.logtime != -1 {
		var logtimeStr = Utils.FmtDuration(logtimeOpt.logtime)
		params := Struct.SlackParams
		params.ThreadTimestamp = event.Timestamp
		attachment := slack.Attachment{
			Color: "good",
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Title: "Résultat",
					Value: logtimeStr,
					Short: true,
				},
			},
			Footer: "Powered by Coalibot",
		}
		params.Attachments = []slack.Attachment{attachment}
		var intra = "intra"
		if !logtimeOpt.intra {
			intra = "badgeuse"
		}
		event.API.PostMessage(event.Channel, "Logtime *"+intra+"* pour *"+logtimeOpt.login+"* entre *"+logtimeOpt.dateBegin.Format("2006-01-02")+"* et *"+logtimeOpt.dateEnd.Format("2006-01-02")+"*", params)
	}
	return true
}

func handleDate(splited []string, logtimeOpt *logopt) {
	(*logtimeOpt).count++

	if len(splited) < (*logtimeOpt).count {
		(*logtimeOpt).error = true
		return
	}
	dateBegin, err := time.Parse("02/01/2006", splited[logtimeOpt.count])

	if err != nil {
		(*logtimeOpt).error = true
		return
	}
	(*logtimeOpt).dateBegin = dateBegin
	dateEnd, err := time.Parse("02/01/2006", splited[logtimeOpt.count+1])
	if err != nil {
		(*logtimeOpt).error = true
		return
	}
	y, m, d := dateEnd.Date()
	(*logtimeOpt).dateEnd = time.Date(y, m, d, 23, 59, 59, int(-time.Nanosecond), dateEnd.Location())
	(*logtimeOpt).count += 2
}

func handleYear(splited []string, logtimeOpt *logopt) {
	(*logtimeOpt).count++
	yearReg, _ := regexp.Compile(`(\b|^)20\d{2}(\b|$)`)
	if len(splited) < (*logtimeOpt).count || !yearReg.MatchString(splited[logtimeOpt.count]) {
		(*logtimeOpt).error = true
		return
	}
	year, _ := strconv.Atoi(splited[logtimeOpt.count])
	(*logtimeOpt).dateBegin = time.Date(year, time.January, 1, 0, 0, 0, 0, now.Location())
	(*logtimeOpt).dateEnd = (*logtimeOpt).dateBegin.AddDate(1, 0, 0).Add(-time.Nanosecond)
	(*logtimeOpt).count++

}

func handleQuarter(splited []string, logtimeOpt *logopt) {
	(*logtimeOpt).count++
	quartReg, _ := regexp.Compile(`^[1-4]$`)
	if len(splited) < logtimeOpt.count || !quartReg.MatchString(splited[logtimeOpt.count]) {
		(*logtimeOpt).error = true
		return
	}
	quarter, _ := strconv.Atoi(splited[logtimeOpt.count])
	year := time.Now().Year()
	yearReg, _ := regexp.Compile(`(\b|^)20\d{2}(\b|$)`)
	if len(splited) > (*logtimeOpt).count+1 && yearReg.MatchString(splited[logtimeOpt.count+1]) {
		year, _ = strconv.Atoi(splited[logtimeOpt.count+1])
		(*logtimeOpt).count++
	}
	(*logtimeOpt).dateBegin = time.Date(year, time.Month((quarter-1)*3+1), 1, 0, 0, 0, 0, now.Location())
	(*logtimeOpt).dateEnd = (*logtimeOpt).dateBegin.AddDate(0, 3, 0).Add(-time.Nanosecond)
	fmt.Println((*logtimeOpt).dateEnd)
	(*logtimeOpt).count++
}

func handleSemester(splited []string, logtimeOpt *logopt) {
	(*logtimeOpt).count++
	quartReg, _ := regexp.Compile(`(\b|^)[1-2](\b|$)`)
	if len(splited) < (*logtimeOpt).count || !quartReg.MatchString(splited[logtimeOpt.count]) {
		(*logtimeOpt).error = true
		return
	}
	semestre, _ := strconv.Atoi(splited[logtimeOpt.count])
	year := time.Now().Year()
	yearReg, _ := regexp.Compile(`(\b|^)20\d{2}(\b|$)`)
	if len(splited) >= (*logtimeOpt).count+1 && yearReg.MatchString(splited[logtimeOpt.count+1]) {
		year, _ = strconv.Atoi(splited[logtimeOpt.count+1])
		(*logtimeOpt).count++
	}
	(*logtimeOpt).dateBegin = time.Date(year, time.Month((semestre-1)*6+1), 1, 0, 0, 0, 0, now.Location())
	(*logtimeOpt).dateEnd = (*logtimeOpt).dateBegin.AddDate(0, 6, 0).Add(-time.Nanosecond)
	(*logtimeOpt).count++
}

func handleMonth(splited []string, logtimeOpt *logopt) {
	var hasYear = false
	(*logtimeOpt).count++
	if len(splited) < (*logtimeOpt).count {
		(*logtimeOpt).error = true
		return
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	month, _, _ := transform.String(t, strings.ToLower(splited[logtimeOpt.count]))
	monthReg, _ := regexp.Compile(`(\b|^)(0[1-9]|[1-9]|1[012])(\b|$)`)
	year := time.Now().Year()
	yearReg, _ := regexp.Compile(`(\b|^)20\d{2}(\b|$)`)
	if len(splited) > (*logtimeOpt).count+1 && yearReg.MatchString(splited[logtimeOpt.count+1]) {
		year, _ = strconv.Atoi(splited[logtimeOpt.count+1])
		hasYear = true
	}
	if value, ok := monthL[month]; ok {
		(*logtimeOpt).dateBegin = time.Date(year, time.Month(value), 1, 0, 0, 0, 0, now.Location())
		(*logtimeOpt).dateEnd = (*logtimeOpt).dateBegin.AddDate(0, 1, 0).Add(-time.Nanosecond)
	} else if monthReg.MatchString(month) {
		monthInt, _ := strconv.Atoi(splited[logtimeOpt.count])
		(*logtimeOpt).dateBegin = time.Date(year, time.Month(monthInt), 1, 0, 0, 0, 0, now.Location())
		(*logtimeOpt).dateEnd = (*logtimeOpt).dateBegin.AddDate(0, 1, 0).Add(-time.Nanosecond)
	} else {
		(*logtimeOpt).error = true
	}
	if hasYear {
		(*logtimeOpt).count += 2
	} else {
		(*logtimeOpt).count++
	}

}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
