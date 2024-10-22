package game

import (
	"ran-tamagotchi/internal/tamagotchi"
)

type Game struct {
	Tama *tamagotchi.Tamagotchi
}

func NewGame(name string) *Game {
	return &Game{
		Tama: tamagotchi.NewTamagotchi(name),
	}
}

func (g *Game) Tick() {
	g.Tama.PassTime()
}
