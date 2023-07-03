package main

import (
	argumentparser "github.com/s9rA16Bf4/ArgumentParser"
)

func main() {
	handler := argumentparser.Constructor()
	handler.Add("test", "t", false, true, "I'm a test")
	handler.AddOptions("test2", "t2", true, true, "I am also a test", []string{"1", "2", "3", "4"})

	handler.Parse()

}
