package lib

import (
	"fmt"

	"github.com/s9rA16Bf4/Go-tools/tools"
)

// Internal handler that all 'add' functions really depend on
func (handler *ArgumentHandler) internalArgumentHandler(new_argument Argument_t) {
	if handler.utilize_dashes {
		// Check if the argument starts with -- or -
		longResult := tools.Contains(new_argument.long_name, []string{"--"})
		shortResult := tools.Contains(new_argument.short_name, []string{"-"})

		if ok := longResult["--"]; !ok {
			new_argument.long_name = fmt.Sprintf("--%s", new_argument.long_name)
		}

		if ok := shortResult["-"]; !ok {
			new_argument.short_name = fmt.Sprintf("-%s", new_argument.short_name)
		}
	}

	handler.arguments = append(handler.arguments, new_argument)
}
