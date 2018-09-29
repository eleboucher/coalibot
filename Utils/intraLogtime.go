package Utils

import (
	"time"

	"gitlab.com/clafoutis/api42"
)

func IntraLogtime(user string, rangeBegin time.Time, rangeEnd time.Time, client *api42.Client42) time.Duration {
	beginAt := rangeBegin.Format("2006-01-02")
	endAt := rangeEnd.Format("2006-01-02")
	params := api42.NewParameter()
	params.AddRange("begin_at", beginAt, endAt)
	params.PerPage = 100
	var locations []api42.Location42
	if beginAt == endAt {
		return 0
	}
	locations, err := client.GetUserLocations(user, params)
	if err != nil || len(locations) == 0 {
		return 0
	}
	var page = 2
	for {
		lastlocation := locations[len(locations)-1].EndAt
		if rangeBegin.Before(*lastlocation) {
			params.Page = page
			data, err := client.GetUserLocations(user, params)
			if err != nil {
				return 0
			}
			locations = append(locations, data...)
			if len(data) == 0 {
				break
			}
			page++
		} else {
			break
		}
	}
	var logtime time.Duration
	for i := 0; i < len(locations); i++ {
		var logEnd time.Time
		if locations[i].EndAt == nil {
			logEnd = time.Now()
		} else {
			logEnd = *locations[i].EndAt
		}
		logtime += logEnd.Sub(*locations[i].BeginAt)
	}
	return logtime
}
