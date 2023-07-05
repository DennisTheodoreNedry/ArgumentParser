package lib

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Check if the entered argument has everything it needs to work
// It works by checking for errors and reports them
func (handler *ArgumentHandler) matchRequirements(argument *Argument_t, value string) {

	if argument.needs_value && value == "" {
		notify.Error(fmt.Sprintf("Argument %s/%s needs a value to work, yet none was given", argument.long_name, argument.short_name), "argumentparser.matchRequirements()")
	}

	// Check if the argument has a specific set of options that needs to be met
	match := false

	for _, option := range argument.options {
		if option == value {
			match = true
		}
	}

	if len(argument.options) > 0 && !match {
		notify.Error(fmt.Sprintf("The value passed to %s/%s can only match %v", argument.long_name, argument.short_name, argument.options), "argumentparser.matchRequirements()")
	}

}
