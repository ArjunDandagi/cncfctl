package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Project struct {
	Name string `json:"name"`
	Repo string `json:"repo_url"`
}

func cncfProjects(c chan []Project) {
	// Retrieve the list of CNCF projects from the Landscape API.
	resp, err := http.Get("https://landscape.cncf.io/data/items.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var cncfProjects []Project

	err = json.Unmarshal(data, &cncfProjects)
	if err != nil {
		log.Fatal(err)
	}
	c <- cncfProjects
	close(c)
}
