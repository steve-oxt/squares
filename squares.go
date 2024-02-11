package main

import "fmt"
import "math/rand"
import "github.com/fatih/color"
import "github.com/rodaine/table"

type team struct {
	name   string
	scores [10]int
	score  int
}

func (t team) Init() team {
	fmt.Print("Please enter the team name: ")
	fmt.Scan(&t.name)
	t.scores = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	t.score = 0
	return t
}
func (t *team) Randomize() {
	rand.Shuffle(len(t.scores), func(i, j int) { t.scores[i], t.scores[j] = t.scores[j], t.scores[i] })
}

type Square struct {
	home  team
	away  team
	board [10][10]string
}

func (s Square) Init() Square {
	var teams [2]team
	teams[0] = team{}.Init()
	teams[1] = team{}.Init()
	s.board = setBoard()
	rand.Shuffle(len(teams), func(i, j int) { teams[i], teams[j] = teams[j], teams[i] })
	s.home = teams[0]
	fmt.Println("Home team: ", s.home.name)
	s.away = teams[1]
	fmt.Println("Away team: ", s.away.name)
	s.home.Randomize()
	fmt.Println("Home team scores: ", s.home.scores)
	s.away.Randomize()
	fmt.Println("Away team scores: ", s.away.scores)
	s.PrintBoard()
	return s
}

func (s *Square) PrintBoard() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New(" ", s.home.scores[0], s.home.scores[1], s.home.scores[2], s.home.scores[3], s.home.scores[4], s.home.scores[5], s.home.scores[6], s.home.scores[7], s.home.scores[8], s.home.scores[9])
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)
	for i := 0; i < 10; i++ {
		tbl.AddRow(s.away.scores[i], s.board[i][0], s.board[i][1], s.board[i][2], s.board[i][3], s.board[i][4], s.board[i][5], s.board[i][6], s.board[i][7], s.board[i][8], s.board[i][9])
	}
	tbl.Print()
}

func (s *Square) Play() {
	playing := true
	first_score := true
	for playing {
		fmt.Println("Current score:")
		fmt.Println(s.home.name,": ", s.home.score, " - ",s.away.name, ": ", s.away.score)
		if !first_score {
			h := -1
			v := -1
			for i := 0; i < 10; i++ {
				if s.home.scores[i] == s.home.score {
					h = i
				}
				if s.away.scores[i] == s.away.score {
					v = i
				}
			}
			
		}
		fmt.Printf("\n1. %s score Safety\n", s.home.name)
		fmt.Printf("2. %s score Field Goal\n", s.home.name)
		fmt.Printf("3. %s score Touchdown, Conversion No Good\n", s.home.name)
		fmt.Printf("4. %s score Touchdown, PAT Good\n", s.home.name)
		fmt.Printf("5. %s score Touchdown, 2 Point Conversion Good\n", s.home.name)
		fmt.Printf("6. %s score Safety\n", s.away.name)
		fmt.Printf("7. %s score Field Goal\n", s.away.name)
		fmt.Printf("8. %s score Touchdown, Conversion No Good\n", s.away.name)
		fmt.Printf("9. %s score Touchdown, PAT Good\n", s.away.name)
		fmt.Printf("10. %s score Touchdown, 2 Point Conversion Good\n", s.away.name)
		fmt.Printf("11. Game Over\n")
		fmt.Print("Enter the number of the play: ")
		var play int
		fmt.Scan(&play)
		switch play {
		case 1:
			s.home.score += 2
			if first_score {
				first_score = false
			}
		case 2:
			s.home.score += 3
			if first_score {
				first_score = false
			}
		case 3:
			s.home.score += 6
			if first_score {
				first_score = false
			}
		case 4:
			s.home.score += 7
			if first_score {
				first_score = false
			}
		case 5:
			s.home.score += 8
			if first_score {
				first_score = false
			}
		case 6:
			s.away.score += 2
			if first_score {
				first_score = false
			}
		case 7:
			s.away.score += 3
			if first_score {
				first_score = false
			}
		case 8:
			s.away.score += 6
			if first_score {
				first_score = false
			}
		case 9:
			s.away.score += 7
			if first_score {
				first_score = false
			}
		case 10:
			s.away.score += 8
			if first_score {
				first_score = false
			}
		case 11:
			fmt.Println("Game Over")
			playing = false
		default:
			fmt.Println("Invalid play")
		}
	}
}

func setBoard() [10][10]string {
	var board [10][10]string
	for entries := 0; entries < 100; entries++ {
		var i, j int
		fmt.Print("Type two numbers: ")
		fmt.Scan(&i, &j)
		if board[i][j] != "" {
			fmt.Println("This square is already taken")
			entries--
			continue
		}
		var name string
		fmt.Print("Enter the name of the player: ")
		fmt.Scanf("%s", &name)
		board[i][j] = name
	}
	return board
}

func main() {
	s := Square{}
	s.Init()
}

/*func (t team) Score() int {
	var score int
	fmt.Print("Please enter the score: ")
	fmt.Scan(&score)
	return score
}*/
