package animation

type Animation int

const (
	Idle Animation = iota
	Hungry
	Playing
	Cleaning
	Healing
	Resetting
	Sick
	Pooping
)

type AnimationController struct {
	Current Animation
	Queue   []Animation
}

func NewAnimationController() *AnimationController {
	return &AnimationController{
		Current: Idle,
		Queue:   make([]Animation, 0),
	}
}

func (ac *AnimationController) SetAnimation(animation Animation) {
	ac.Queue = append(ac.Queue, animation)
}

func (ac *AnimationController) Update() Animation {
	if len(ac.Queue) > 0 {
		ac.Current = ac.Queue[0]
		ac.Queue = ac.Queue[1:]
	}
	return ac.Current
}
