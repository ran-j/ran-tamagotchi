package console_render

import (
	"fmt"
	"os"
	"os/exec"
	Anim "ran-tamagotchi/internal/animation"
	"ran-tamagotchi/internal/tamagotchi"
	"runtime"
)

type ConsoleRender struct{}

func NewConsoleRender() *ConsoleRender {
	return &ConsoleRender{}
}

func (c *ConsoleRender) Render(spriteChan <-chan Anim.Animation, statusChan <-chan tamagotchi.TamagotchiOutput, tamagotchi *tamagotchi.Tamagotchi) {
	var spriteData, statusData string
	spriteReceived, statusReceived := false, false

	for !spriteReceived || !statusReceived {
		select {
		case sprite := <-spriteChan:
			spriteData = MakeTamagochiSprit(sprite)
			spriteReceived = true
		case status := <-statusChan:
			statusData = MakeStatusSprite(status, tamagotchi)
			statusReceived = true
		}
	}

	if spriteData != "" && statusData != "" {
		c.ClearScreen()
		fmt.Println(spriteData)
		fmt.Println(statusData)
		fmt.Println(MakeMenuSprite(tamagotchi))

		statusData = ""
		spriteData = ""
		statusReceived = false
		spriteReceived = false
	}
}

func (c *ConsoleRender) ClearScreen() {
	fmt.Print("\033[H\033[2J")
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func MakeTamagochiSprit(animation Anim.Animation) string {
	switch animation {
	case Anim.Idle:
		return `
		╭─────╮
		│ ^_^ │
		╰─────╯
		`
	case Anim.Hungry:
		return `
		╭─────╮
		│ >_< │
		╰─────╯
		`
	case Anim.Playing:
		return `
		╭─────╮
		│ ^o^ │
		╰─────╯
		`
	case Anim.Cleaning:
		return `
		╭─────╮
		│ ^_^ │
		│  ~  │
		╰─────╯
		`
	case Anim.Healing:
		return `
		╭─────╮
		│ +_+ │
		╰─────╯
		`
	case Anim.Resetting:
		return `
		╭─────╮
		│ O_O │
		╰─────╯
		`
	case Anim.Sick:
		return `
		╭─────╮
		│ x_x │
		╰─────╯
		`
	case Anim.Pooping:
		return `
		╭─────╮
		│ ~_~ │
		│  *  │
		╰─────╯
		`
	default:
		return `
		╭─────╮
		│ ?_? │
		╰─────╯
		`
	}
}

func MakeStatusSprite(output tamagotchi.TamagotchiOutput, tama *tamagotchi.Tamagotchi) string {

	if !tama.Alive {
		return fmt.Sprintf("%s morreu de %s com a idade: %v.", tama.Name, tama.DiedReason, tama.Age)
	}

	stringOutput := fmt.Sprintf("%s's status: Fome: %d, Felicidade: %d, Higiene: %d, Idade: %d, Doente: %t e Cocô: %t.",
		tama.Name, tama.Hunger, tama.Happiness, tama.Hygiene, tama.Age, tama.Sick, tama.Poop)

	stringOutput += "\n"
	switch output.Status {
	case tamagotchi.Sick:
		stringOutput += fmt.Sprintf("%s está doente!", tama.Name)
	case tamagotchi.Died:
		stringOutput += fmt.Sprintf("%s morreu de %s.", tama.Name, output.Options[0])
	case tamagotchi.Pooped:
		stringOutput += fmt.Sprintf("%s fez cocô na tela!", tama.Name)
	case tamagotchi.WantsToPlay:
		stringOutput += fmt.Sprintf("%s quer brincar!", tama.Name)
	case tamagotchi.VerySad:
		stringOutput += fmt.Sprintf("%s está muito triste porque ninguém brincou com ele.", tama.Name)
	case tamagotchi.Rebooted:
		stringOutput += fmt.Sprintf("%s foi resetado.", tama.Name)
	case tamagotchi.Cured:
		stringOutput += fmt.Sprintf("%s foi curado.", tama.Name)
	case tamagotchi.Cleaned:
		stringOutput += fmt.Sprintf("%s foi limpo. Higiene: %d", tama.Name, output.Options[0])
	}

	return stringOutput
}

func MakeMenuSprite(tama *tamagotchi.Tamagotchi) string {
	if !tama.Alive {
		return `6 - Resetar`
	}
	return `
	1 - Dar comida
	2 - Dar petisco
	3 - Brincar
	4 - Limpar
	5 - Curar
	6 - Resetar
	`
}
