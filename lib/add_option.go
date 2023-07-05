package lib

// Adds an argument but with options
func (handler *ArgumentHandler) AddOptions(long_name string, short_name string, needs_value bool, required bool, desc string, options []string) {
	handler.internalArgumentHandler(Argument_t{
		long_name:   long_name,
		short_name:  short_name,
		needs_value: needs_value,
		required:    required,
		desc:        desc,
		options:     options,
	})
}
