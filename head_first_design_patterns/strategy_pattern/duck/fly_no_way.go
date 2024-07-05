package duck

import "fmt"

type FlyNoWay struct {
}

func (f *FlyNoWay) fly() {
	fmt.Println("못날아요 ㅠㅠ")
}
