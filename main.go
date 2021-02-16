package main

import(
	"fmt"
	"log"
	"strconv"
)

func main() {
	game := Game{
		IsCircle: true,
		PlayerN: 1,
		Board: [][]rune{{'1','2','3'}, {'4','5','6'}, {'7','8','9'}},
		Win: [][]rune{{'1','2','3'}, {'4','5','6'}, {'7','8','9'}, {'1','4','7'}, {'2','5','8'}, {'3','6','9'}, {'1','5','9'}, {'3','5','7'}},
	}

	fmt.Println(game.Win)

	game.GetPlayerN()
	game.GetPlayerNames()
/*
	game.Draw()
	game.PrintScores()
	game.Move()
*/
	game.PlayRound()
}

func IsNumber(s string) (int, bool) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, false
	} else {
		return i,true
	}
}

type Game struct {
	IsCircle bool
	PlayerN int
	Score []int
	Players []string
	Board [][]rune
	Win [][]rune
}

func (g * Game) GetPlayerN() {
	for {
		fmt.Print("Define number of players: ")
		var temp string
		_, err := fmt.Scanln(&temp)
		if err != nil {
			log.Println("err: Unable to scan number of players!")
		}
		i, b := IsNumber(temp)
		if b && i >= 1 && i <= 2 {
			(*g).PlayerN = i
			break
		} else {
			fmt.Println("Enter a valid number (1 or 2)")
		}
	}
}

func (g Game) PrintScores() {
	fmt.Println("Scoreboard:")
	for i, v := range g.Score {
		fmt.Print("\t", g.Players[i], ": ", v, "\n")
	}
}

func (g * Game) GetPlayerNames() {
	for i := 0; i < g.PlayerN; i++ {
		var temp string
		fmt.Print("Player", i+1, " name: ")
		fmt.Scanln(&temp)
		g.Players = append(g.Players, temp)
		g.Score = append(g.Score, 0)
	}
}

func (g Game) Draw() {
	for i, row := range g.Board {
		fmt.Print("\t")
		for j, v := range row {
			fmt.Print(string(v))
			if j < len(row)-1 {
				fmt.Print(" | ")
			}
		}
		fmt.Printf("\n")
		if i < len(g.Board)-1 {
			fmt.Print("\t")
		}
		if i < len(g.Board)-1 {
			for k := 0; k < len(row); k++ {
				fmt.Print("--")
				if k < len(row)-1 {
					if k > 0 {
						fmt.Print("-")
					}
					fmt.Print("+")
				}
			}
			fmt.Print("\n")
		}
	}
}

func (g * Game) Move() {
	if g.IsCircle {
		fmt.Print(g.Players[0])
	} else {
		fmt.Print(g.Players[1])
	}
	fmt.Println("'s turn!")

	fmt.Print("Select where to place ")
	if g.IsCircle {
		fmt.Print("o: ")
	} else {
		fmt.Print("x: ")
	}
	for {	
		var temp string
		_, err := fmt.Scanln(&temp)
		if err != nil {
			log.Println("err: Unable to select a spot!")
		}
		i, b := IsNumber(temp)
		if b {
			if i > 9 || i < 1 {
				fmt.Print("\tThe number must be between 1 and 9.\nTry again: ")
			} else {
				// 1 2 3
				// 4 5 6
				// 7 8 9

				// 1 2 3
				// 1 2 3
				// 1 2 3

				// 1
				// 2
				// 3
				x, y := (i+2)/3, i%3
				if y == 0 {
					y = 3
				}
				x, y = x-1, y-1
				fmt.Println("\t\tx:", x, "y:", y)
				if g.Board[x][y] != rune(i+'0') {
					fmt.Print("\tThis slot is already taken: ", x, " ", y, " ", g.Board[x][y],".\nTry again: ")
				} else {
					if g.IsCircle {
						(*g).Board[x][y] = 'o'
						(*g).IsCircle = false
					} else {
						(*g).Board[x][y] = 'x'
						(*g).IsCircle = true
					}
					//g.Draw()
					break
				}
			}
		} else {
			fmt.Print("\tThat is not a correct number.\nTry again: ")
		}
	}
}

func (g * Game) CheckWin() {
	
}

func (g * Game) PlayRound() {
	g.Draw()
  // maximum number of moves is 9
	for i := 0; i < 9; i++ {
		g.Move()
		g.Draw()
	}
}
