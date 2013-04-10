package tictactoe

func FakeBoard() []byte {
	board = make([]byte, 9)
	for i := range board {
		board[i] = '1' + byte(i)
	}
	return board
}

func FakePopulatedBoard() {
	board = MakeBoard()
	PopulateBoard(board)
}

var Row = []int{0, 1, 2}

func PutComputerInBoard() {
	board[0] = computerMark
	board[1] = computerMark
	board[2] = computerMark
	board[3] = computerMark
	board[4] = computerMark
	board[5] = computerMark
	board[6] = computerMark
	board[7] = computerMark
	board[8] = computerMark
}

func TieGame() {
	board[0] = computerMark
	board[1] = playerMark
	board[2] = computerMark
	board[3] = computerMark
	board[4] = playerMark
	board[5] = playerMark
	board[6] = playerMark
	board[7] = computerMark
	board[8] = playerMark
}

func PutPlayerInBoard() {
	board[0] = playerMark
	board[1] = playerMark
	board[2] = playerMark
}

func PopulateNonWinningBoard() {
	board[0] = playerMark
	board[1] = computerMark
	board[2] = playerMark
}

func ClearBoard() {
	board[0] = 49
	board[1] = 50
	board[2] = 51
	board[3] = 52
	board[4] = 53
	board[5] = 54
	board[6] = 55
	board[7] = 56
	board[8] = 57
}

var mockboard []byte

func OneMoveTestBoard() []byte {
	mockboard = make([]byte, boardSize)
	PopulateBoard(mockboard)
	board[5] = playerMark
	return board
}
