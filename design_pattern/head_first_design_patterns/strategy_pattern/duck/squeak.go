package duck

import "fmt"

type Squeak struct {
}

func (s *Squeak) Quack() {
	fmt.Println("삑삑~")
}
