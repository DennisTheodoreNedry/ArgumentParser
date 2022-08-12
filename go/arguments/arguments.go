package arguments

import (
	"fmt"
	"os"

	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

type argument struct {
	longName          string   // Long command for this argument, i.e. '--help'
	shortName         string   // Short command for this argument, i.e. '-h' instead of '--help'
	desc              string   // Description of the command
	argument_required bool     // Does this command require a value to function?
	argument_value    string   // This will contain the passed value if something was passed [deprecated]
	set               bool     // Has this argument been passed? [deprecated]
	options           []string // Available options for this command
}

var defined_arguments []argument // This array will contain all the defined arguments

/*
*  Adds an argument but without any optional values for said argument
**/
func Argument_add(longName string, shortName string, argument_required bool, desc string) {
	defined_arguments = append(defined_arguments, argument{longName, shortName, desc, argument_required, "NULL", false, []string{"NULL"}})
}

/*
*  Adds an argument but with the possibility to use define optional values for said argument
**/
func Argument_add_with_options(longName string, shortName string, argument_required bool, desc string, options []string) {
	defined_arguments = append(defined_arguments, argument{longName, shortName, desc, argument_required, "NULL", false, options})
}

/*
*  Parses all entered arguments, check if they contain the correct value and does all magic
**/
func Argument_parse() map[string]string {

	toReturn := make(map[string]string) // Only the flags entered will be found here

	for i := 1; i <= len(os.Args[1:]); i++ {
		found := false
		i_arg := os.Args[i] // The input argument

		for j, def_arg := range defined_arguments {
			if i_arg == def_arg.longName || i_arg == def_arg.shortName {

				if def_arg.argument_required {
					if i+1 > len(os.Args[1:]) {
						notify.Error(fmt.Sprintf("The argument %s needs an argument to work!", "arguments.argument_parse()", i_arg), "arguments.argument_parse()")
					}

					possible_options := false
					i_value := os.Args[i+1] // The input value

					for _, opt := range def_arg.options {
						if opt == i_value {
							defined_arguments[j].argument_value = i_value // Save the value (deprecated)
							// I believe the j variable becomes totally deprecated with the new way of handling

							toReturn[i_arg] = i_value

							i += 1 // So that we skip one row
							possible_options = true
							break
						}
					}

					if !possible_options {
						notify_error_msg := fmt.Sprintf("Unknown option %s to argument %s, possible options are [", i_value, i_arg)
						for i, option := range def_arg.options {
							notify_error_msg += option
							if i+1 < len(def_arg.options) {
								notify_error_msg += ","
							}
						}
						notify_error_msg += "]"

						notify.Error(notify_error_msg, "arguments.argument_parse()")
					}

				} else { // The argument was called but it doesn't require anything to work
					toReturn[i_arg] = ""
				}

				defined_arguments[j].set = true // (deprecated)
				found = true                    // Indicate that we found the flag
				break
			}
		}

		if !found {
			if i_arg == "-h" || i_arg == "--help" {
				Argument_help()
			} else {
				notify.Error(fmt.Sprintf("The argument %s was not defined!", i_arg), "arguments.argument_parse()")
			}
		}

	}

	return toReturn
}

// ----  DEPRECATED FUNCTIONS START ---- //
func Argument_check(arg_name string) bool { // Checks if a argument has been called
	notify.Warning("This function is deprecated and instead you should utilize the map being returned when parsing")
	var toReturn bool = false
	for _, arg := range defined_arguments {
		if arg.longName == arg_name || arg.shortName == arg_name {
			toReturn = arg.set
			break
		}
	}
	return toReturn
}

func Argument_get(arg_name string) string { // Gets the provided value for a set argument, and will return NULL otherwise
	notify.Warning("This function is deprecated and instead you should utilize the map being returned when parsing")
	var toReturn string = "NULL"
	for _, arg := range defined_arguments {
		if (arg.longName == arg_name || arg.shortName == arg_name) && arg.set {
			toReturn = arg.argument_value
			break
		}
	}
	return toReturn
}

// ----  DEPRECATED FUNCTIONS END ---- //

/*
*  Prints all defined arguments
**/
func Argument_help() {
	fmt.Println("#### Definied Arguments ####")
	for id := 0; id < len(defined_arguments); id++ {
		value := defined_arguments[id].longName + ", " + defined_arguments[id].shortName

		if defined_arguments[id].argument_required {
			value += " <value> "
		}
		value += " | " + defined_arguments[id].desc

		if defined_arguments[id].options[0] != "NULL" { // There were some predefined options defined
			options := " | options are ["
			for i, opt := range defined_arguments[id].options {
				options += opt
				if i+1 < len(defined_arguments[id].options) {
					options += ", "
				}
			}

			options += "]"
			value += options
		}

		fmt.Println(value)
	}

	os.Exit(0)
}
