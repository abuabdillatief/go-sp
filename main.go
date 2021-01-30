package main

import (
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/abuabdillatief/go-sp/api"
)

func main() {
	now := time.Now()
	rosterFile, err := os.OpenFile("rosters.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening the fi le rosters.txt: %s", err)
	}
	defer rosterFile.Close()
	w := io.MultiWriter(os.Stdout, rosterFile)
	log.SetOutput(w)

	teams, err := api.GetAllTeams()
	if err != nil {
		log.Fatalf("error while getting all teams: %s", err)
	}
	var wg sync.WaitGroup
	wg.Add(len(teams))
	results := make(chan []api.Roster)

	for _, team := range teams {
		go func(team api.Team) {
			roster, err := api.GetRosters(team.ID)
			if err != nil {
				log.Fatalf("error getting rosters: %v", err)
			}
			results <- roster
			wg.Done()
		}(team)
	}

	display(results)
	wg.Wait()
	close(results)

	log.Printf("took %v", time.Now().Sub(now).String())
}

func display(results chan []api.Roster) {
	for tr := range results {
		for _, r := range tr {
			log.Println("==================")
			log.Printf("Name: %s\n", r.Person.FullName)
			log.Printf("ID: %d\n", r.Person.ID)
			log.Printf("Postion: %s\n", r.Position.Name)
		}
	}
}
