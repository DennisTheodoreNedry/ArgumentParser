package lib

import (
	"fmt"
	"os"
)

// Prints a pretty help screen and exits
func (handler *ArgumentHandler) Help(ignore_value string) string {
	fmt.Println("#### Definied Arguments ####")
	for _, argument := range handler.arguments {
		line := fmt.Sprintf("%s/%s", argument.long_name, argument.short_name)

		// The command needs one or more values to work
		if argument.needs_value {
			line += " <value> "
		}

		// Add the description
		line += " | " + argument.desc

		// Check if the argument has a specific set of options
		if len(argument.options) > 0 {
			line += " | options are {"

			for i, option := range argument.options {
				line += option

				if i+1 < len(argument.options) {
					line += ", "
				}

			}
			line += "}"
		}

		// Is the argument required for the program to worK?
		if argument.required {
			line += " | REQUIRED"
		}

		fmt.Println(line)
	}
	os.Exit(0)
	return ""
}
