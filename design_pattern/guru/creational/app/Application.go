package app

import (
	"tucker_algorithm/design_pattern/guru/creational/buttons"
	"tucker_algorithm/design_pattern/guru/creational/checkboxs"
	"tucker_algorithm/design_pattern/guru/creational/factories"
)

type Application struct {
	buttons.Button
	checkboxs.Checkbox
}

func NewApplication(factory factories.GUIFactory) *Application {
	return &Application{
		Button:   factory.CreateButton(),
		Checkbox: factory.CreateCheckbox(),
	}
}

func (a *Application) Paint() {
	a.Button.PaintButton()
	a.Checkbox.PaintCheckbox()
}
