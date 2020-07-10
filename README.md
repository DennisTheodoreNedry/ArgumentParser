# ArgumentParser
A simpel argument parser for C++

The argument parser consists out of 5 functions which are easy to use, take a look at the example below to get a good look at it.


### Example
The following is how a `main` function could look like
```
#include "args.hpp"

int main(int argc, char** argv){
  arg _arg;
  _arg.addArgument("--help", false, "Prints the help screen"); // This tells us that the flag to trigger this event is '--help' and that it doesn't need an argument to work
  _arg.addArgument("--echo", "-e", true, "Prints a message to the screen"); // Tells us that we can use '--echo' and '-e' but it does require an argument to work

  _arg.parseArgs(argc, argv); // This takes a look at what the user has entered, and sets the correct flags to there respectfully values
  if (argc > 1){
    if (_arg.checkArg("--help")){ // Checks if the '--help' flag has been triggered, and returns true if that se case
      _arg.help(); // This prints the help screen which shows all the arguments, their different formats and if they require an argument to work
    }else if (_arg.checkArg("--echo")){ // Did we enter this argument? Its totally okay if we entered '-e' instead
        std::cout << _arg.getValue("--echo") << std::endl; // What ever the cause, lets print the value
    }
  }else{ _arg.help(); return 1; }

  return 0;
}

```
Result
```
./test --help
#### Definied Arguments ####
--help  | Prints the help screen
--echo, -e <value> | Prints a message to the screen

./test
#### Definied Arguments ####
--help  | Prints the help screen
--echo, -e <value> | Prints a message to the screen

./test --echo "Hello world"
Hello world

./test -e "Hello world"
Hello world
```
