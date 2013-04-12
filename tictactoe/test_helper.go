package tictactoe

func FakeBoard() []byte {
	mockboard = make([]byte, 9)
	for i := range board {
		mockboard[i] = '1' + byte(i)
	}
	return mockboard
}

func FakePopulatedBoard() {
	board = MakeBoard()
	PopulateBoard(board)
}

var mockboard []byte
var FakeOpenSpacesEmpty = []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
var FakeOpenSpacesNotEmpty = []int{2, 3, 5, 6, 8}

var Row = []int{0, 1, 2}

func PutComputerInBoard() {
	for cell := range board {
		board[cell] = computerMark
	}
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
	for cell := range board {
		board[cell] = '1' + byte(cell)
	}
	for cell := range mockboard {
		mockboard[cell] = '1' + byte(cell)
	}
}

func MakeMoveTestBoard() {
	board[0] = playerMark
	board[5] = computerMark
}

func MakeBlockingMove() []byte {
	mockboard = make([]byte, boardSize)
	PopulateBoard(mockboard)
	mockboard[0] = playerMark
	mockboard[1] = playerMark
	mockboard[4] = computerMark
	return board
}

func OneMoveTestBoard() []byte {
	mockboard = make([]byte, boardSize)
	PopulateBoard(mockboard)
	mockboard[0] = playerMark
	mockboard[4] = computerMark
	return board
}

func MakesWinningMove() []byte {
	mockboard = make([]byte, boardSize)
	PopulateBoard(mockboard)
	mockboard[0] = computerMark
	mockboard[1] = computerMark
	mockboard[3] = playerMark
	mockboard[5] = playerMark
	mockboard[7] = playerMark
	return board
}

func MakesWinningMoveLateGame() []byte {
	mockboard = make([]byte, boardSize)
	PopulateBoard(mockboard)
	mockboard[7] = computerMark
	mockboard[2] = computerMark
	mockboard[1] = playerMark
	mockboard[3] = playerMark
	mockboard[5] = playerMark
	return board
}
