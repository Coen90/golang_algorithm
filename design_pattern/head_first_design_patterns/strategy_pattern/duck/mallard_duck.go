package duck

import "fmt"

type mallardDuck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func InitMallardDuck() Duck {
	return &mallardDuck{
		flyBehavior:   new(FlyWithWings),
		quackBehavior: new(Quack),
	}
}

func (m *mallardDuck) Display() {
	fmt.Println("mallardDuck")
}

func (m *mallardDuck) Swim() {
	fmt.Println("어푸어푸")
}

func (m *mallardDuck) PerformFly() {
	m.flyBehavior.fly()
}

func (m *mallardDuck) PerformQuack() {
	m.quackBehavior.Quack()
}

func (m *mallardDuck) SetFlyBehavior(flyBehavior FlyBehavior) {
	m.flyBehavior = flyBehavior
}

func (m *mallardDuck) SetQuackBehavior(quackBehavior QuackBehavior) {
	m.quackBehavior = quackBehavior
}
