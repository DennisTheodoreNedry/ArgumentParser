package main

import (
	"fmt"

	argumentparser "github.com/s9rA16Bf4/ArgumentParser"
)

func foo(value string) string {
	fmt.Printf("I got the value %s and will return 'done'\n", value)
	return "done"
}

func foo2(value string) string {
	fmt.Println("I can limit options but also act as a function call")
	return "done"
}

func main() {
	// True sent to the constructor will make sure that each command is prefixed with a `--` or `-`
	// depending if it's a long or short argument
	handler := argumentparser.Constructor(true)

	handler.Add("test", "t", false, false, "I'm a test")
	handler.AddOptions("test2", "t2", true, true, "I am a required test", []string{"1", "2", "3", "4"})
	handler.AddFunction("foo", "fo", true, false, "I will peform a function call when entered", foo)
	handler.AddFunctionOptions("foo2", "fo2", true, false, "I will peform a function call when entered", foo2, []string{"1", "2", "3", "4"})

	// The function arg foo/fo will be executed during handler.Parse() if it's entered
	parsed_result := handler.Parse()

	// You can after parsing the contents simply combine a for loop and switch case
	// To see what was entered and handle each case
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
}
