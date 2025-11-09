package main

import (
	"log"

	gochess "github.com/tomascaceres14/go-chess"
)

func GetTurn(turn bool) string {
	str := "white"

	if !turn {
		str = "black"
	}

	return str
}

func main() {

	//reader := bufio.NewReader(os.Stdin)

	game := gochess.NewChessEngine()

	if _, err := game.NewGameFENString("rnbqkbnr/pp1ppppp/8/2p5/4P3/8/PPPP1PPP/RNBQKBNR w KQkq c9 0 2", "whites", "blacks"); err != nil {
		log.Fatal(err)
	}

	// id, err := game.NewGame("whites", "blacks")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println("Partido iniciado. ID:", id)

	// turn := true

	// for {

	// 	fmt.Println("Movimiento de", GetTurn(turn))
	// 	fmt.Print("Desde: ")
	// 	from, _ := reader.ReadString('\n')
	// 	from = strings.TrimSpace(from)

	// 	fmt.Print("Hacia: ")
	// 	to, _ := reader.ReadString('\n')
	// 	to = strings.TrimSpace(to)

	// 	if err := game.Move(from, to, turn); err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}

	// 	turn = !turn

	// }
}
