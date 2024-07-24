package tamagotchi

import (
	"fmt"
	Anim "ran-tamagoch/internal/animation"
	"time"
)

const (
	ChangeToSick = 20
	ChangeToPlay = 20
	ChangeToDie  = 0
	ChangeToPoop = 30
)

type TamagotchiAction int

const (
	Entree TamagotchiAction = iota
	Snack
	Play
	Clean
	Heal
	Reset
)

func (a TamagotchiAction) String() string {
	return [...]string{"entree", "snack", "play", "clean", "heal", "reset"}[a]
}

type TamagotchiOutputTypes int

const (
	Sick TamagotchiOutputTypes = iota
	Status
	WantsToPlay
	VerySad
	Rebooted
	Cured
	Cleaned
	Played
	Fed
	Pooped
	Died
	InvalidFeedOption
)

type TamagotchiOutput struct {
	Message string                `json:"message"`
	Details string                `json:"details,omitempty"`
	Status  TamagotchiOutputTypes `json:"status,omitempty"`
	Options []interface{}         `json:"options,omitempty"`
}

type Tamagotchi struct {
	Name                string
	Hunger              int
	Happiness           int
	Hygiene             int
	Age                 int
	Alive               bool
	DiedReason          string
	Sick                bool
	Poop                bool
	PoopCounter         int
	PlayRequest         bool
	PlayRequestTime     time.Time
	Output              chan TamagotchiOutput
	AnimationController *Anim.AnimationController
	AnimationChannel    chan Anim.Animation
}

func NewTamagotchi(name string) *Tamagotchi {
	return &Tamagotchi{
		Name:                name,
		Hunger:              0,
		Happiness:           100,
		Hygiene:             50,
		Age:                 0,
		Alive:               true,
		Sick:                false,
		Poop:                false,
		AnimationController: Anim.NewAnimationController(),
		Output:              make(chan TamagotchiOutput),
		AnimationChannel:    make(chan Anim.Animation),
	}
}

func (t *Tamagotchi) Run(action TamagotchiAction) {
	switch action {
	case Entree:
		t.Feed("entree")
	case Snack:
		t.Feed("snack")
	case Play:
		t.Play()
	case Clean:
		t.Clean()
	case Heal:
		t.Heal()
	case Reset:
		t.Reset()
	}
}

func (t *Tamagotchi) generateOutput(condition TamagotchiOutputTypes, options ...string) {
	switch condition {
	case Sick:
		t.Output <- TamagotchiOutput{
			Message: fmt.Sprintf("%s está doente!", t.Name),
			Status:  condition,
		}
	case Died:
		if len(options) == 0 {
			panic("No reason provided")
		}
		reason := options[0]
		t.Output <- TamagotchiOutput{
			Message: fmt.Sprintf("%s morreu de %s.", t.Name, reason),
			Status:  condition,
			Options: []interface{}{reason},
		}
	case Pooped:
		t.Output <- TamagotchiOutput{
			Message: fmt.Sprintf("%s fez cocô na tela!", t.Name),
			Status:  condition,
		}
	case Status:
		t.Output <- TamagotchiOutput{
			Message: fmt.Sprintf("%s's status", t.Name),
			Details: fmt.Sprintf("Fome: %d, Felicidade: %d, Higiene: %d, Idade: %d, Doente: %t", t.Hunger, t.Happiness, t.Hygiene, t.Age, t.Sick),
			Status:  condition,
			Options: []interface{}{t.Hunger, t.Happiness, t.Hygiene, t.Age, t.Sick},
		}
	case WantsToPlay:
		t.Output <- TamagotchiOutput{
			Message: fmt.Sprintf("%s quer brincar!", t.Name),
			Status:  condition,
		}
	case VerySad:
		t.Output <- TamagotchiOutput{
			Message: fmt.Sprintf("%s está muito triste porque ninguém brincou com ele.", t.Name),
			Status:  condition,
		}
	case Rebooted:
		t.Output <- TamagotchiOutput{
			Message: fmt.Sprintf("%s foi resetado.", t.Name),
			Status:  condition,
		}
	case Cured:
		t.Output <- TamagotchiOutput{
			Message: fmt.Sprintf("%s foi curado.", t.Name),
			Status:  condition,
		}
	case Cleaned:
		t.Output <- TamagotchiOutput{
			Message: fmt.Sprintf("%s foi limpo. Higiene: %d", t.Name, t.Hygiene),
			Status:  condition,
			Options: []interface{}{t.Hygiene},
		}
	case Played:
		t.Output <- TamagotchiOutput{
			Message: fmt.Sprintf("%s brincou. Felicidade: %d", t.Name, t.Happiness),
			Status:  condition,
			Options: []interface{}{t.Happiness},
		}
	case Fed:
		food := "food"
		if len(options) == 0 {
			panic("No food option provided")
		}
		food = options[0]
		t.Output <- TamagotchiOutput{
			Message: fmt.Sprintf("%s foi alimentado com %s. Fome: %d", t.Name, food, t.Hunger),
			Status:  condition,
			Options: []interface{}{food},
		}
	case InvalidFeedOption:
		t.Output <- TamagotchiOutput{
			Message: "Opção inválida.",
			Status:  condition,
		}
	default:
		panic("Invalid condition")
	}
}
