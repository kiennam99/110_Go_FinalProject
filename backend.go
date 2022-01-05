package main

//"strings"

type Chess struct {
	chess_type string
	party      int
	moved      bool
}

type Game struct {
	board                      [8][8]Chess
	id                         int
	player1                    string
	player2                    string
	last_moved_x, last_moved_y int
	turn                       int
}

func abs(n int) int {
	if n < 0 {
		n = (-n)
	}
	return n
}

func ternary(flag bool, state1, state2 interface{}) interface{} {
	if flag {
		return state1
	}
	return state2
}

func (gm *Game) init(Id int, PlayerName1 string, PlayerName2 string) bool {
	order := []string{"Rook", "Knight", "Bishop", "King", "Queen", "Bishop", "Knight", "Rook"}
	for i := 0; i < 8; i++ {
		gm.board[0][i] = Chess{
			chess_type: order[i],
			party:      0,
			moved:      false,
		}
	}
	for i := 0; i < 8; i++ {
		gm.board[7][i] = Chess{
			chess_type: order[i],
			party:      1,
			moved:      false,
		}
	}
	for i := 0; i < 8; i++ {
		gm.board[1][i] = Chess{
			chess_type: "Pawn",
			party:      0,
			moved:      false,
		}
		gm.board[6][i] = Chess{
			chess_type: "Pawn",
			party:      1,
			moved:      false,
		}
	}
	for i := 2; i < 6; i++ {
		for j := 0; j < 8; j++ {
			gm.board[i][j] = Chess{
				chess_type: "blank",
				party:      3,
				moved:      true,
			}
		}
	}
	gm.player1 = PlayerName1
	gm.player2 = PlayerName2
	gm.id = Id
	gm.last_moved_x, gm.last_moved_y = -1, -1
	gm.turn = 1
	return true
}

func (gm *Game) move(cor1, cor2 string) bool { //move(x1 int, y1 int, x2 int, y2 int) bool {
	cory := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"D": 3,
		"E": 4,
		"F": 5,
		"G": 6,
		"H": 7,
	}
	corx := map[string]int{
		"1": 7,
		"2": 6,
		"3": 5,
		"4": 4,
		"5": 3,
		"6": 2,
		"7": 1,
		"8": 0,
	}
	x1, x2 := corx[cor1[1:2]], corx[cor2[1:2]]
	y1, y2 := cory[cor1[0:1]], cory[cor2[0:1]]

	if x1 < 0 || x1 > 7 || x2 < 0 || x2 > 7 || y1 < 0 || y1 > 7 || y2 < 0 || y2 > 7 || (x1 == x2 && y1 == y2) {
		return false
	} else if gm.board[x1][y1].party == gm.board[x2][y2].party {
		return false
	} else if gm.board[x1][y1].chess_type == "blank" {
		return false
	} else if gm.board[x1][y1].party != gm.turn {
		return false
	}
	cur := gm.board[x1][y1]
	pawn_move_two := true
	if cur.chess_type == "Rook" {
		if x1 != x2 && y1 != y2 {
			return false
		}
		isBlocked := false
		if x1 == x2 {
			for i := ternary(y1 < y2, 1, -1).(int); y1+i != y2; i += ternary(y1 < y2, 1, -1).(int) {
				if gm.board[x1][y1+i].chess_type != "blank" {
					isBlocked = true
				}
			}
		} else if y1 == y2 {
			for i := ternary(x1 < x2, 1, -1).(int); x1+i != x2; i += ternary(x1 < x2, 1, -1).(int) {
				if gm.board[x1+i][y1].chess_type != "blank" {
					isBlocked = true
				}
			}
		}
		if isBlocked {
			return false
		}
	} else if cur.chess_type == "Knight" {
		if x1 == x2 || y1 == y2 {
			return false
		} else if abs(x1-x2)+abs(y1-y2) != 3 {
			return false
		}
	} else if cur.chess_type == "Bishop" {
		if x1 == x2 || y1 == y2 {
			return false
		} else if abs(x1-x2) != abs(y1-y2) {
			return false
		}
		isBlocked := false
		for i, j := ternary(x1 < x2, 1, -1).(int), ternary(y1 < y2, 1, -1).(int); x1+i != x2 && y1+j != y2; i, j = i+ternary(x1 < x2, 1, -1).(int), j+ternary(y1 < y2, 1, -1).(int) {
			if gm.board[x1+i][y1+j].chess_type != "blank" {
				isBlocked = true
			}
		}
		if isBlocked {
			return false
		}
	} else if cur.chess_type == "King" {
		if abs(x1-x2) > 1 || abs(y1-y2) > 1 {
			return false
		}
	} else if cur.chess_type == "Queen" {
		if x1 != x2 && y1 != y2 && (abs(x1-x2) != abs(y1-y2)) {
			return false
		}
		isBlocked := false
		if x1 == x2 {
			for i := ternary(y1 < y2, 1, -1).(int); y1+i != y2; i += ternary(y1 < y2, 1, -1).(int) {
				if gm.board[x1][y1+i].chess_type != "blank" {
					isBlocked = true
				}
			}
		} else if y1 == y2 {
			for i := ternary(x1 < x2, 1, -1).(int); x1+i != x2; i += ternary(x1 < x2, 1, -1).(int) {
				if gm.board[x1+i][y1].chess_type != "blank" {
					isBlocked = true
				}
			}
		} else {
			for i, j := ternary(x1 < x2, 1, -1).(int), ternary(y1 < y2, 1, -1).(int); x1+i != x2 && y1+j != y2; i, j = i+ternary(x1 < x2, 1, -1).(int), j+ternary(y1 < y2, 1, -1).(int) {
				if gm.board[x1+i][y1+j].chess_type != "blank" {
					isBlocked = true
				}
			}
		}
		if isBlocked {
			return false
		}
	} else if cur.chess_type == "Pawn" {
		tar := gm.board[x2][y2]
		if cur.party == 0 {
			if !cur.moved && x2 == x1+2 && y1 == y2 {
				if gm.board[x1+1][y1].chess_type != "blank" || gm.board[x1+2][y1].chess_type != "blank" {
					return false
				} else {
					gm.last_moved_x = x2
					gm.last_moved_y = y2
					pawn_move_two = false
				}
			} else if x2 == x1+1 && y1 == y2 {
				if gm.board[x1+1][y1].chess_type != "blank" {
					return false
				}
			} else if x2 == x1+1 && abs(y1-y2) == 1 {
				if tar.chess_type == "blank" {
					if gm.last_moved_x != -1 {
						if gm.board[gm.last_moved_x][gm.last_moved_y].party == 0 || gm.last_moved_x != x1 {
							return false
						} else {
							gm.board[gm.last_moved_x][gm.last_moved_y] = Chess{
								chess_type: "blank",
								party:      3,
								moved:      true,
							}
						}
					} else {
						return false
					}
				}
			} else {
				return false
			}
		} else {
			if !cur.moved && x2 == x1-2 && y1 == y2 {
				if gm.board[x1-1][y1].chess_type != "blank" || gm.board[x1-2][y1].chess_type != "blank" {
					return false
				} else {
					gm.last_moved_x = x2
					gm.last_moved_y = y2
					pawn_move_two = false
				}
			} else if x2 == x1-1 && y1 == y2 {
				if gm.board[x1-1][y1].chess_type != "blank" {
					return false
				}
			} else if x2 == x1-1 && abs(y1-y2) == 1 {
				if tar.chess_type == "blank" {
					if gm.last_moved_x != -1 {
						if gm.board[gm.last_moved_x][gm.last_moved_y].party == 1 || gm.last_moved_x != x1 {
							return false
						} else {
							gm.board[gm.last_moved_x][gm.last_moved_y] = Chess{
								chess_type: "blank",
								party:      3,
								moved:      true,
							}
						}
					} else {
						return false
					}
				}
			} else {
				return false
			}
		}
	}
	cur.moved = true
	gm.board[x2][y2] = cur
	gm.board[x1][y1] = Chess{
		chess_type: "blank",
		party:      3,
		moved:      true,
	}
	if pawn_move_two {
		gm.last_moved_x, gm.last_moved_y = -1, -1
	}
	gm.turn = ternary(gm.turn == 0, 1, 0).(int)
	return true
}

