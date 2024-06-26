package lib

import (
	"fmt"

	notify "github.com/DennisTheodoreNedry/notify_handler"
)

// Checks if the entered value is an argument
func (handler *ArgumentHandler) isAnArgument(value string, raiseOnError bool) *Argument_t {
	var match *Argument_t = nil

	for _, defined_argument := range handler.arguments {

		if value == defined_argument.long_name || value == defined_argument.short_name {
			match = &defined_argument
			break
		}
	}

	// Unknown command, print and exit
	if raiseOnError && match == nil {
		notify.Error(fmt.Sprintf("Found unknown argument '%s'", value), "argumentparser.isAnArgument()", 1)
	}

	return match
}
