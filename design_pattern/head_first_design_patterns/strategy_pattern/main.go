package main

import (
	"fmt"
	"tucker_algorithm/head_first_design_patterns/strategy_pattern/duck"
)

func main() {
	var mallardDuck duck.Duck
	mallardDuck = duck.InitMallardDuck()
	mallardDuck.Swim()         //어푸어푸
	mallardDuck.Display()      //mallardDuck
	mallardDuck.PerformQuack() //꽉꽉!
	mallardDuck.PerformFly()   //Food Duck!!
	fmt.Println()
	mallardDuck.SetFlyBehavior(new(duck.FlyNoWay))
	mallardDuck.SetQuackBehavior(new(duck.Squeak))
	mallardDuck.PerformQuack() //삑삑!
	mallardDuck.PerformFly()   //못날아요 ㅠㅠ
}
