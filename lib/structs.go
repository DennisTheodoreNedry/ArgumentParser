package lib

type ArgumentHandler struct {
	arguments      []Argument_t
	utilize_dashes bool // Should each argument start with a `--` or `-`
}

type Argument_t struct {
	long_name     string              // Long command for this argument, i.e. '--help'
	short_name    string              // Short command for this argument, i.e. '-h' instead of '--help'
	desc          string              // Description of the command
	needs_value   bool                // Does this command require a value to work?
	required      bool                // Is this command required for the program to work?
	options       []string            // Possible options for this command
	function_call func(string) string // Possible function to call
}