/*func (gm *Game) castling(x, y int) bool {
	if x < 0 || x > 7 || y < 0 || y > 7 {
		return false
	} else if gm.board[x][y].chess_type == "blank" {
		return false
	} else if gm.board[x][y].party != gm.turn || gm.board[x][y].chess_type != "Rook" || gm.board[x][y].moved {
		return false
	}
	if gm.board[x][y].party == 0 {

	}
}*/

func (gm *Game) promotion(x, y int, name string) bool {
	if gm.board[x][y].chess_type != "Pawn" {
		return false
	} else {
		if gm.board[x][y].party == 0 && x == 7 {
			gm.board[x][y].chess_type = name
		} else if gm.board[x][y].party == 1 && x == 0 {
			gm.board[x][y].chess_type = name
		} else {
			return false
		}
	}
	return true
}

func (gm Game) print(function func(string, ...interface{})) {
	function("   |")
	for i := 0; i < 8; i++ {
		function(" %7d |", i)
	}
	function("\n")
	for id1, i := range gm.board {
		function("%d  |", id1)
		for _, j := range i {
			function(" %7s |", j.chess_type)
		}
		function("\n")
	}
	function("----------------------------------------------------------------------------------------------------\n")
}

func (gm Game) winner() int {
	p1, p2 := true, true
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if gm.board[i][j].chess_type == "King" {
				if gm.board[i][j].party == 0 {
					p1 = false
				} else if gm.board[i][j].party == 1 {
					p2 = false
				}
			}
		}
	}
	if p1 != p2 {
		if p1 {
			return 0
		} else {
			return 1
		}
	} else {
		return 3
	}
}

/*
func main() {
	var gm Game
	gm.init(5, "A", "B")
	gm.print()
	for gm.winner() == 3 {
		//var x1, x2, y1, y2 int
		var cor1, cor2 string
		//fmt.Scanln(&x1, &y1, &x2, &y2)
		fmt.Scanln(&cor1, &cor2)
		if gm.move(cor1, cor2) {
			gm.print(fmt.Printf)
		} else {
			fmt.Printf("Invalid operation\n")
		}
	}
}
*/
