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
clihelp.Print(
    clihelp.F("--flag1", "bool", "Explanation of flag1"),
    clihelp.F("--flag2", "string", "Explanation of flag2"),
    clihelp.F("--flag3", "", "Explanation of flag3"),
    clihelp.F("--flag4", "int", "Explanation of flag4"),
)
```
### Output
```bash
Options:
  --flag1    bool      Explanation of flag1
  --flag2    string    Explanation of flag2
  --flag3              Explanation of flag3
  --flag4    int       Explanation of flag4
```
### Design

Formatting only

No flag parsing

No opinionated layout

Intended to be embedded into custom help output


That’s it.  

---

