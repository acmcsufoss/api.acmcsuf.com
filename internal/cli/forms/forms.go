package forms

import (
	"github.com/charmbracelet/huh"
)

func GetIdInteractive() (string, error) {
	var id string
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Announcement ID to delete").
				CharLimit(400).
				Value(&id),
		),
	).Run()
	return id, err
}
