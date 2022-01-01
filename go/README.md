# ArgumentParser

Compared to the cpp version of this project, the golang version is more advanced as you now can set expected values, and raise an error when those are not met.

The argument parser consists out of 5 functions which are easy to use, 


### Example
The following is how a `main` function could look like
```
package main

import (
	"fmt"
	"os"

	arg "github.com/s9rA16Bf4/ArgumentParser/go/arguments"
	"github.com/s9rA16Bf4/go-evil/utility/notify"
)

func main() {

	arg.Argument_add("--help", "-h", false, "Shows all available arguments and their purpose", []string{"NULL"})
	arg.Argument_add("--verbose", "-v", true, "How verbose should the program be, options are [1,2,3]", []string{"0", "1", "2", "3"})
	arg.Argument_add("--debug", "-d", true, "Debug iptions, options are [false, true]", []string{"false", "true"})
	arg.Argument_add("--echo", "-e", true, "Echos the provided value in the console", []string{"NULL"})

	arg.Argument_parse() // Lets check what the user entered

	if len(os.Args[0:]) > 1 { // The user entered something
		if arg.Argument_check("-h") {
			arg.Argument_help()
		} else {
			if arg.Argument_check("-v") {
				notify.Verbose_lvl = arg.Argument_get("-v")
				notify.Notify_log("Setting verbose level to "+notify.Verbose_lvl, notify.Verbose_lvl, "1")
			}

			if arg.Argument_check("-d") {
				if arg.Argument_get("-d") == "true" {
					// set debug true here
				}
			}
			if arg.Argument_check("-e") {
				fmt.Println(arg.Argument_get("-e"))
			}
		}
	} else {
		notify.Notify_error("No argument was provided, run '--help'/'-h' to have a look at the arguments available", "main.main()")
	}
}

```
Result
```
./test --help
#### Definied Arguments ####
--help, -h | Shows all available arguments and their purpose
--verbose, -v <value>  | How verbose should the program be, options are [1,2,3]
--debug, -d <value>  | Debug iptions, options are [false, true]
--echo, -e <value>  | Echos the provided value in the console

./test
#### Error ####
msg: No argument was provided, run '--help'/'-h' to have a look at the arguments available
where: main.main()

./test --echo "Hello world"
Hello world

./test -e "Hello world"
Hello world
```
