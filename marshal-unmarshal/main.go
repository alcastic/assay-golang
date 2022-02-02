package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"released"`
	Color  bool     `json:"color,omitempty"`
	Actors []string `json:"actors"`
}

func main() {
	var movies = []Movie{
		{
			Title:  "Casablanca",
			Year:   1942,
			Color:  false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"},
		},
		{
			Title:  "Cool Hand Luke",
			Year:   1967,
			Color:  true,
			Actors: []string{"Paul Newman"},
		},
	}
	data, err := json.MarshalIndent(movies, "", "  ") // json.Marshal(movies) => no identation
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(string(data))

	jsonMovie := "{\"title\":\"Bullit\",\"released\":1968,\"color\":true,\"actors\":[\"Steve MacQueen\",\"Jacqueline Bisset\"]}"
	var movie Movie
	json.Unmarshal([]byte(jsonMovie), &movie)
	fmt.Println(movie)
}
