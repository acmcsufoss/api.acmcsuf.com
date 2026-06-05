package linter

//
// Rules for linting
// Right now just deals with bad imports but can be added onto later
//
type lintRule struct {

	// Path to check
	path string

	// Files to skip in a path
	skipFiles []string

	// Imports that should not be in a file
	badImports []string

}

func LinterRules() []lintRule {
	var rules []lintRule

	// ====================== API RULES ====================== 
	handlerRule := lintRule{
		path: "internal/api/handlers/",
		skipFiles: []string{".", "swagger.go"},
		badImports: []string{"github.com/acmcsufoss/api.acmcsuf.com/internal/api/store/dbmodels"}, 
	}
	rules = append(rules, handlerRule)

	// CLI and API linter rules could probably be split in the future for faster test times

	// ====================== CLI RULES ====================== 
	cliPath := "internal/cli" 
	announcementsCLIRule := lintRule{
		path: cliPath + "/announcements/",
		skipFiles: []string{".", "root.go"},
		badImports: []string{"github.com/acmcsufoss/api.acmcsuf.com/internal/api/store/dbmodels"}, 
	}

	rules = append(rules, announcementsCLIRule)
	
	eventsCLIRule := lintRule{
		path: cliPath + "/events/",
		skipFiles: []string{".", "root.go"},
		badImports: []string{"github.com/acmcsufoss/api.acmcsuf.com/internal/api/store/dbmodels"}, 
	}

	rules = append(rules, eventsCLIRule)

	officersCLIRule := lintRule{
		path: cliPath + "/officers/",
		skipFiles: []string{".", "root.go"},
		badImports: []string{"github.com/acmcsufoss/api.acmcsuf.com/internal/api/store/dbmodels"}, 
	}

	rules = append(rules, officersCLIRule)

	return rules
}

