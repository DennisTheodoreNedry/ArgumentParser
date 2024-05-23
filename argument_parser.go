package argumentparser

import (
	"github.com/DennisTheodoreNedry/ArgumentParser/lib"
)

// Creates a handler and returns it
// This will be the handler that you will utilize
func Constructor(utilize_dashes bool) *lib.ArgumentHandler {
	var constructed lib.ArgumentHandler
	constructed.SetUtilizeDash(utilize_dashes)

	// Add default argument
	constructed.AddFunction("help", "h", false, false, "Prints a help screen with all arguments", constructed.Help)

	return &constructed
}
