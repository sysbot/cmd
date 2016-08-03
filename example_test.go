package cmd_test

import (
	"os"

	"zro.io/cmd"
)

var cmdGet = cmd.Command{
	UsageLine: "get [-d] ...",
	Short:     "compile packages and dependencies",
	Long: `
Get downloads go packages...

The -d flag instructs get...
				`,
}

var helpPath = cmd.Command{
	UsageLine: "gopath",
	Short:     "GOPATH environment variable",
	Long: `
The Go path is ...

Lots of help stuff here.
				`,
}

var flagD = cmdGet.Flag.String("d", "", "")

func init() {
	// Define our command name/description
	cmd.Name = "example"
	cmd.Desc = "example command"

	// register our commands
	cmd.Register(&cmdGet)
	cmd.Register(&helpPath)

	// This prevents flag/command compile dependency loop
	cmdGet.Run = runGet
}

func runGet(c *cmd.Command, args []string) {
	// Do Stuff
}

func ExampleAllHelp() {
	os.Args = []string{"example", "help"}
	os.Stderr = os.Stdout

	cmd.Parse()
	// Output:
	// example command
	//
	// Usage:
	//
	// example command [arguments]
	//
	// The commands are:
	//
	// 	get         compile packages and dependencies
	//
	// Use "example help [command]" for more information about a command.
	//
	// Additional help topics:
	//
	// 	gopath      GOPATH environment variable
	//
	// Use "example help [topic]" for more information about that topic.
	//
}

func ExampleSubHelp() {
	os.Args = []string{"example", "help", "get"}
	os.Stderr = os.Stdout

	cmd.Parse()
	// Output:
	// usage: example get [-d] ...
	//
	// Get downloads go packages...
	//
	// The -d flag instructs get...
}

func ExampleExecutePage() {
	os.Args = []string{"example", "gopath"}
	os.Stderr = os.Stdout

	cmd.Parse()
	// Output:
	// example: unknown subcommand "gopath"
	// Run 'example help' for usage.
}
