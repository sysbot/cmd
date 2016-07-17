package cmd_test

import (
	"os"

	"zro.net/go/cmd"
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
	// Only for test to pass
	os.Args = []string{"example", "help"}
	os.Stderr = os.Stdout

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

func Example() {
	// cmd.Parse() executes your command though also returns.
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
