package utils

import (
	"fmt"
	"os"

	"github.com/tidwall/pretty"
)

func PrettyPrintJSON(json []byte) {
	prettyJSON := pretty.Pretty(json)
	colorfulJSON := pretty.Color(prettyJSON, nil)
	fmt.Print(string(colorfulJSON))
}

func PrettyPrintJSONErr(json []byte) {
	prettyJSON := pretty.Pretty(json)
	colorfulJSON := pretty.Color(prettyJSON, nil)
	fmt.Fprint(os.Stderr, string(colorfulJSON))
}
