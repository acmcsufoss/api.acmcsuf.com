package interactors

//go:generate go run ./../../../cmd/generate_openapi_interactors/main.go --output ./generated_interactors.go --package interactors --service ./../../../api/openapi/service.yaml --store ./../../../stores/stores.go --store-interface api.Store --store-package github.com/acmcsufoss/api.acmcsuf.com/stores --store-import github.com/acmcsufoss/api.acmcsuf.com/api
