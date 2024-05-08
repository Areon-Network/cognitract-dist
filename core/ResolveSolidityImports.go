package core

import "regexp"

// This only resolves imports that begin with "@"
func ResolveSolidityImports(sourceCode string) []string {
	importSyntaxRegex := regexp.MustCompile(`import\s+["'](@[^"']+)["']\s*;`)
	matches := importSyntaxRegex.FindAllStringSubmatch(sourceCode, -1)

	paths := make([]string, len(matches))
	for i, match := range matches {
		paths[i] = match[1]
	}

	return paths
}
