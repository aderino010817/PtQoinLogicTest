package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Player represents a player in the game.
type Player struct {
	name      string
	diceCount int
	points    int
}

// RollDice simulates rolling the dice for a player.
func (p *Player) RollDice() []int {
	diceResults := make([]int, p.diceCount)
	for i := 0; i < p.diceCount; i++ {
		diceResults[i] = rand.Intn(6) + 1
	}
	return diceResults
}

// PlayGame simulates the game with N players and M dice rolls.
func (p *Player) PlayGame(N, M int) {
	players := make([]Player, 0, N)
	for i := 1; i <= N; i++ {
		players = append(players, Player{name: fmt.Sprintf("Pemain #%d", i), diceCount: M})
	}

	round := 1
	for len(players) > 1 {
		fmt.Printf("[\nPemain = %d, Dadu = %d\n==================\n\"Round %d\"\n", N, M, round)

		diceResults := make([]struct {
			player Player
			rolls  []int
		}, 0)

		for i := range players {
			player := &players[i]
			if player.diceCount > 0 {
				playerRoll := player.RollDice()
				fmt.Printf("%s (%d): %v\n", player.name, player.points, playerRoll)
				diceResults = append(diceResults, struct {
					player Player
					rolls  []int
				}{player: *player, rolls: playerRoll})
			}
		}

		fmt.Println("Setelah evaluasi:")
		for _, result := range diceResults {
			for _, roll := range result.rolls {
				if roll == 6 {
					result.player.points++
				} else if roll == 1 {
					nextPlayerIndex := (indexOfPlayer(result.player, players) + 1) % len(players)
					nextPlayer := &players[nextPlayerIndex]
					nextPlayer.diceCount++
					// Hapus angka 1 dari hasil lemparan
					rollIndex := indexOf(1, result.rolls)
					if rollIndex != -1 {
						result.rolls = append(result.rolls[:rollIndex], result.rolls[rollIndex+1:]...)
					}
				}
			}

			validResults := filter(result.rolls, func(result int) bool {
				return result >= 1 && result <= 5
			})

			result.player.diceCount = len(validResults)
			fmt.Printf("%s (%d): %v\n", result.player.name, result.player.points, append(validResults, make([]int, result.player.diceCount-len(validResults))...))
		}

		fmt.Println("==================")
		round++

		for i := len(players) - 1; i >= 0; i-- {
			player := &players[i]
			if player.diceCount == 0 {
				fmt.Printf("%s telah selesai bermain.\n", player.name)
				players = append(players[:i], players[i+1:]...)
			}
		}
	}

	winner := players[0]
	for _, player := range players {
		if player.points > winner.points {
			winner = player
		}
	}
	fmt.Printf("%s adalah pemenang dengan %d poin!\n", winner.name, winner.points)
}

// Helper functions

func indexOfPlayer(player Player, players []Player) int {
	for i, p := range players {
		if p == player {
			return i
		}
	}
	return -1
}

func indexOf(value int, arr []int) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

func filter(arr []int, condition func(int) bool) []int {
	result := make([]int, 0)
	for _, v := range arr {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var p Player
	N := 2 // masukkan jumlah player disini
	M := 4 // masukkan jumlah putaran dadu disini
	p.PlayGame(N, M)
}
