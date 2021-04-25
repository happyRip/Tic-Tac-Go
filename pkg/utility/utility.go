package utility

type Pair struct {
	X, Y int
}

func ClearBoard() [3][3]string {
	return [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
}
