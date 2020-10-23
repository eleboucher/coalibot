package fortyTwo

import (
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/eleboucher/coalibot/utils"
	"github.com/slack-go/slack"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var now = time.Now()
var usage = "```" + `Usage:
bc logtime [-b] [LOGIN]
bc logtime [-b] -d BEGIN END [LOGIN]
bc logtime [-b] -y YEAR [LOGIN]
bc logtime [-b] -m MONTH [YEAR] [LOGIN]
bc logtime [-b] -t TRIMESTER [YEAR] [LOGIN]
bc logtime [-b] -s SEMESTER [YEAR] [LOGIN]
bc logtime [-b] -w [LOGIN]

logtime displays the total cluster logtime
default parameter is the current month

-b  --badgeuse            displays school presence instead of cluster logtime

-d BEGIN END displays logtime between begin_date and end_date
									 date format: %dd/%MM/%yyyy
-y YEAR             displays logtime during the given year
-m MONTH [YEAR]     displays logtime during the given month
-t TRIMESTER [YEAR] displays logtime during the given trimester
-s SEMESTER [YEAR]  displays logtime during the given semester
-w --week			displays logtime during the last 7 days

NOTE: school presence of the month is only updated at the end of the month.
School presence is only stored since Nov 2017.` + "```"

type logopt struct {
	count     int
	badgeuse  bool
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
	"january":   1,
	"february":  2,
	"march":     3,
	"april":     4,
	"may":       5,
	"june":      6,
	"july":      7,
	"august":    8,
	"september": 9,
	"october":   10,
	"november":  11,
	"december":  12,
}

var zeroDate = time.Time{}

func Logtime(option string, event *utils.Message) bool {
	var logtimeOpt = logopt{count: 0, badgeuse: false, dateBegin: zeroDate, dateEnd: zeroDate, login: "", logtime: -1, error: false}
	var splited = strings.Split(option, " ")
	if splited[logtimeOpt.count] == "-b" || splited[logtimeOpt.count] == "--badgeuse" {
		logtimeOpt.badgeuse = true
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
		case "--week", "-w":
			handleWeek(splited, &logtimeOpt)
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

	if len(splited) > logtimeOpt.count && (splited[logtimeOpt.count] == "-b" || splited[logtimeOpt.count] == "--badgeuse") {
		logtimeOpt.badgeuse = true
		logtimeOpt.count++
	}
	if !logtimeOpt.error {
		if len(splited) > logtimeOpt.count {
			logtimeOpt.login, logtimeOpt.error = utils.GetLogin(splited[logtimeOpt.count], event)
		} else {
			logtimeOpt.login, logtimeOpt.error = utils.GetLogin("", event)
		}
	}
	if len(splited) > logtimeOpt.count && !logtimeOpt.error {
		logtimeOpt.count++
	}
	if len(splited) > logtimeOpt.count {
		logtimeOpt.error = true
	}
	if logtimeOpt.error {
		params := utils.SlackParams
		params.ThreadTimestamp = event.Timestamp
		utils.PostMsg(event, slack.MsgOptionText(usage, false), slack.MsgOptionTS(event.Timestamp))
		return false
	}
	if logtimeOpt.dateBegin != zeroDate && logtimeOpt.dateEnd != zeroDate {
		switch logtimeOpt.badgeuse {
		case true:
			logtimeOpt.logtime = utils.Logtime(logtimeOpt.login, logtimeOpt.dateBegin, logtimeOpt.dateEnd, event.FortyTwo)
		case false:
			logtimeOpt.logtime = utils.IntraLogtime(logtimeOpt.login, logtimeOpt.dateBegin, logtimeOpt.dateEnd, event.FortyTwo)
		}
	}

	if logtimeOpt.logtime != -1 {
		var logtimeStr = utils.FmtDuration(logtimeOpt.logtime)
		attachment := slack.Attachment{
			Color: "good",
			Fields: []slack.AttachmentField{
				{
					Title: "Logtime",
					Value: logtimeStr,
					Short: true,
				},
			},
			Footer: "Powered by Coalibot",
		}
		var intra = "badgeuse"
		if !logtimeOpt.badgeuse {
			intra = "intra"
		}
		utils.PostMsg(event, slack.MsgOptionText("Logtime *"+intra+"* for *"+logtimeOpt.login+"* between *"+logtimeOpt.dateBegin.Format("2006-01-02")+"* and *"+logtimeOpt.dateEnd.Format("2006-01-02")+"*", false), slack.MsgOptionAttachments(attachment), slack.MsgOptionTS(event.Timestamp))
	}
	return true
}

func handleDate(splited []string, logtimeOpt *logopt) {
	logtimeOpt.count++

	if len(splited) < logtimeOpt.count+2 {
		logtimeOpt.error = true
		return
	}
	dateBegin, err := time.Parse("02/01/2006", splited[logtimeOpt.count])

	if err != nil {
		logtimeOpt.error = true
		return
	}
	logtimeOpt.dateBegin = dateBegin
	dateEnd, err := time.Parse("02/01/2006", splited[logtimeOpt.count+1])
	if err != nil {
		logtimeOpt.error = true
		return
	}
	y, m, d := dateEnd.Date()
	logtimeOpt.dateEnd = time.Date(y, m, d, 23, 59, 59, int(-time.Nanosecond), dateEnd.Location())
	logtimeOpt.count += 2
}

func handleYear(splited []string, logtimeOpt *logopt) {
	logtimeOpt.count++
	yearReg, _ := regexp.Compile(`(\b|^)20\d{2}(\b|$)`)
	if len(splited) <= logtimeOpt.count || !yearReg.MatchString(splited[logtimeOpt.count]) {
		logtimeOpt.error = true
		return
	}
	year, _ := strconv.Atoi(splited[logtimeOpt.count])
	logtimeOpt.dateBegin = time.Date(year, time.January, 1, 0, 0, 0, 0, now.Location())
	logtimeOpt.dateEnd = logtimeOpt.dateBegin.AddDate(1, 0, 0).Add(-time.Nanosecond)
	logtimeOpt.count++
}

func handleWeek(splited []string, logtimeOpt *logopt) {
	y, m, d := time.Now().Date()
	logtimeOpt.dateEnd = time.Date(y, m, d, 0, 0, 0, 0, now.Location())
	logtimeOpt.dateBegin = logtimeOpt.dateEnd.AddDate(0, 0, -7).Add(-time.Nanosecond)
	logtimeOpt.count++
}

func handleQuarter(splited []string, logtimeOpt *logopt) {
	logtimeOpt.count++
	quartReg, _ := regexp.Compile(`(?i)^[1-4]|automne|ete|été|printemps|hiver|spring|fall|winter|summer$`)
	if len(splited) <= logtimeOpt.count || !quartReg.MatchString(splited[logtimeOpt.count]) {
		logtimeOpt.error = true
		return
	}
	var quarter int
	switch strings.ToLower(splited[logtimeOpt.count]) {
	case "hiver", "winter":
		quarter = 1
	case "printemps", "spring":
		quarter = 2
	case "ete", "été", "summer":
		quarter = 3
	case "fall", "automne":
		quarter = 4
	default:
		quarter, _ = strconv.Atoi(splited[logtimeOpt.count])
	}
	year := time.Now().Year()
	yearReg, _ := regexp.Compile(`(\b|^)20\d{2}(\b|$)`)
	if len(splited) > logtimeOpt.count+1 && yearReg.MatchString(splited[logtimeOpt.count+1]) {
		year, _ = strconv.Atoi(splited[logtimeOpt.count+1])
		logtimeOpt.count++
	}
	logtimeOpt.dateBegin = time.Date(year, time.Month((quarter-1)*3+1), 1, 0, 0, 0, 0, now.Location())
	logtimeOpt.dateEnd = logtimeOpt.dateBegin.AddDate(0, 3, 0).Add(-time.Nanosecond)
	logtimeOpt.count++
}

func handleSemester(splited []string, logtimeOpt *logopt) {
	logtimeOpt.count++
	quartReg, _ := regexp.Compile(`(\b|^)[1-2](\b|$)`)
	if len(splited) <= logtimeOpt.count || !quartReg.MatchString(splited[logtimeOpt.count]) {
		logtimeOpt.error = true
		return
	}
	semestre, _ := strconv.Atoi(splited[logtimeOpt.count])
	year := time.Now().Year()
	yearReg, _ := regexp.Compile(`(\b|^)20\d{2}(\b|$)`)
	if len(splited) >= logtimeOpt.count+1 && yearReg.MatchString(splited[logtimeOpt.count+1]) {
		year, _ = strconv.Atoi(splited[logtimeOpt.count+1])
		logtimeOpt.count++
	}
	logtimeOpt.dateBegin = time.Date(year, time.Month((semestre-1)*6+1), 1, 0, 0, 0, 0, now.Location())
	logtimeOpt.dateEnd = logtimeOpt.dateBegin.AddDate(0, 6, 0).Add(-time.Nanosecond)
	logtimeOpt.count++
}

func handleMonth(splited []string, logtimeOpt *logopt) {
	var hasYear = false
	logtimeOpt.count++
	if len(splited) <= logtimeOpt.count {
		logtimeOpt.error = true
		return
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	month, _, _ := transform.String(t, strings.ToLower(splited[logtimeOpt.count]))
	monthReg, _ := regexp.Compile(`(\b|^)(0[1-9]|[1-9]|1[012])(\b|$)`)
	year := time.Now().Year()
	yearReg, _ := regexp.Compile(`(\b|^)20\d{2}(\b|$)`)
	if len(splited) > logtimeOpt.count+1 && yearReg.MatchString(splited[logtimeOpt.count+1]) {
		year, _ = strconv.Atoi(splited[logtimeOpt.count+1])
		hasYear = true
	}
	if value, ok := monthL[month]; ok {
		logtimeOpt.dateBegin = time.Date(year, time.Month(value), 1, 0, 0, 0, 0, now.Location())
		logtimeOpt.dateEnd = logtimeOpt.dateBegin.AddDate(0, 1, 0).Add(-time.Nanosecond)
	} else if monthReg.MatchString(month) {
		monthInt, _ := strconv.Atoi(splited[logtimeOpt.count])
		logtimeOpt.dateBegin = time.Date(year, time.Month(monthInt), 1, 0, 0, 0, 0, now.Location())
		logtimeOpt.dateEnd = logtimeOpt.dateBegin.AddDate(0, 1, 0).Add(-time.Nanosecond)
	} else {
		logtimeOpt.error = true
	}
	if hasYear {
		logtimeOpt.count += 2
	} else {
		logtimeOpt.count++
	}

}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
