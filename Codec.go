package uint2str

import (
	"fmt"
	"reflect"
)

// Sample of alphabet for encode/decode
const ALPHABET = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Codec holds alphabet and methods for encode/decode numbers to/from string
type Codec struct {
	alphabet      string
	alphabetSlice []rune

	alphabetMapInt64 map[rune]int64
	baseInt64        int64

	alphabetMapUint64 map[rune]uint64
	baseUint64        uint64
}

// NewCodec returns a Codec, that can encode numeric types to string and decode it back by given alphabet
func NewCodec(alphabet string) *Codec {
	c := new(Codec)
	c.setAlphabet(alphabet)
	return c
}

// GetAlphabet returns alphabet
func (c Codec) GetAlphabet() string {
	return c.alphabet
}

// Encode integer types(int,int64,uint8 for example) to string
func (c Codec) Encode(v interface{}) (encoded string, err error) {
	val := reflect.ValueOf(v)
	if !val.IsValid() {
		return "", fmt.Errorf("v is invalid")
	}
	val = getValThroughPointers(val)
	switch val.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return c.encodeInt64(val.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return c.encodeUint64(val.Uint()), nil
	default:
		return "", fmt.Errorf("unsupported type %s", val.Type().String())
	}
}

// Decode string to given type
func (c Codec) Decode(str string, v interface{}) (err error) {
	if str == "" {
		return fmt.Errorf("str is empty")

	}
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return fmt.Errorf("v is invalid")
	}
	if rv.Kind() != reflect.Ptr {
		return fmt.Errorf("non pointer %s", rv.Type().String())
	}
	in := getValThroughPointers(rv)
	switch in.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		decoded, err := c.decodeInt64(str)
		if err != nil {
			return err
		}
		in.SetInt(decoded)
		break
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		decoded, err := c.decodeUint64(str)
		if err != nil {
			return err
		}
		in.SetUint(decoded)
		break
	default:
		return fmt.Errorf("unsupported type %s", rv.Type().String())
	}
	return
}

func getValThroughPointers(v reflect.Value) (value reflect.Value) {
	if v.Kind() != reflect.Ptr {
		return v
	}
	value = v
	for {
		if value.Kind() == reflect.Ptr && value.IsValid() {
			value = value.Elem()
		} else {
			break
		}
	}
	return
}

func (c Codec) encodeInt64(in int64) (out string) {
	if in == 0 {
		return string(c.GetAlphabet()[0])
	}
	var buffer []rune
	for in > 0 {
		buffer = append(buffer, c.alphabetSlice[in%c.baseInt64])
		in /= c.baseInt64
	}
	reverted := c.runSliceRev(buffer)
	return string(reverted)
}

func (c Codec) encodeUint64(in uint64) (out string) {
	if in == 0 {
		return string(c.GetAlphabet()[0])
	}
	var buffer []rune
	for in > 0 {
		buffer = append(buffer, c.alphabetSlice[in%c.baseUint64])
		in /= c.baseUint64
	}
	reverted := c.runSliceRev(buffer)
	return string(reverted)
}

func (c Codec) decodeInt64(str string) (out int64, err error) {
	for _, character := range str {
		chIdxInAlphabet, exists := c.alphabetMapInt64[character]
		if !exists {
			return 0, fmt.Errorf("character `%c` does not exist in alphabet `%s`", character, c.GetAlphabet())
		}
		out = out*c.baseInt64 + chIdxInAlphabet
	}
	return out, nil
}

func (c Codec) decodeUint64(str string) (out uint64, err error) {
	for _, character := range str {
		chIdxInAlphabet, exists := c.alphabetMapUint64[character]
		if !exists {
			return 0, fmt.Errorf("character `%c` does not exist in alphabet `%s`", character, c.GetAlphabet())
		}
		out = out*c.baseUint64 + chIdxInAlphabet
	}
	return out, nil
}

func (c *Codec) setAlphabet(alphabet string) {
	c.alphabet = alphabet
	c.alphabetSlice = []rune(alphabet)

	// int64 stuff
	int64Map := make(map[rune]int64)
	for i, rune := range c.alphabet {
		int64Map[rune] = int64(i)
	}
	c.alphabetMapInt64 = int64Map
	c.baseInt64 = int64(len(int64Map))

	// uint64 stuff
	uint64Map := make(map[rune]uint64)
	for i, rune := range c.alphabet {
		uint64Map[rune] = uint64(i)
	}
	c.alphabetMapUint64 = uint64Map
	c.baseUint64 = uint64(len(uint64Map))
}

func (c Codec) runSliceRev(in []rune) []rune {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
	return in
}
