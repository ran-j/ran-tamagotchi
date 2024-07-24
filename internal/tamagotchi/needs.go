package tamagotchi

import (
	"math/rand"
	Anim "ran-tamagotchi/internal/animation"
	"time"
)

func (t *Tamagotchi) PassTime() {
	if t.Alive {
		t.Hunger += 5
		t.Happiness -= 5
		t.Hygiene -= 5
		t.Age += 1

		if t.Hunger > 100 {
			t.Hunger = 100
			t.Die("fome")
		}
		if t.Happiness < 0 {
			t.Happiness = 0
		}
		if t.Hygiene < 0 {
			t.Hygiene = 0
		}

		t.checkSick()

		if t.Poop {
			t.PoopCounter++
		}

		t.checkPlayRequest()

		currentAnimation := t.AnimationController.Update()
		if currentAnimation == Anim.Idle && t.Sick {
			t.AnimationController.SetAnimation(Anim.Sick)
		}

		t.AnimationChannel <- currentAnimation
		t.generateOutput(Status)
	}
}

func (t *Tamagotchi) checkPlayRequest() {
	if !t.PlayRequest && rand.Intn(100) < ChangeToPlay {
		t.PlayRequest = true
		t.PlayRequestTime = time.Now()
		t.generateOutput(WantsToPlay)
	}
	if t.PlayRequest && time.Since(t.PlayRequestTime).Minutes() >= 2 {
		t.Happiness -= 20
		t.PlayRequest = false
		t.generateOutput(VerySad)
	}
}
