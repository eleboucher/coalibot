# API 42 - Client - GO

_By **clafoutis**_

_This client has been made to ease access to [42's API](https://api.intra.42.fr/apidoc) using golang._

## Implemented

This client implements:
* Scopes:
	* Achievements `/v2/achievements` [`GET`](https://api.intra.42.fr/apidoc/2.0/achievements/index.html)
	* Campus `/v2/campus` [`GET`](https://api.intra.42.fr/apidoc/2.0/campus/index.html)
	* Campus Users `/v2/campus_users` [`GET`](https://api.intra.42.fr/apidoc/2.0/campus_users/index.html)
	* Coalitions `/v2/coalitions` [`GET`](https://api.intra.42.fr/apidoc/2.0/coalitions/index.html)
	* Coalition Users `/v2/coalitions_users` [`GET`](https://api.intra.42.fr/apidoc/2.0/coalitons_users/index.html)
	* Cursus `/v2/cursus` [`GET`](https://api.intra.42.fr/apidoc/2.0/cursus/index.html)
	* Cursus Users `/v2/cursus_users` [`GET`](https://api.intra.42.fr/apidoc/2.0/cursus_users/index.html)
	* Events `/v2/events` [`GET`](https://api.intra.42.fr/apidoc/2.0/events/index.html)
	* Expertises `/v2/expertises` [`GET`](https://api.intra.42.fr/apidoc/2.0/expertises/index.html)
	* Feedbacks `/v2/feedbacks` [`GET`](https://api.intra.42.fr/apidoc/2.0/feedbacks/index.html)
	* Languages `/v2/languages` [`GET`](https://api.intra.42.fr/apidoc/2.0/languages/index.html)
	* Language Users `/v2/languages_users` [`GET`](https://api.intra.42.fr/apidoc/2.0/languages_users/index.html)
	* Locations `/v2/locations` [`GET`](https://api.intra.42.fr/apidoc/2.0/locations/index.html)
	* Messages `/v2/messages` [`GET`](https://api.intra.42.fr/apidoc/2.0/messages/index.html)
	* Notions `/v2/notions` [`GET`](https://api.intra.42.fr/apidoc/2.0/notions/index.html)
	* Partnerships `/v2/partnerships` [`GET`](https://api.intra.42.fr/apidoc/2.0/partnerships/index.html)
	* Partnership Users `/v2/partnerships_users` [`GET`](https://api.intra.42.fr/apidoc/2.0/partnerships_users/index.html)
	* Projects `/v2/projects` [`GET`](https://api.intra.42.fr/apidoc/2.0/projects/index.html)
	* Project Session `/v2/project_sessions` [`GET`](https://api.intra.42.fr/apidoc/2.0/project_sessions/index.html)
	* Project Users `/v2/projects_users` [`GET`](https://api.intra.42.fr/apidoc/2.0/projects_users/index.html)
	* Scales `/v2/scales` [`GET`](https://api.intra.42.fr/apidoc/2.0/scales/index.html)
	* Scale Teams `/v2/scale_teams` [`GET`](https://api.intra.42.fr/apidoc/2.0/scale_teams/index.html)
	* Skills `/v2/skills` [`GET`](https://api.intra.42.fr/apidoc/2.0/skills/index.html)
	* Subnotions `/v2/subnotions` [`GET`](https://api.intra.42.fr/apidoc/2.0/subnotions/index.html)
	* Tags `/v2/tags` [`GET`](https://api.intra.42.fr/apidoc/2.0/tags/index.html)
	* Teams `/v2/teams` [`GET`](https://api.intra.42.fr/apidoc/2.0/teams/index.html)
	* Team Uploads `/v2/teams_uploads` ['GET'](https://api.intra.42.fr/apidoc/2.0/teams_uploads/index.html)
	* Titles `/v2/titles` [`GET`](https://api.intra.42.fr/apidoc/2.0/titles/index.html)
	* Title Users `/v2/titles_users` [`GET`](https://api.intra.42.fr/apidoc/2.0/titles_users/index.html)
	* Topics `/v2/topics` [`GET`](https://api.intra.42.fr/apidoc/2.0/topics/index.html)
	* Users `/v2/users` [`GET`](https://api.intra.42.fr/apidoc/2.0/users/index.html)
	* Votes `/v2/votes` [`GET`](https://api.intra.42.fr/apidoc/2.0/votes/index.html)
* Page Filters
* Sort Filters
* Range Filters
* Filters

__All implemented scopes have been fully implemented. The only exceptions are child objects that does not have their scope implemented are not stored YET__

## Documentation

The client has been made using my more generalist [API Client](api).

__Please note that this project is being developed and everything documented here might very well change over time.__

In order to start using this client you have initiate it like this:
```golang
import (
	"gitlab.com/clafoutis/api42"
	"os"
)

// ...

func main(){
// ...
uid := os.Getenv("API_UID") // This is how I think you should store and get your keys, but there are many other ways
secret := os.Getenv("API_SECRET")
client, err := api42.NewAPI(uid, secret)
// ...
}
```

You can edit the limit of request by second. It is set to 2/s by default considering [42's API](https://api.intra.42.fr/apidoc) default app limit.  
There are many ways to achieve this:
```golang
// Usage: client.SetLimit(limit, delay int) -- limit: requests -- delay: seconds
client.SetLimit(3, 1) // To set the limit to 3 requests by seconds
// Usage: client.SetDelay(delay, int) -- delay: seconds
client.SetDelay(4) // To set the delay to 4 seconds
```

You can then send requests manually:
```golang
import (
	"gitlab.com/clafoutis/api42"
	"fmt"
	"os"
)

func main(){
	var users []api42.Users

	uid := os.Getenv("API_UID")
	secret := os.Getenv("API_SECRET")
	client, err := api42.NewAPI(uid, secret)

	// Usage : api42.NewParameter(args... interface{})
	// Explanation below
	params := api42.NewParameter()
	params.PerPage = 100
	// Usage: client.Get(scope string, params *api42.RequestParameter, respHost interface{})
	// scope: path to the scope
	// params: filters and page
	// respHost: pointer to the requested type
	rawResp, err := client.Get("/v2/users", nil, &users)
	// rawResp will contain the raw body of the request
	//...
}
```

As seen earlier can modify request filters, pages, sorts and on by passing a RequestParameter struct to `Get`:

```golang
type RequestFilters map[string]interface{}
type RequestCustoms map[string]interface{}
type RequestRanges map[string][2]int

type RequestParameter struct {
	Filters			RequestFilters // On a get URL looks like ?filter[key]=value
	Ranges			RequestRanges // On a get URL looks like ?range[key]=value1,value2
	Customs			RequestCustoms // Is a custom parameter, on a get URL looks like ?key=value
	Sort			string // On a get URL, looks like ?sort=value
	Page, PerPage	int // page and per_page parameters
}

api42.NewParameter(args... interface{})
// if parameter is of type:
// map[string]interface{} --> 1 Adds to filters - 2 Adds to customs
// map[string][2]int --> Adds to range
// int --> 1 Page - 2 perPage
// string --> sort
```

To simplify the usage of `Get` and other requests, you can use the methods that are made for specific scopes:

```golang
client.GetUsers(nil) // Gets all users on the first pages
client.GetUser("clabouri") // Gets user of login 'clabouri' (this accepts ints and strings)
client.GetTitlesByUser("clabouri", nil) // Gets the first page of titles for user "clabouri"
user.GetTitles(nil) // Same as above
```

To get all the pages of a scope with a single function call, you can pass a parameter with `Page` set to `-1`. Note that this makes as much request as needed regarding the `PerPage` value. __The API might not react as expected so use this carefully.__


`POST` requests have been implemented but are not handled by any scope yet, and `RequestParams` might be modified in order to ease `POST` requests.
