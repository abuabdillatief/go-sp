package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Team ...
type Team struct {
	ID              int        `json:"id"`
	Name            string     `json:"name"`
	Link            string     `json:"link"`
	Venue           Venue      `json:"venue,omitempty"`
	Abbreviation    string     `json:"abbreviation"`
	TeamName        string     `json:"teamName"`
	LocationName    string     `json:"locationName"`
	FirstYearOfPlay string     `json:"firstYearOfPlay"`
	Division        Division   `json:"division"`
	Conference      Conference `json:"conference"`
	Franchise       Franchise  `json:"franchise"`
	ShortName       string     `json:"shortName"`
	OfficialSiteURL string     `json:"officialSiteUrl"`
	FranchiseID     int        `json:"franchiseId"`
	Active          bool       `json:"active"`
}

//Venue ...
type Venue struct {
	Name     string   `json:"name"`
	Link     string   `json:"link"`
	City     string   `json:"city"`
	TimeZone TimeZone `json:"timeZone"`
}

//TimeZone ...
type TimeZone struct {
	ID     string `json:"id"`
	Offset int    `json:"offset"`
	Tz     string `json:"tz"`
}

//Division ...
type Division struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

//Conference ...
type Conference struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

//Franchise ...
type Franchise struct {
	FranchiseID int    `json:"franchiseId"`
	TeamName    string `json:"teamName"`
	Link        string `json:"link"`
}

type nhlTeamsResponse struct {
	Teams []Team `json:"teams"`
}

//GetAllTeams ...
func GetAllTeams() ([]Team, error) {
	var resp nhlTeamsResponse
	res, err := http.Get(fmt.Sprintf("%s/teams", baseURL))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&resp)
	return resp.Teams, err
}
