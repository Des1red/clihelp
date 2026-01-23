// Package clihelp provides a minimal helper for rendering
// aligned CLI flag descriptions.
//
// It is intended to be used by applications that implement
// their own help or usage output, while keeping formatting
// consistent and readable.
//
// Example usage:
//
//	fmt.Println("Flags:")
//	clihelp.Print(
//	    clihelp.F("--flag1", "value", "Explanation"),
//	    clihelp.F("--flag2", "value", "Explanation"),
//		clihelp.F("--flag3", "", "Explanation"),
//	)
package clihelp

import "fmt"

// Flag describes a single CLI flag for help output.
type Flag struct {
	Name string // flag name, e.g. --listen
	Arg  string // argument name, e.g. address
	Desc string // description
}

// F creates a Flag with the provided name, argument, and description.
func F(name, arg, desc string) Flag {
	return Flag{
		Name: name,
		Arg:  arg,
		Desc: desc,
	}
}

// Print prints aligned flag descriptions to stdout.
func Print(flags ...Flag) {
	flagW, argW := calcWidths(flags)
	for _, f := range flags {
		fmt.Printf(
			"  %-*s  %-*s  %s\n",
			flagW,
			f.Name,
			argW,
			f.Arg,
			f.Desc,
		)
	}
}

// calcWidths calculates the max width for flag and argument columns.
func calcWidths(flags []Flag) (flagW, argW int) {
	for _, f := range flags {
		if len(f.Name) > flagW {
			flagW = len(f.Name)
		}
		if len(f.Arg) > argW {
			argW = len(f.Arg)
		}
	}
	return
}
