# gopwn 
![Build Status](https://github.com/hupe1980/gopwn/workflows/build/badge.svg) 
[![Go Reference](https://pkg.go.dev/badge/github.com/hupe1980/gopwn.svg)](https://pkg.go.dev/github.com/hupe1980/gopwn)
> Golang CTF framework and exploit development module

This module is strictly for educational purposes only. Usage of the methods and tools for attacking targets without prior mutual consent is illegal. It is the end user's responsibility to obey all applicable laws. Developers assume no liability and are not responsible for any misuse or damage caused by this module.

:warning: This is experimental and subject to breaking changes.

## Usage
```go
package main

import (
  "bytes"
  "fmt"

  "github.com/hupe1980/gopwn"
)

func main() {
  p, _ := gopwn.NewProcess([]string{"./ctfbinary"})
  p.SendLine(append(bytes.Repeat([]byte("A"), 200), gopwn.P32L(0xdeadbeef)...))
  out, _ := p.RecvLine()
  fmt.Println(string(out))
}
```

### Packing Integers
```go
//32Bit LittelEndian
b := gopwn.P32L(0xdeadbeef)
assert.Equal(t, []byte("\xef\xbe\xad\xde"), b) // true
i := gopwn.U32L([]byte("\xef\xbe\xad\xde"))
assert.Equal(t, uint32(0xdeadbeef), i) // true
```

### Assembly and Disassembly
```go
insn, _ := gopwn.AssembleI386("mov eax, 0")
fmt.Println(gopwn.HexString(insn))
```
Outputs:
```
b800000000
```
```go
assembly, _ := gopwn.DisamI386([]byte("\xb8\x5d\x00\x00\x00"), 0)
fmt.Println(assembly)
```
Outputs:
```
0x0           b8 5d 00 00 00                mov eax, 0x5d
```

### Misc Tools
Generate unique sequences to find offsets in your buffer causing a crash:
```go
assert.Equal(t, []byte("aaaabaaacaaadaaa"), gopwn.Cyclic(16)) // true
assert.Equal(t, 4, gopwn.CyclicFind([]byte("baaa")) // true
```

### Binary Analysis and Manipulation
```go
elf, _ := gopwn.NewELF("./ctfbinary")
```

```go
pe, _ := gopwn.NewPE("./ctfbinary.exe")
```

```go
macho, _ := gopwn.NewMACHO("./ctfbinary")
```

### Documentation
See [godoc](https://pkg.go.dev/github.com/hupe1980/gopwn).

### Examples
See more complete [examples](https://github.com/hupe1980/gopwn/tree/main/_examples).

## CLI
```
gopwn command-line interface

Usage:
  gopwn [command]

Available Commands:
  cave        Search for code caves
  checksec    Check binary security settings
  completion  Prints shell autocompletion scripts for gopwn
  cyclic      Generation of unique sequences
  help        Help about any command

Flags:
  -h, --help      help for gopwn
  -v, --version   version for gopwn

Use "gopwn [command] --help" for more information about a command.
```

### Installing
You can install the pre-compiled binary in several different ways

#### deb/rpm/apk:
Download the .deb, .rpm or .apk from the [releases page](https://github.com/hupe1980/gopwn/releases) and install them with the appropriate tools.

#### manually:
Download the pre-compiled binaries from the [releases page](https://github.com/hupe1980/gopwn/releases) and copy to the desired location.

## License
[MIT](LICENCE)
