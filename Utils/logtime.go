package Utils

import (
	"bufio"
	"encoding/csv"
	"os"
	"regexp"
	"strconv"
	"time"
	"unicode"

	"gitlab.com/clafoutis/api42"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type name struct {
	firstName string
	lastName  string
}

func Logtime(user string, rangeBegin time.Time, rangeEnd time.Time, client *api42.Client42) time.Duration {
	data, err := client.GetUser(user)
	if err != nil || data.FirstName == "" || data.LastName == "" {
		return 0
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	firstName, _, _ := transform.String(t, data.FirstName)
	lastName, _, _ := transform.String(t, data.LastName)
	var name = name{firstName: firstName, lastName: lastName}
	var current = rangeBegin
	var duration int
	for !current.After(rangeEnd) {
		csvFile, _ := os.Open("logtime/presence-" + current.Format("2006") + "-" + current.Format("01") + ".csv")
		reader := csv.NewReader(bufio.NewReader(csvFile))
		reader.Comma = ';'
		reader.TrimLeadingSpace = true
		reader.TrailingComma = true
		rows, err := reader.ReadAll()

		csvFile.Close()
		if err == nil {
			duration += getHourByName(name, rows)
		}
		current = current.AddDate(0, 1, 0)
	}
	ret, _ := time.ParseDuration(strconv.Itoa(duration) + "h")
	return ret
}

func getHourByName(name name, data [][]string) int {
	nameReg, _ := regexp.Compile("(?i)(\b|^)" + name.lastName + "(\\d|), " + name.firstName + "(\\d|)(\b|$)")
	var duration = 0
	for i := 0; i < len(data); i++ {
		if nameReg.MatchString(data[i][0]) {
			tmp, _ := strconv.Atoi(data[i][2][:len(data[i][2])-1])
			duration += tmp
		}
	}
	return duration
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}
