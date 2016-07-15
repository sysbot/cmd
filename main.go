package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	commands []*Command
)

var (
	// Desc description for our root command
	Desc string

	// Name of command we are running. defaults to command name executed
	Name string
)

func init() {
	Name = os.Args[0]
}

// Register a command with the cmd system
func Register(c *Command) {
	commands = append(commands, c)
}

// Parse loads arguments from the command line and processes them
func Parse() {
	flag.Usage = printUsage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	if len(args) < 1 {
		printUsage()
	}

	if args[0] == "help" {
		help(args[1:])
		return
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Runnable() {
			cmd.Flag.Usage = func() { cmd.Usage() }
			if cmd.CustomFlags {
				args = args[1:]
			} else {
				cmd.Flag.Parse(args[1:])
				args = cmd.Flag.Args()
			}
			cmd.Run(cmd, args)
			return
		}
	}

	fmt.Fprintf(os.Stderr, "%s: unknown subcommand %q\nRun '%s help' for usage.\n", Name, args[0], Name)
	os.Exit(2)
}
