# uint2str

Package provides ability to convert number to string by given alphabet(use default as in example, or setup your own).
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

## Basic

### Encode

```go
package main

import (
	"github.com/maxbaldin/uint2str"
	"log"
)

func main() {
	c := uint2str.NewCodec(uint2str.ALPHABET)

	number := 100
	encoded, err := c.Encode(number)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d -> %s", number, encoded)
}
```

Output for this example:

```
100 -> bM
```

### Decode

```go
package main

import (
	"github.com/maxbaldin/uint2str"
	"log"
)

func main() {
	c := uint2str.NewCodec(uint2str.ALPHABET)

	str := "bM"
	var number int
	err := c.Decode(str, &number)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s -> %d", str, number)
}

```

Output for this example:

```
bM -> 100
```

### Supported types

Both encode and decode: int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64

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
