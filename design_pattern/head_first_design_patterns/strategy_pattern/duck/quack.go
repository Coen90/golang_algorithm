package duck

import "fmt"

type Quack struct {
}

func (q *Quack) Quack() {
	fmt.Println("꽉꽉!")
}
