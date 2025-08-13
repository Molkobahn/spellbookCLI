package main

import(
	"net/http"
	"encoding/json"
	"fmt"
)

type Spells struct {
	Count   int `json:"count"`
	Results []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		Level int    `json:"level"`
		URL   string `json:"url"`
	} `json:"results"`
}

func GetSpellsRequest(level string) (Spells, error) {
	var url string
	if level != "" {
		url = baseURL +"/spells" + "?level=" + level
	} else {
		url = baseURL + "/spells"
	}
	method := "GET"
	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, nil)


	if err != nil {
		fmt.Println(err)
		return Spells{}, err
	}
	req.Header.Add("Accept", "application/json")


	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Spells{}, err
	}
	defer res.Body.Close()
	
	decoder := json.NewDecoder(res.Body)
	spells := Spells{}
	err = decoder.Decode(&spells)
	if err != nil {
		fmt.Printf("Failed to decode json: %v", err)
		return Spells{}, err
	}
	return spells, nil
}

func commandSpells(cfg * Config, args ...string) error {
	spellLevel := ""
	if len(args) >= 1 {
		spellLevel = args[0]
	}
	spells, err := GetSpellsRequest(spellLevel)
	if err != nil {
		return err
	}
	for i, spell := range spells.Results {
		fmt.Printf("%v. %s\n", i, spell.Name)
	} 
	return nil
}