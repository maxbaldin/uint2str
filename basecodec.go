package uint2str

import "unicode/utf8"

// Default alphabet
const ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// BaseCodec implements common methods for all Shortener codecs
type BaseCodec struct {
	alphabet    []rune
	alphabetMap map[rune]uint32
	base        uint32
}

func (us *BaseCodec) getAlphabet() (characters []rune) {
	return us.alphabet
}

func (us *BaseCodec) setAlphabet(in string) {
	us.alphabet = []rune(in)
	us.alphabetMap = map[rune]uint32{}
	for i, rune := range us.alphabet {
		us.alphabetMap[rune] = uint32(i)
	}
	us.base = uint32(len(us.alphabet))
}

func (us BaseCodec) reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func newBaseCodec(alphabet string) *BaseCodec {
	sh := new(BaseCodec)
	sh.setAlphabet(alphabet)
	return sh
}
