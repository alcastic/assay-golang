package main

import (
	"fmt"
	"sort"
)

type Game struct {
	Rank        int
	Title       string
	Platform    string
	ReleaseYear int
}

func (g *Game) String() string {
	return fmt.Sprintf("#%d - %s", g.Rank, g.Title)
}

// Basic sort by field 'Rank', implementing interface sort.Interface
type GamesByRank []*Game

func (g GamesByRank) Len() int { return len(g) }

func (g GamesByRank) Less(i, j int) bool {
	return g[i].Rank < g[j].Rank
}

func (g GamesByRank) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

func printGames(games []*Game) {
	for _, el := range games {
		fmt.Println(el)
	}
}

func main() {
	var games = []*Game{
		{Rank: 4, Title: "The Legend of Zelda: Majora's Mask", Platform: "Nintendo 64", ReleaseYear: 2000},
		{Rank: 1, Title: "The Legend of Zelda: Ocarina of Time", Platform: "Nintendo 64", ReleaseYear: 1998},
		{Rank: 5, Title: "Super Mario 64", Platform: "Nintendo 64", ReleaseYear: 1996},
		{Rank: 2, Title: "Perfect Dark", Platform: "Nintendo 64", ReleaseYear: 2000},
		{Rank: 3, Title: "GoldenEye 007", Platform: "Nintendo 64", ReleaseYear: 1997},
	}

	var sortedGames = make([]*Game, len(games))
	copy(sortedGames, games)
	sort.Sort(GamesByRank(sortedGames))

	printGames(sortedGames)
}
