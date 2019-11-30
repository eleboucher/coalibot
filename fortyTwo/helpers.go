package fortyTwo

import (
	"fmt"
	"strings"
	"time"

	"github.com/genesixx/coalibot/utils"
	"gitlab.com/clafoutis/api42"
)

// getMainCampus returns the primary campus' name of the user.
// If the campus can't be found, "Unknown Campus" is returned.
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

// getTitle returns the formated selected title of the User.
// If no title is selected, or if the selected title can't be found, this simply returns the user's login.
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

// getLogtime returns the logtime of the user within the last seven days.
func getLogtime(user string, client *api42.Client42) time.Duration {
	y, m, d := time.Now().Date()
	rangeBegin := time.Date(y, m, d, 0, 0, 0, 0, time.Now().Location())
	rangeEnd := rangeBegin.AddDate(0, 0, -7)
	return utils.IntraLogtime(user, rangeEnd, rangeBegin, client)
}

func getCoalitionEmoji(slug string) string {
	coasDict := map[string]string {
		// 42 Madrid's coalitions
		"atlantis": "coa-atlantis",
		"metropolis": "coa-metropolis",
		"wakanda": "coa-wakanda",
	}
	if emoji, ok := coasDict[slug]; ok {
		return emoji
	}
	return slug
}

// getCoasRepr returns the user's coaltion's slug and color.
func getCoasRepr(user string, client *api42.Client42, blocs []api42.Bloc42, coalitions []api42.Coalition42) (string, string) {
	coaData := getCoalition(1, blocs, coalitions)
	color := "#D40000"
	slug := ""
	if coaData == nil && len(coalitions) >= 1 {
		coaData = &coalitions[0]
	}
	if coaData != nil {
		return coaData.Color, fmt.Sprintf(":%s:", getCoalitionEmoji(coaData.Slug))
	}
	return color, slug
}

// getCoalition returns the coalition corresponding to a cursus
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

// getNameCoa returns the slig of the coalition corresponding to a cursus.
func getNameCoa(cursusID int, blocs []api42.Bloc42, coas []api42.Coalition42) string {
	if coa := getCoalition(cursusID, blocs, coas); coa != nil {
		return coa.Slug
	}
	return ""
}

// cusrusLevels returns the string representation of all the cursuses of the user.
func cursusLevels(cursus []api42.CursusUser42, blocs []api42.Bloc42, coas []api42.Coalition42, client *api42.Client42) string {
	var builder strings.Builder
	for _, c := range cursus {
		if c.HasCoalition {
			name := getNameCoa(c.CursusID, blocs, coas)
			switch len(name) {
			case 0:
			default:
				builder.WriteString(fmt.Sprintf(":%s: • ", getCoalitionEmoji(name)))
			}
		}
		builder.WriteString(fmt.Sprintf("%s — _%02.2f_\n", c.Cursus.Name, c.Level))
	}
	return builder.String()
}

// hasDoneIntership returns an emoticon corresponding to the Internship status.
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
		if v.Project.ID == 1055 {
			indexInternProject = k
		} else if v.Project.ID == 1090 {
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
