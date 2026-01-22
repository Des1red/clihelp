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
    {"--listen", "address", "Listen address"},
    {"--mode", "string", "direct | tor"},
})
```
### Output
```bash
Options:
  --listen  address  Listen address
  --mode    string   direct | tor
```
### Design

Formatting only

No flag parsing

No opinionated layout

Intended to be embedded into custom help output


That’s it.  

---

