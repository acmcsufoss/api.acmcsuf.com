package interactors

//go:generate -command interactors_generator go run ./../../../cli/interactors_generator/cmd
//go:generate interactors_generator --config ./interactors.json --output ./interactors_generated.go