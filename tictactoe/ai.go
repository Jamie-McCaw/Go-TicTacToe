package tictactoe

import (
	"math"
)

var PosInfinity float64 = math.Inf(1)
var NegInfinity float64 = math.Inf(-1)

const Zero = 0
const PosScore = 10000
const NegScore = -10000

func PlaceMoveForAi(move int, piece byte) {
	board[move] = piece
}

func UndoAiMove(move int) {
	board[move] = '1' + byte(move)
}

func Opponent(piece byte) byte {
	if piece == playerMark {
		return computerMark
	}
	return playerMark
}

func BetterMove(score float64, bestScore float64) bool {
	return score >= bestScore
}

func ScoreIsNotBetter(alpha float64, beta float64) bool {
	return alpha >= beta
}

func HaveNotFoundBestMove(score float64, alpha float64) bool {
	return score > alpha
}

func CheckForWin(piece byte) int {
	if GameIsWon() {
		if WinningPlayer == piece {
			return PosScore
		} else if WinningPlayer == Opponent(piece) {
			return NegScore
		}
	}
	return 0
}

func GetBestMove(index int, piece byte, alpha float64, beta float64) float64 {
	PlaceMoveForAi(index, piece)
	score := -AlphaBeta(Opponent(piece), alpha, beta)
	UndoAiMove(index)
	return score
}

func AlphaBeta(piece byte, alpha float64, beta float64) float64 {
	openMoves, newAlpha := OpenMoves(), alpha

	if GameOver() {
		return float64(CheckForWin(piece))
	} else {
		for index := range openMoves {
			score := GetBestMove(openMoves[index], piece, -beta, -newAlpha)
			if HaveNotFoundBestMove(score, newAlpha) {
				newAlpha = score
			}
			if ScoreIsNotBetter(newAlpha, beta) {
				return newAlpha
			}
		}
	}
	return newAlpha
}

func ComputerTurn() {
	openMoves, bestScore, bestMove := OpenMoves(), NegInfinity, Zero

	for index := range openMoves {
		score := GetBestMove(openMoves[index], computerMark, NegInfinity, PosInfinity)
		if BetterMove(score, bestScore) {
			bestScore = score
			bestMove = openMoves[index]
		}
	}
	PlaceMoveForAi(bestMove, computerMark)
}
