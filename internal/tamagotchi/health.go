package tamagotchi

import (
	"math/rand"
	Anim "ran-tamagotchi/internal/animation"
)

func (t *Tamagotchi) checkSick() {
	if (t.Hunger < 20 || t.Hygiene < 20 || t.PoopCounter > 5) && rand.Intn(100) < ChangeToSick {
		t.Sick = true
		t.AnimationController.SetAnimation(Anim.Sick)
		t.AnimationChannel <- Anim.Sick
		t.generateOutput(Sick)
	}
}

func (t *Tamagotchi) checkPoop() {
	if rand.Intn(100) < ChangeToPoop {
		t.Poop = true
		t.PoopCounter = 0
		t.AnimationController.SetAnimation(Anim.Pooping)
		t.AnimationChannel <- Anim.Pooping
		t.generateOutput(Pooped)
	}
}

func (t *Tamagotchi) Die(reason string) {
	t.Alive = false
	t.DiedReason = reason
	t.generateOutput(Died, reason)
}
