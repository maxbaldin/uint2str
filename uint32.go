// Package uint2str contains implementation of Uint32Codec, that can Encode/Decode uint32 by given alphabet
package uint2str

import (
    "bytes"
    "fmt"
)

// Uint32Codec is basic implementation of Codec interface
type Uint32Codec struct {
    BaseCodec
}

// NewUint32Codec returns pointer to uint32 codec instance
func NewUint32Codec(alphabet string) *Uint32Codec {
    sh := new(Uint32Codec)
    sh.setAlphabet(alphabet)
    return sh
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
