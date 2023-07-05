package lib

// Basic way of adding an argument
func (handler *ArgumentHandler) Add(long_name string, short_name string, needs_value bool, required bool, desc string) {
	handler.internalArgumentHandler(Argument_t{
		long_name:   long_name,
		short_name:  short_name,
		needs_value: needs_value,
		required:    required,
		desc:        desc,
	})
}
