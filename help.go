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
//	clihelp.PrintFlags([]clihelp.Flag{
//	    {"--listen", "address", "Listen address"},
//	    {"--mode", "string", "Execution mode"},
//	})
package clihelp

import "fmt"

// Flag describes a single CLI flag for help output.
type Flag struct {
	Name string // flag name, e.g. --listen
	Arg  string // argument name, e.g. address
	Desc string // description
}

// PrintFlags prints aligned flag descriptions to stdout.
func PrintFlags(flags []Flag) {
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
