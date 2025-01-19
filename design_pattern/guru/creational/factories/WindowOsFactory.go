package factories

import (
	"tucker_algorithm/design_pattern/guru/creational/buttons"
	"tucker_algorithm/design_pattern/guru/creational/checkboxs"
)

type WindowOsFactory struct {
}

func (w *WindowOsFactory) CreateButton() buttons.Button {
	return &buttons.WindowOsButton{}
}

func (w *WindowOsFactory) CreateCheckbox() checkboxs.Checkbox {
	return &checkboxs.WindowOsCheckbox{}
}
