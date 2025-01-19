package buttons

import "fmt"

type MacOsButton struct {
}

func (m *MacOsButton) PaintButton() {
	fmt.Println("You have created MacOSButton.")
}
