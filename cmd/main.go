package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	gochess "github.com/tomascaceres14/go-chess"
)

func TurnToString(turn bool) string {
	str := "white"

	if !turn {
		str = "black"
	}

	return str
}

func main() {

	game := gochess.NewChessEngine()

	// id, err := game.NewGameFENString("r3kbnr/pppq1ppp/2n5/3ppb2/3PPB2/2N5/PPPQ1PPP/R3KBNR w KQkq - 6 6", "whites", "blacks")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	id, err := game.NewGame("whites", "blacks")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Match started. ID:", id)

	reader := bufio.NewReader(os.Stdin)

	for {

		turn := game.GetTurn()
		fmt.Println(TurnToString(turn), "'s turn")
		fmt.Print("From: ")
		from, _ := reader.ReadString('\n')
		from = strings.TrimSpace(from)

		fmt.Print("To: ")
		to, _ := reader.ReadString('\n')
		to = strings.TrimSpace(to)

		if err := game.Move(from, to, turn); err != nil {
			fmt.Println(err)
			continue
		}

	}
}
