package argumentparser

import (
	"fmt"
	"os"

	"github.com/s9rA16Bf4/Go-tools/tools"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type ArgumentHandler struct {
	arguments []Argument_t
}

type Argument_t struct {
	longName       string   // Long command for this argument, i.e. '--help'
	shortName      string   // Short command for this argument, i.e. '-h' instead of '--help'
	desc           string   // Description of the command
	needs_value    bool     // Does this command require a value to work?
	needs_multiple bool     // Does the command need multiple values to work?
	required       bool     // Is this command required for the program to work?
	options        []string // Possible options for this command
}

// Creates a handler and returns it
// This will be the handler that you will utilize
func Constructor() *ArgumentHandler {
	var constructed ArgumentHandler

	// Add default value
	constructed.Add("--help", "-h", false, false, "Prints the help screen")

	return &constructed
}

// Basic way of adding an argument
func (self *ArgumentHandler) Add(longName string, shortName string, needs_value bool, needs_multiple bool, required bool, desc string) {
	self.AddOptions(longName, shortName, needs_value, needs_multiple, required, desc, []string{})
}

// Adds an argument but with options
func (self *ArgumentHandler) AddOptions(longName string, shortName string, needs_value bool, needs_multiple bool, required bool, desc string, options []string) {

	// Check if the argument starts with -- or -
	longResult := tools.Contains(longName, []string{"--"})
	shortResult := tools.Contains(shortName, []string{"-"})

	if ok := longResult["--"]; !ok {
		longName = fmt.Sprintf("--%s", longName)
	}

	if ok := shortResult["-"]; !ok {
		shortName = fmt.Sprintf("-%s", shortName)
	}

	self.arguments = append(self.arguments, Argument_t{longName, shortName, desc, needs_value, needs_multiple, required, options})
}

// Check if the entered argument has everything it needs to work
// It primarly works by checking for errors
func (self *ArgumentHandler) matchRequirements(argument *Argument_t, values []string) {

	if argument.needs_value && len(values) == 0 {
		notify.Error(fmt.Sprintf("Argument %s/%s needs a value to work, yet none was given", argument.longName, argument.shortName), "argumentparser.matchRequirements()")
	}

	if !argument.needs_value && len(values) != 0 {
		notify.Error(fmt.Sprintf("Argument %s/%s needs no values to work, yet %d was given", argument.longName, argument.shortName, len(values)), "argumentparser.matchRequirements()")
	}

	// Check if the argument has a specific set of options that needs to be met
	for _, option := range argument.options {
		for _, value := range values {
			if option != value {
				notify.Error(fmt.Sprintf("Argument %s/%s has a specific set of options, %s did not match it", argument.longName, argument.shortName, value), "argumentparser.matchRequirements()")
			}
		}
	}

}

// Checks if the entered value is an argument
func (self *ArgumentHandler) isAnArgument(value string, raiseOnError bool) *Argument_t {
	var match *Argument_t = nil

	for _, defined_argument := range self.arguments {
		result := tools.Contains(value, []string{defined_argument.longName, defined_argument.shortName})

		longOK := result[defined_argument.longName]
		shortOK := result[defined_argument.shortName]

		if longOK || shortOK {
			match = &defined_argument
			break
		}
	}

	// Unknown command, print and exit
	if raiseOnError && match != nil {
		notify.Error(fmt.Sprintf("Found unknown argument '%s'", value), "argumentparser.isAnArgument()")
	}

	return match
}

// Checks if any required argument was left out aka not called
func (self *ArgumentHandler) checkRequired(keys []string) {
	for _, argument := range self.arguments {

		// This command is not required to work
		if !argument.required {
			continue
		}

		match := false
		for _, key := range keys {
			if argument.shortName == key {
				match = true
			}
		}

		if !match {
			notify.Error(fmt.Sprintf("Argument %s/%s is required but was not called", argument.longName, argument.shortName), "argumentparser.checkRequired()")
		}

	}
}

// Parses all entered values and returns a map containing the results
func (self *ArgumentHandler) Parse() map[string][]string {
	parsed_contents := map[string][]string{}
	keys := []string{}
	input := os.Args[1:]

	for i, input_argument := range input {
		// Prints an error and exits if the argument wasn't known
		argument := self.isAnArgument(input_argument, true)
		valueToArgument := []string{}

		// Is the argument help?
		if argument.longName == "--help" {
			self.help()
		}

		// Found the command, so now we need to check if we match it
		// For this to work we need to fast forward and obtain everything either until EOF or until the next command has been encountered
		line := input[i+1]

		for ok := self.isAnArgument(line, false); ok == nil; {
			// Reached EOF
			if (i + 1) >= len(input) {
				break
			}

			// Add the value
			valueToArgument = append(valueToArgument, line)

			// Grab the next value
			i++
			line = input[i]
		}

		// Does the passed values work with the argument?
		self.matchRequirements(argument, valueToArgument)

		// Finalize the parsing
		parsed_contents[argument.longName] = valueToArgument
		parsed_contents[argument.shortName] = valueToArgument
		keys = append(keys, argument.shortName)

	}

	// Any required commands that did not get called?

	return parsed_contents
}

// Prints a pretty help screen and exits
func (self *ArgumentHandler) help() {
	fmt.Println("#### Definied Arguments ####")
	for _, argument := range self.arguments {
		line := fmt.Sprintf("%s/%s", argument.longName, argument.shortName)

		// The command needs one or more values to work
		if argument.needs_value {
			line += " <value> "

			if argument.needs_multiple {
				line += " ... "
			}
		}

		// Add the description
		line += " | " + argument.desc

		// Check if the argument has a specific set of options
		if len(argument.options) > 0 {
			line += " | options are {"

			for i, option := range argument.options {
				line += option

				if i+1 < len(argument.options) {
					line += ", "
				}

			}
			line += "}"
		}

		fmt.Println(line)
	}
}
