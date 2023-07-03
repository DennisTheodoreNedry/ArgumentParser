# ArgumentParser

Compared to the cpp version of this project, the golang version is more advanced as you now can set expected values, and raise an error when those are not met.

The argument parser consists out of 5 functions which are easy to use, 


### Example
The following is how a `main` function could look like
```
package main

import (
	arg "github.com/s9rA16Bf4/ArgumentParser/go/arguments"
	patcher "github.com/s9rA16Bf4/No_CD_Cracks/Jurassic_Park_Trespasser/utility/patcher"
	"github.com/s9rA16Bf4/notify_handler/go/notify"
)

func main() {
	arg.Argument_add("--exe", "-x", true, "Path to Trespasser exe [REQUIRED]")
	arg.Argument_add("--smks", "-s", true, "Path to a folder containing the four smk's [REQUIRED]")
	arg.Argument_add("--levels", "-l", true, "Path to all the different levels and other materials [REQUIRED]")
	parsed_flags := arg.Argument_parse()

	if len(parsed_flags) > 0 {
		var path_to_exe string
		var path_to_smks string
		var path_to_lvl string

		if value, entered := parsed_flags["-x"]; entered { // Patch exe
			path_to_exe = value
		} else {
			notify.Error("No exe file was provided", "main.main()")
		}

		if value, entered := parsed_flags["-s"]; entered { // Patch exe
			path_to_smks = value
		} else {
			notify.Error("No folder was provided", "main.main()")
		}

		if value, entered := parsed_flags["-l"]; entered {
			path_to_lvl = value
		} else {
			notify.Error("Levels were not provided", "main.main()")
		}

		patcher.Begin_patch(path_to_exe, path_to_smks, path_to_lvl)

	} else {
		notify.Error("No argument was provided, run '--help'/'-h' to have a look at the arguments available", "main.main()")
	}
}

```
Result
```
./trespasser_patcher.exe -h
#### Definied Arguments ####
--exe, -x <value>  | Path to Trespasser exe [REQUIRED]
--smks, -s <value>  | Path to a folder containing the four smk's [REQUIRED]
--levels, -l <value>  | Path to all the different levels and other materials [REQUIRED]

./trespasser_patcher.exe
#### Error ####
msg: No exe file was provided
where: main.main()

```
