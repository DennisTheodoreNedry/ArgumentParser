package lib

// Adds an argument that will execute a function when entered
// Any value entered along side the argument will be passed to the function
func (handler *ArgumentHandler) AddFunction(long_name string, short_name string, needs_value bool, required bool, desc string, foo func(string) string) {
	handler.internalArgumentHandler(Argument_t{
		long_name:     long_name,
		short_name:    short_name,
		needs_value:   needs_value,
		required:      required,
		desc:          desc,
		function_call: foo,
	})
}

// Adds an argument that will execute a function when entered
// Any value entered along side the argument will be passed to the function
func (handler *ArgumentHandler) AddFunctionOption(long_name string, short_name string, needs_value bool, required bool, desc string, foo func(string) string, options []string) {
	handler.internalArgumentHandler(Argument_t{
		long_name:     long_name,
		short_name:    short_name,
		needs_value:   needs_value,
		required:      required,
		desc:          desc,
		options:       options,
		function_call: foo,
	})
}
