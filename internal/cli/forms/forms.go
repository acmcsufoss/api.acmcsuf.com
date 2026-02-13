package forms

import (
	"github.com/charmbracelet/huh"
)

// TODO: Use DTO models

func GetIdInteractive() (string, error) {
	var id string
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter resource ID").
				CharLimit(400).
				Value(&id),
		),
	).Run()
	return id, err
}
