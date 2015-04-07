package uint2str

import "unicode/utf8"

// Default alphabet
const ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// BaseCodec implements common methods for all Shortener codecs
type BaseCodec struct {
	alphabet []rune
}

func (us *BaseCodec) getAlphabet() (characters []rune) {
	return us.alphabet
}

func (us *BaseCodec) setAlphabet(in string) {
	us.alphabet = []rune(in)
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
