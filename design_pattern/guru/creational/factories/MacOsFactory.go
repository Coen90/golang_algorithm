package factories

import (
	"tucker_algorithm/design_pattern/guru/creational/buttons"
	"tucker_algorithm/design_pattern/guru/creational/checkboxs"
)

type MacOsFactory struct {
}

func (m *MacOsFactory) CreateButton() buttons.Button {
	return &buttons.MacOsButton{}
}

func (m *MacOsFactory) CreateCheckbox() checkboxs.Checkbox {
	return &checkboxs.MacOsCheckbox{}
}
