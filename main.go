package cmd

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	commands []*Command
)

var (
	// Defaults defines whether we print flag defaults
	Defaults bool = false

	// Desc description for our root command
	Desc string

	// Name of command we are running. defaults to command name executed
	Name string

	// PrefixArgs defines if we allow arguments to show up before flags. If
	// this is allowed arguments (not starting with -) before and after flags
	// will be picked up and given to the command
	PrefixArgs bool
)

func init() {
	Name = os.Args[0]
}

// Register a command with the cmd system
func Register(c *Command) {
	commands = append(commands, c)
}

// splitArgs takes our args and pulls out the arguments from the beginning
// returning us the arguments and the remaining flags
func splitArgs(args []string) ([]string, []string) {
	for idx, arg := range args {
		if arg[0] == '-' {
			return args[:idx], args[idx:]
		}
	}

	// no flags found
	return []string{}, args
}

// Parse loads arguments from the command line and processes them
func Parse() error {
	flag.Usage = printUsage
	flag.Parse()
	log.SetFlags(0)

	args := flag.Args()
	if len(args) < 1 {
		printUsage()
		return nil
	}

	if args[0] == "help" {
		help(args[1:])
		return nil
	}

	for _, cmd := range commands {
		if cmd.Name() == args[0] && cmd.Runnable() {
			cmd.Flag.Usage = func() { cmd.Usage() }

			if cmd.CustomFlags {
				args = args[1:]
			} else if PrefixArgs {
				var flags []string
				args, flags = splitArgs(args[1:])
				cmd.Flag.Parse(flags)
				args = append(args, cmd.Flag.Args()...)
			} else {
				cmd.Flag.Parse(args[1:])
				args = cmd.Flag.Args()
			}
			cmd.Run(cmd, args)
			return nil
		}
	}

	fmt.Fprintf(os.Stderr, "%s: unknown subcommand %q\nRun '%s help' for usage.\n", Name, args[0], Name)
	return errors.New("command not found")
}
