package FortyTwo

import (
	"fmt"
	"strings"
	"time"

	"github.com/genesixx/coalibot/Utils"
	"gitlab.com/clafoutis/api42"
)

func getNumber(user *api42.User42) string {
	switch user.Phone {
	case "":
		return "Hidden"
	default:
	}
	return user.Phone
}

func getMainCampus(user *api42.User42) string {
	selected := -1
	for _, campus := range user.CampusUsers {
		if campus.IsPrimary {
			selected = campus.CampusID
		}
	}
	for _, campus := range user.Campus {
		if campus.ID == selected {
			return campus.Name
		}
	}
	return "Unknown Campus"
}

func getTitle(user *api42.User42) string {
	selected := -1
	for _, title := range user.TitleUsers {
		if title.Selected {
			selected = title.TitleID
		}
	}
	for _, title := range user.Titles {
		if title.ID == selected {
			return strings.Replace(title.Name, "%login", user.Login, -1)
		}
	}
	return user.Login
}

func getLogtime(user string, client *api42.Client42) time.Duration {
	y, m, d := time.Now().Date()
	rangeBegin := time.Date(y, m, d, 0, 0, 0, 0, time.Now().Location())
	rangeEnd := rangeBegin.AddDate(0, 0, -7)
	return Utils.IntraLogtime(user, rangeEnd, rangeBegin, client)
}

func getCoasRepr(user string, client *api42.Client42, blocs []api42.Bloc42, coalitions []api42.Coalition42) (string, string) {
	coaData := getCoalition(1, blocs, coalitions)
	color := "#D40000"
	slug := ""
	if coaData == nil && len(coalitions) >= 1 {
		coaData = &coalitions[0]
	}
	if coaData != nil {
		return coaData.Color, fmt.Sprintf(":%s:", coaData.Slug)
	}
	return color, slug
}

func getCoalition(cursusID int, blocs []api42.Bloc42, coas []api42.Coalition42) *api42.Coalition42 {
	for _, bloc := range blocs {
		if bloc.CursusID != cursusID {
			continue
		}
		for _, bCoa := range bloc.Coalitions {
			for _, coa := range coas {
				if coa.ID == bCoa.ID {
					return &coa
				}
			}
		}
	}
	return nil
}

func getNameCoa(cursusID int, blocs []api42.Bloc42, coas []api42.Coalition42) string {
	if coa := getCoalition(cursusID, blocs, coas); coa != nil {
		return coa.Slug
	}
	return ""
}

func cursusLevels(cursus []api42.CursusUser42, blocs []api42.Bloc42, coas []api42.Coalition42, client *api42.Client42) string {
	var builder strings.Builder
	for _, c := range cursus {
		if c.HasCoalition {
			name := getNameCoa(c.CursusID, blocs, coas)
			switch len(name) {
			case 0:
			default:
				builder.WriteString(fmt.Sprintf(":%s: • ", name))
			}
		}
		builder.WriteString(fmt.Sprintf("%s — _%02.2f_\n", c.Cursus.Name, c.Level))
	}
	return builder.String()
}

func hasDoneIntership(user *api42.User42) string {
	var stage = ":negative_squared_cross_mark:"
	var indexInternProject = -1
	var indexContractProject = -1
	for k, v := range user.Projects {
		if v.Project.ID == 118 {
			indexInternProject = k
		} else if v.Project.ID == 119 {
			indexContractProject = k
		}
	}
	if indexInternProject != -1 && indexContractProject != -1 &&
		user.Projects[indexContractProject].Status == "finished" &&
		(user.Projects[indexInternProject].FinalMark == nil || *user.Projects[indexInternProject].FinalMark > 0) {
		switch user.Projects[indexInternProject].Status {
		case "finished":
			stage = ":white_check_mark:"
		case "in_progress":
			stage = ":clock1:"
		}
	}
	return stage
}
