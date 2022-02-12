package main

import (
	"fmt"
	"sort"
	"strings"
)

type Game struct {
	Rank        int
	Title       string
	Platform    string
	ReleaseYear int
}

func (g *Game) String() string {
	return fmt.Sprintf("#%d - %s - %d - %s", g.Rank, g.Title, g.ReleaseYear, g.Platform)
}

// Basic: sort by field 'Rank', implementing interface sort.Interface
type basicSortGamesByRank []*Game

func (g basicSortGamesByRank) Len() int { return len(g) }

func (g basicSortGamesByRank) Less(i, j int) bool {
	return g[i].Rank < g[j].Rank
}

func (g basicSortGamesByRank) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

// Basic: sort by field 'ReleaceYear', implementing interface sort.Interface
type basicSortGamesByReleaseYear []*Game

func (g basicSortGamesByReleaseYear) Len() int { return len(g) }

func (g basicSortGamesByReleaseYear) Less(i, j int) bool {
	return g[i].ReleaseYear < g[j].ReleaseYear
}

func (g basicSortGamesByReleaseYear) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}

// Advanced: reusing methods Len and Swap when creating a sorting by field
type customGameSort struct {
	games  []*Game
	isLess func(x, y *Game) bool
}

func (g customGameSort) Len() int { return len(g.games) }
func (g customGameSort) Less(i, j int) bool {
	return g.isLess(g.games[i], g.games[j])
}
func (g customGameSort) Swap(i, j int) {
	g.games[i], g.games[j] = g.games[j], g.games[i]
}

func main() {
	var games = []*Game{
		{Rank: 4, Title: "The Legend of Zelda: Majora's Mask", Platform: "Nintendo 64", ReleaseYear: 2000},
		{Rank: 1, Title: "The Legend of Zelda: Ocarina of Time", Platform: "Nintendo 64", ReleaseYear: 1998},
		{Rank: 5, Title: "Super Mario 64", Platform: "Nintendo 64", ReleaseYear: 1996},
		{Rank: 2, Title: "Perfect Dark", Platform: "Nintendo 64", ReleaseYear: 2000},
		{Rank: 3, Title: "GoldenEye 007", Platform: "Nintendo 64", ReleaseYear: 1997},
	}
	var printGames = func(games []*Game) {
		for _, el := range games {
			fmt.Println(el)
		}
	}
	var sortedGames = make([]*Game, len(games))

	fmt.Print("Basic - sortByRank\n")
	copy(sortedGames, games)
	sort.Sort(basicSortGamesByRank(sortedGames))
	printGames(sortedGames)

	fmt.Print("\nBasic - sortByReleaseYear\n")
	copy(sortedGames, games)
	sort.Sort(basicSortGamesByReleaseYear(sortedGames))
	printGames(sortedGames)

	fmt.Print("\nAdvanced - sortByRank\n")
	copy(sortedGames, games)
	sort.Sort(customGameSort{
		games: sortedGames,
		isLess: func(x, y *Game) bool {
			return x.Rank < y.Rank
		},
	})
	printGames(sortedGames)

	fmt.Print("\nAdvanced - sortByReleaseYear\n")
	copy(sortedGames, games)
	sort.Sort(customGameSort{
		games: sortedGames,
		isLess: func(x, y *Game) bool {
			return x.ReleaseYear < y.ReleaseYear
		},
	})
	printGames(sortedGames)

	fmt.Print("\nAdvanced - reverse sortByTitle\n")
	copy(sortedGames, games)
	sort.Sort(sort.Reverse(customGameSort{
		games: sortedGames,
		isLess: func(x, y *Game) bool {
			return strings.Compare(x.Title, y.Title) < 0
		},
	}))
	printGames(sortedGames)

}
