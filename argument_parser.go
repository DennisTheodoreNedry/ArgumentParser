package argumentparser

import "fmt"

type ArgumentHandler struct {
	arguments []Argument_t
}

type Argument_t struct {
	longName    string   // Long command for this argument, i.e. '--help'
	shortName   string   // Short command for this argument, i.e. '-h' instead of '--help'
	desc        string   // Description of the command
	needs_value bool     // Does this command require a value to work?
	required    bool     // Is this command required for the program to work?
	options     []string // Possible options for this command
}

// Creates a handler and returns it
// This will be the handler that you will utilize
func Constructor() *ArgumentHandler {
	var constructed ArgumentHandler
	return &constructed
}

// Basic way of adding an argument
func (self *ArgumentHandler) Add(longName string, shortName string, needs_value bool, required bool, desc string) {
	self.AddOptions(longName, shortName, needs_value, required, desc, []string{})
}

// Adds an argument but with options
func (self *ArgumentHandler) AddOptions(longName string, shortName string, needs_value bool, required bool, desc string, options []string) {
	self.arguments = append(self.arguments, Argument_t{longName, shortName, desc, needs_value, required, options})
}

// Parses all entered values and returns a map containing the results
func (self *ArgumentHandler) Parse() {
	for _, argument := range self.arguments {
		fmt.Println(argument.longName)
	}
}

// Prints a pretty help screen and exits
func (self *ArgumentHandler) Help() {

}
