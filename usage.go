// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package cmd

import (
	"bufio"
	"fmt"
	"html/template"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

var usageTemplate = `{{.Desc}}

Usage:

{{.Name}} command [arguments]

The commands are:
{{range .Cmds}}{{if .Runnable}}
	{{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use "{{.Name}} help [command]" for more information about a command.

Additional help topics:
{{range .Cmds}}{{if not .Runnable}}
	{{.Name | printf "%-11s"}} {{.Short}}{{end}}{{end}}

Use "{{.Name}} help [topic]" for more information about that topic.

`

var helpTemplate = `{{if .Runnable}}usage: go {{.UsageLine}}

{{end}}{{.Long | trim}}
`

// An errWriter wraps a writer, recording whether a write error occurred.
type errWriter struct {
	w   io.Writer
	err error
}

func (w *errWriter) Write(b []byte) (int, error) {
	n, err := w.w.Write(b)
	if err != nil {
		w.err = err
	}
	return n, err
}

func errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
	os.Exit(1)
}

// tmpl executes the given template text on data, writing the result to w.
func tmpl(w io.Writer, text string, data interface{}) {
	t := template.New("top")
	t.Funcs(template.FuncMap{"trim": strings.TrimSpace, "capitalize": capitalize})
	template.Must(t.Parse(text))
	ew := &errWriter{w: w}
	err := t.Execute(ew, data)
	if ew.err != nil {
		// I/O error writing. Ignore write on closed pipe.
		if strings.Contains(ew.err.Error(), "pipe") {
			os.Exit(1)
		}
		errorf("writing output: %v", ew.err)
	}
	if err != nil {
		panic(err)
	}
}

func capitalize(s string) string {
	if s == "" {
		return s
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToTitle(r)) + s[n:]
}

func printUsage() {
	bw := bufio.NewWriter(os.Stderr)
	usageData := struct {
		Name string
		Desc string
		Cmds []*Command
	}{Name, Desc, commands}

	tmpl(bw, usageTemplate, usageData)
	bw.Flush()
	os.Exit(2)
}

// help implements the 'help' command.
func help(args []string) {
	if len(args) == 0 {
		printUsage()
		// not exit 2: succeeded at 'go help'.
		return
	}
	if len(args) != 1 {
		fmt.Fprintf(os.Stderr, "usage: go help command\n\nToo many arguments given.\n")
		os.Exit(2) // failed at 'go help'
	}

	arg := args[0]

	for _, cmd := range commands {
		if cmd.Name() == arg {
			tmpl(os.Stdout, helpTemplate, cmd)
			// not exit 2: succeeded at 'go help cmd'.
			return
		}
	}

	fmt.Fprintf(os.Stderr, "Unknown help topic %#q.  Run 'go help'.\n", arg)
	os.Exit(2) // failed at 'go help cmd'
}
