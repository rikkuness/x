package popx

import (
	"fmt"
	"regexp"
)

var SQLTemplateFuncs = map[string]interface{}{
	"identifier": Identifier,
}

var identifierPattern = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9_]*$")

func Identifier(i string) (string, error) {
	if !identifierPattern.MatchString(i) {
		return "", fmt.Errorf("invalid SQL identifier '%s'", i)
	}
	return i, nil
}
