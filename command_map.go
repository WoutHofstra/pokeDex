package main

import (
	"fmt"
	"net/http"
	"io"
	"encoding/json"
)

type LocationAreaResponse struct {
	Count 		int		`json:"count"`
	Next		*string		`json:"next"`
	Previous	*string		`json:"previous"`
	Results 	[]Result	`json:"results"`
}

type Result struct {
	Name string	`json:"name"`
	Url string	`json:"url"`
}

type Config struct {
	Next *string
	Previous *string
}


func commandMap(config *Config) error {

	var res *http.Response
	var err error


	if config.Next == nil && config.Previous == nil {
		res, err = http.Get("https://pokeapi.co/api/v2/location-area/")
	} else {
		res, err = http.Get(*config.Next)
	}
	if err != nil {
                return err
        }
        defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var AreaResponse LocationAreaResponse
	err = json.Unmarshal(data, &AreaResponse)
	if err != nil {
		return err
	}

	config.Next = AreaResponse.Next
	config.Previous = AreaResponse.Previous

	for _, r := range AreaResponse.Results {
		fmt.Println(r.Name)
	}
	return nil
}

func commandMapBack(config *Config) error {

	if config.Previous == nil {
		fmt.Println("You're on the first page!")
		return nil
	}

	res, err := http.Get(*config.Previous)
        if err != nil {
                return err
        }

        defer res.Body.Close()

        data, err := io.ReadAll(res.Body)
        if err != nil {
                return err
        }

        var AreaResponse LocationAreaResponse
        err = json.Unmarshal(data, &AreaResponse)
        if err != nil {
                return err
        }

        config.Next = AreaResponse.Next
        config.Previous = AreaResponse.Previous

        for _, r := range AreaResponse.Results {
                fmt.Println(r.Name)
        }
        return nil
}
