# uint2str

Package provides ability to convert uint32 to string by given alphabet(use default as in example, or setup your own).
Useful when you need possibility to convert number to string and back(url shortener for example).

# Advantages

* Tiny
* Fast
* UTF-8 is fully supported
* Well tested([![GoCover](http://gocover.io/_badge/github.com/MaxBaldin/uint2str)](http://gocover.io/github.com/MaxBaldin/uint2str)) and benchmarked
* Has no 3rd party dependencies

# Installation

```
go get github.com/maxbaldin/uint2str
```

# Usage

Basic usage:
```go
package main

import (
    "log"
    "github.com/maxbaldin/uint2str"
)

func main() {
    codec := uint2str.NewUint32Codec(uint2str.ALPHABET)
    encoded := codec.Encode(1234)
    log.Println(encoded)
    decoded, err := codec.Decode(encoded)
    if err != nil {
        log.Fatalf("Unable decode integer from string `%s`. Error: %s", encoded, err.Error())
    }
    log.Println(decoded)
}
```

Output for this example:

```
2015/04/06 23:11:05 t4
2015/04/06 23:11:05 1234
```

# Documentation

[![GoDoc](https://godoc.org/github.com/MaxBaldin/uint2str?status.svg)](https://godoc.org/github.com/MaxBaldin/uint2str)

# Tests and performance

Run tests:
```
go test github.com/maxbaldin/uint2str
```

Run benchmarks:
```
go test -bench github.com/maxbaldin/uint2str
```

# License

BSD 3
