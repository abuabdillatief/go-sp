package main

import (
	"io"
	"log"
	"os"
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
	for _, team := range teams {
		log.Println("===========================")
		log.Printf("Name: %s", team.Name)
		log.Printf("Venue: %v", team.Venue)
		log.Println("===========================")
	}
	log.Printf("took %v", time.Now().Sub(now).String())
}
