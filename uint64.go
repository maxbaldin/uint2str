package uint2str

import (
	"bytes"
	"fmt"
)

// Uint64Codec contains alphabet and methods to encode uint32 to string and decode back
type Uint64Codec struct {
	BaseCodec
	alphabetMap map[rune]uint64
	base        uint64
}

// NewUint64Codec returns pointer to uint64 codec instance
func NewUint64Codec(alphabet string) *Uint64Codec {
	sh := new(Uint64Codec)
	sh.init(alphabet)
	return sh
}

func (us *Uint64Codec) init(alphabet string) {
	us.setAlphabet(alphabet)
	us.alphabetMap = map[rune]uint64{}
	for i, rune := range us.alphabet {
		us.alphabetMap[rune] = uint64(i)
	}
	us.base = uint64(len(us.alphabet))
}

// Encode uint64 to string
func (us Uint64Codec) Encode(in uint64) (out string) {
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

// Decode uint64 from string
func (us Uint64Codec) Decode(in string) (out uint64, err error) {
	for _, character := range in {
		chIdxInAlphabet, exists := us.alphabetMap[character]
		if !exists {
			return 0, fmt.Errorf("character %s does not exist in alphabet", string(character))
		}
		out = out*us.base + chIdxInAlphabet
	}
	return out, nil
}
