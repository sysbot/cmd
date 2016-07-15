package cmd_test

import "zro.net/go/cmd"

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
	// This prevents flag/command compile dependency loop
	cmdGet.Run = runGet

	// register our commands
	cmd.Register(&cmdGet)
	cmd.Register(&helpPath)
}

func runGet(c *cmd.Command, args []string) {
	// Do Stuff
}

func Example() {
	cmd.Parse()
	// Calling cmd.Parse() from main() processes/executes your command
	// It will return though there wont be anything left for you to do
	// as far as executing your command goes
}
