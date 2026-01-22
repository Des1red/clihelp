# clihelp

Minimal helper for rendering aligned CLI help output in Go.

## Installation

```bash
go get github.com/Des1red/clihelp
```


### Usage
```go
import "github.com/Des1red/clihelp"

fmt.Println("Options:")
clihelp.PrintFlags([]clihelp.Flag{
    {"--flag1", "bool", "Explenation of flag1"},
    {"--flag2", "string", "Explenation of flag2"},
    {"--flag3", "", "Explenation of flag3"},
    {"--flag4", "int", "Explenation of flag4"},
})
```
### Output
```bash
Options:
  --flag1    bool     Explenation of flag1
  --flag2    string   Explenation of flag2
  --flag3             Explenation of flag3
  --flag4    int      Explenation of flag4
```
### Design

Formatting only

No flag parsing

No opinionated layout

Intended to be embedded into custom help output


That’s it.  

---

