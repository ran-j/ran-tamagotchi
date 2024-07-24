package main

import (
	"fmt"
	"time"

	console_render "ran-tamagoch/internal/console-render"
	"ran-tamagoch/internal/game"
	"ran-tamagoch/internal/tamagotchi"
)

type CLIInput struct{}

func NewCLIInput() *CLIInput {
	return &CLIInput{}
}

func (c *CLIInput) HandleInput(tama *tamagotchi.Tamagotchi) {
	for {
		var input string
		fmt.Scanln(&input)
		switch input {
		case "1":
			tama.Run(tamagotchi.Entree)
		case "2":
			tama.Run(tamagotchi.Snack)
		case "3":
			tama.Run(tamagotchi.Play)
		case "4":
			tama.Run(tamagotchi.Clean)
		case "5":
			tama.Run(tamagotchi.Heal)
		case "6":
			tama.Run(tamagotchi.Reset)
		default:
			fmt.Println("Ação inválida.")
		}
	}
}

func main() {
	cliInput := NewCLIInput()
	g := game.NewGame("Tama")
	render := console_render.NewConsoleRender()

	spriteChannel := g.Tama.AnimationChannel
	statusChannel := make(chan tamagotchi.TamagotchiOutput)

	go func() {
		cliInput.HandleInput(g.Tama)
	}()

	go func() {
		for output := range g.Tama.Output {
			statusChannel <- output
		}
	}()

	go func() {
		for {
			g.Tick()
			time.Sleep(3 * time.Second)
		}
	}()

	go func() {
		for {
			render.Render(spriteChannel, statusChannel, g.Tama)
			// randomSecondsFrom1to60 := rand.Intn(60) + 1
			time.Sleep(1 * time.Second)
		}
	}()

	select {}
}
