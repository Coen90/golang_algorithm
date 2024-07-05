package duck

type Duck interface {
	Swim()
	Display()
	PerformFly()
	PerformQuack()
	SetFlyBehavior(flyBehavior FlyBehavior)
	SetQuackBehavior(quackBehavior QuackBehavior)
}
