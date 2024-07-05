package duck

import "fmt"

type FlyWithWings struct {
}

func (f *FlyWithWings) fly() {
	fmt.Println("Food Duck!!")
}
