package lib

import (
	"fmt"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

// Checks if any required argument was left out aka not called
func (handler *ArgumentHandler) checkRequired(keys []string) {
	for _, argument := range handler.arguments {

		// This command is not required to work
		if !argument.required {
			continue
		}

		match := false
		for _, key := range keys {
			if argument.short_name == key {
				match = true
			}
		}

		if !match {
			notify.Error(fmt.Sprintf("Argument %s/%s is required but was not called", argument.long_name, argument.short_name), "argumentparser.checkRequired()")
		}

	}
}
