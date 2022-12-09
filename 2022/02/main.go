package main

import (
	"bufio"
	"fmt"
	"os"
)

type Play interface {
	Name() string
	Score() int
	ScoreAgainst(play Play) int
	GetWinningPlay() Play
	GetLosingPlay() Play
	GetDrawPlay() Play
}

const RawPaper = "B"
const PaperName = "P"

type Paper struct{}

func (p Paper) Name() string {
	return RawPaper
}

func (p Paper) Score() int {
	return 2
}

func (p Paper) ScoreAgainst(play Play) int {
	playName := play.Name()
	if playName == RawPaper {
		return 3
	} else if playName == RawScissors {
		return 0
	} else {
		return 6
	}
}

func (p Paper) GetWinningPlay() Play {
	return Scissors{}
}

func (p Paper) GetLosingPlay() Play {
	return Rock{}
}

func (p Paper) GetDrawPlay() Play {
	return Paper{}
}

const RawRock = "A"
const RockName = "R"

type Rock struct{}

func (r Rock) Name() string {
	return RawRock
}

func (r Rock) Score() int {
	return 1
}

func (r Rock) ScoreAgainst(play Play) int {
	playName := play.Name()
	if playName == RawPaper {
		return 0
	} else if playName == RawScissors {
		return 6
	} else {
		return 3
	}
}

func (r Rock) GetWinningPlay() Play {
	return Paper{}
}

func (r Rock) GetLosingPlay() Play {
	return Scissors{}
}

func (r Rock) GetDrawPlay() Play {
	return Rock{}
}

const RawScissors = "C"
const ScissorsName = "S"

type Scissors struct{}

func (s Scissors) Name() string {
	return RawScissors
}

func (s Scissors) Score() int {
	return 3
}

func (s Scissors) ScoreAgainst(play Play) int {
	playName := play.Name()
	if playName == RawPaper {
		return 6
	} else if playName == RawScissors {
		return 3
	} else {
		return 0
	}
}

func (s Scissors) GetWinningPlay() Play {
	return Rock{}
}

func (s Scissors) GetLosingPlay() Play {
	return Paper{}
}

func (s Scissors) GetDrawPlay() Play {
	return Scissors{}
}

type Round struct {
	myself   Play
	opponent Play
}

type Strategy = func(opponent Play) Play
type Strategies = map[string]Strategy

func fromString(raw string) Play {
	if raw == RawRock {
		return Rock{}
	} else if raw == RawPaper {
		return Paper{}
	} else if raw == RawScissors {
		return Scissors{}
	}

	panic(raw)
}

func fromStringWithStrategies(strategies Strategies, raw string, opponent Play) Play {
	strategy, ok := strategies[raw]
	if !ok {
		panic(raw)
	}

	return strategy(opponent)
}

func getGamePlays(filename string, strategies Strategies, c chan Round) {
	readFile, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		opponent := fromString(string(line[0]))
		myself := fromStringWithStrategies(strategies, string(line[2]), opponent)

		c <- Round{myself: myself, opponent: opponent}
	}

	readFile.Close()
	close(c)
}

func getTotalScore(filename string, strategies Strategies) int {
	c := make(chan Round)

	go getGamePlays(filename, strategies, c)

	totalScore := 0

	for round := range c {
		totalScore += round.myself.Score() + round.myself.ScoreAgainst(round.opponent)
	}

	return totalScore
}

func main() {
	input := os.Args[1]

	strategies := Strategies{
		"X": func(opponent Play) Play { return Rock{} },
		"Y": func(opponent Play) Play { return Paper{} },
		"Z": func(opponent Play) Play { return Scissors{} },
	}
	totalScore := getTotalScore(input, strategies)

	fmt.Println(totalScore)

	strategies = Strategies{
		"X": func(opponent Play) Play { return opponent.GetLosingPlay() },
		"Y": func(opponent Play) Play { return opponent.GetDrawPlay() },
		"Z": func(opponent Play) Play { return opponent.GetWinningPlay() },
	}
	totalScore = getTotalScore(input, strategies)

	fmt.Println(totalScore)
}
