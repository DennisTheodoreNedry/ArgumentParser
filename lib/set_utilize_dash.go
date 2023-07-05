package lib

// Enables/Disables the option to utilize '--'/'-' in your arguments
func (handler *ArgumentHandler) SetUtilizeDash(status bool) {
	handler.utilize_dashes = status
}
