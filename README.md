# Argument Parser

Simpel and effective argument parser, is like your average joe but comes with some nice perks compared to the previous version.

Three methods of adding an argument to your program.
| Function         | Description     | Constraints |
|--------------|-----------|-----------|
| `Add` | Basic function to add arguments | Can't automatically limit the the input with options, can't call a function on input. |
| `AddOptions` | Same as `Add` but requires a string array of possible options to enter  | Can't call a function on input. |
| `AddFunction`| Executes a function when an argument is triggered | The function definition must be `func(string) string`, can't automatically limit the input. |
| `AddFunctionOptions`| Executes a function when an argument is triggered, but can limit the input | The function definition must be `func(string) string`. |

Each argument function must atleast have the following information to work.
| Variable         | Type     | Description |
|--------------|-----------|-----------|
| `long_name` | string | Long hand version of the argument, e.g. `--help` |
| `short_name` | string  | Short hand version of the argument, e.g. `-h` |
| `needs_value`| bool | Does the argument require a value if it's passed in? |
| `required`| bool | Is the argument required for the program to work? |
| `desc`| string | Argument description, e.g. `Sets the verbose level` |


## Help function
There is a built in help function which is printed when `--help`/`-h` or `help/h` is passed.
```
#### Definied Arguments ####
--help/-h | Prints a help screen with all arguments
--test/-t | I'm a test
--test2/-t2 <value>  | I am a required test | options are {1, 2, 3, 4} | REQUIRED
--foo/-fo <value>  | I will peform a function call when entered
```

## To not or to not utilize dashes
You get the option if you would like your arguments to utilize `--`/`-` or not. 
This can be handled in the `Constructor` function by passing in `true/false` depending if you want it or not.

## Return of parsed values
After running `Parse()` a map of the format `map[string]string` is returned, the easiest way to access the value is by utilizing a combination of a for loop and switch case.

```
for key, value := range parsed_result {
		switch key {
		case "test":
			fmt.Println("Argument 'test' was entered")

		case "test2":
			fmt.Printf("Argument 'test2' was entered with value %s\n", value)

		case "foo":
			fmt.Printf("The function argument 'foo' returned %s\n", value)
		}
	}
```

<b>Note:</b> The map returned will contain entries of all entered arguments and will contain both the `long_name` and `short_name` versions of the argument.

## Example
An example of how you can utilize each function can be found under `example/main.go`.
