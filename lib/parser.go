package lib

import "os"

// Parses all entered values and returns a map containing the results
func (handler *ArgumentHandler) Parse() map[string]string {
	parsed_contents := map[string]string{}
	input := os.Args[1:]
	keys := []string{}

	for i := 0; i < len(input); i++ {
		input_argument := input[i]
		// Prints an error and exits if the argument wasn't known
		argument := handler.isAnArgument(input_argument, true)

		valueToArgument := ""

		// Do we need to check a value?
		if argument.needs_value {
			// Do we even have a value?
			if (i + 1) < len(input) {
				valueToArgument = input[i+1]

				// Fast forward
				i++
			}
			handler.matchRequirements(argument, valueToArgument)
		}

		// Is the argument a function?
		if argument.function_call != nil {
			valueToArgument = argument.function_call(valueToArgument)
		}

		// Finalize the parsing
		parsed_contents[argument.long_name] = valueToArgument
		parsed_contents[argument.short_name] = valueToArgument
		keys = append(keys, argument.short_name)
	}

	// Any required commands that did not get called?
	handler.checkRequired(keys)

	return parsed_contents
}
