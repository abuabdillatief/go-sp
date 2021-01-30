package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Roster ...
type Roster struct {
	Person       Person   `json:"person"`
	JerseyNumber string   `json:"jerseyNumber"`
	Position     Position `json:"position"`
	Link         string   `json:"link"`
}

//Person ...
type Person struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Link     string `json:"link"`
}

//Position ...
type Position struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Abbreviation string `json:"abbreviation"`
}

type nhlRostersResponse struct {
	Rosters []Roster `json:"rosters"`
}

//GetRosters ...
func GetRosters(teamID int) ([]Roster, error) {
	var response nhlRostersResponse
	url := fmt.Sprintf("%s/teams/%d/roster", baseURL, teamID)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)

	return response.Rosters, err
}
