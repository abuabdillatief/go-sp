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
	roosterFile, err := os.OpenFile("roosters.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening the fi le roosters.txt: %s", err)
	}
	defer roosterFile.Close()
	w := io.MultiWriter(os.Stdout, roosterFile)
	log.SetOutput(w)

	teams, err := api.GetAllTeams()
	if err != nil {
		log.Fatalf("error while getting all teams: %s", err)
	}
	for _, team := range teams {
		log.Println("===========================")
		log.Printf("Name: %s", team.Name)
		log.Println("===========================")
	}
	log.Printf("took %v", time.Now().Sub(now).String())
}
