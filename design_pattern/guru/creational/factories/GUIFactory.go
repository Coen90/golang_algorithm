package factories

import (
	"tucker_algorithm/design_pattern/guru/creational/buttons"
	"tucker_algorithm/design_pattern/guru/creational/checkboxs"
)

type GUIFactory interface {
	CreateButton() buttons.Button
	CreateCheckbox() checkboxs.Checkbox
}
