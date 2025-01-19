package buttons

import "fmt"

type WindowOsButton struct {
}

func (w *WindowOsButton) PaintButton() {
	fmt.Println("You have created WindowsOsButton.")
}
