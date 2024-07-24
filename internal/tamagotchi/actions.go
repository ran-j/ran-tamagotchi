package tamagotchi

import (
	Anim "ran-tamagotchi/internal/animation"
)

func (t *Tamagotchi) Feed(option string) {
	if t.Alive {
		switch option {
		case "entree":
			t.Hunger -= 20
		case "snack":
			t.Hunger -= 10
		default:
			t.generateOutput(InvalidFeedOption)
			return
		}
		if t.Hunger < 0 {
			t.Hunger = 0
		}
		t.AnimationController.SetAnimation(Anim.Hungry)
		t.AnimationChannel <- Anim.Hungry
		t.generateOutput(Fed, option)
		t.checkPoop()
		t.checkSick()
	}
}

func (t *Tamagotchi) Play() {
	if t.Alive {
		t.Happiness += 10
		t.Hunger += 5
		t.PlayRequest = false
		if t.Happiness > 100 {
			t.Happiness = 100
		}
		t.AnimationController.SetAnimation(Anim.Playing)
		t.AnimationChannel <- Anim.Playing
		t.generateOutput(Played)
	}
}

func (t *Tamagotchi) Clean() {
	if t.Alive {
		if t.Poop {
			t.Hygiene += 10
			if t.Hygiene > 100 {
				t.Hygiene = 100
			}
			t.Poop = false
			t.PoopCounter = 0
			t.AnimationController.SetAnimation(Anim.Cleaning)
			t.AnimationChannel <- Anim.Cleaning
			t.generateOutput(Cleaned)
		}
	}
}

func (t *Tamagotchi) Heal() {
	if t.Alive {
		if t.Sick {
			t.Sick = false
			t.AnimationController.SetAnimation(Anim.Healing)
			t.AnimationController.SetAnimation(Anim.Idle)
			t.AnimationChannel <- Anim.Healing
			t.generateOutput(Cured)
		}
	}
}

func (t *Tamagotchi) Reset() {
	t.Hunger = 50
	t.Happiness = 50
	t.Hygiene = 50
	t.Age = 0
	t.Alive = true
	t.DiedReason = ""
	t.Sick = false
	t.Poop = false
	t.PoopCounter = 0
	t.PlayRequest = false
	t.AnimationController.SetAnimation(Anim.Resetting)
	t.AnimationController.SetAnimation(Anim.Idle)
	t.AnimationChannel <- Anim.Resetting
	t.generateOutput(Rebooted)
}
