package utils

import (
	"fmt"

	"github.com/tidwall/pretty"
)

func PrettyPrintJSON(json []byte) {
	prettyJSON := pretty.Pretty(json)
	colorfulJSON := pretty.Color(prettyJSON, nil)
	fmt.Print(string(colorfulJSON))
}
