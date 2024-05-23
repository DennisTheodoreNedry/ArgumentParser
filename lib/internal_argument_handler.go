package lib

import (
	"fmt"

	gotools "github.com/DennisTheodoreNedry/Go-tools"
)

// Internal handler that all 'add' functions really depend on
func (handler *ArgumentHandler) internalArgumentHandler(new_argument Argument_t) {
	if handler.utilize_dashes {
		// Check if the argument starts with -- or -
		longResult := gotools.Contains(new_argument.long_name, []string{"--"})
		shortResult := gotools.Contains(new_argument.short_name, []string{"-"})

		if ok := longResult["--"]; !ok {
			new_argument.long_name = fmt.Sprintf("--%s", new_argument.long_name)
		}

		if ok := shortResult["-"]; !ok {
			new_argument.short_name = fmt.Sprintf("-%s", new_argument.short_name)
		}
	}

	handler.arguments = append(handler.arguments, new_argument)
}
