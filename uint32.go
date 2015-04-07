// Package uint2str contains implementation of Uint32Codec, that can Encode/Decode uint32 by given alphabet
package uint2str

import (
	"bytes"
	"fmt"
)

// Uint32Codec contains alphabet and methods to encode uint32 to string and decode back
type Uint32Codec struct {
	BaseCodec
	alphabetMap map[rune]uint32
	base        uint32
}

// NewUint32Codec returns pointer to uint32 codec instance
func NewUint32Codec(alphabet string) *Uint32Codec {
	sh := new(Uint32Codec)
	sh.init(alphabet)
	return sh
}

func (us *Uint32Codec) init(alphabet string) {
	us.setAlphabet(alphabet)
	us.alphabetMap = map[rune]uint32{}
	for i, rune := range us.alphabet {
		us.alphabetMap[rune] = uint32(i)
	}
	us.base = uint32(len(us.alphabet))
}

// Encode uint32 to string
func (us Uint32Codec) Encode(in uint32) (out string) {
	if in == 0 {
		return string(us.getAlphabet()[0])
	}
	var buffer bytes.Buffer
	for in > 0 {
		buffer.WriteRune(us.getAlphabet()[in%us.base])
		in /= us.base
	}
	return us.reverse(buffer.String())
}

// Decode uint32 from string
func (us Uint32Codec) Decode(in string) (out uint32, err error) {
	for _, character := range in {
		chIdxInAlphabet, exists := us.alphabetMap[character]
		if !exists {
			return 0, fmt.Errorf("character %s does not exist in alphabet", string(character))
		}
		out = out*us.base + chIdxInAlphabet
	}
	return out, nil
}
