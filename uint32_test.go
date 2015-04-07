package uint2str

import (
	"fmt"
	"testing"
)

var Uint32Instance = NewUint32Codec(ALPHABET)

var uint32EncDecTestData = map[uint32]string{
	0:          "a",
	100:        "bM",
	3466876560: "dWMO72",
	4294967295: "eQPpmd",
}

func TestUint32EncodeDecodeEquals(t *testing.T) {
	iterationsCnt := 100500
	for i := 0; i < iterationsCnt; i++ {
		num := uint32(i)
		encoded := Uint32Instance.Encode(num)
		decoded, err := Uint32Instance.Decode(encoded)
		if err != nil {
			t.Error(err)
		}
		if num != decoded {
			t.Errorf("Error while decoding %d", num)
		}
	}
}

func TestUint32CustomAlphabet(t *testing.T) {
	alphabet := "abc"
	codec := NewUint32Codec(alphabet)
	num := uint32(4294967295)
	expectedEncoded := "bacaacacccabccbbbbcba"
	encoded := codec.Encode(num)
	if encoded != expectedEncoded {
		t.Errorf("Error while encoding %d", num)
	}
	decoded, err := codec.Decode(encoded)
	if err != nil {
		t.Error(err)
	}
	if decoded != num {
		t.Errorf("Error while decoding %s", encoded)
	}
}

func TestUint32Encode(t *testing.T) {
	for num, expected := range uint32EncDecTestData {
		encoded := Uint32Instance.Encode(num)
		if encoded != expected {
			t.Errorf("Error while encoding %d", num)
		}
	}
}

func TestUint32Decode(t *testing.T) {
	for expected, encoded := range uint32EncDecTestData {
		decoded, err := Uint32Instance.Decode(encoded)
		if err != nil {
			t.Error(err)
		}
		if decoded != expected {
			t.Errorf("Error while decoding %s", encoded)
		}
	}
}

func TestUint32DecodeError(t *testing.T) {
	invalidCharacter := "-"
	_, err := Uint32Instance.Decode(invalidCharacter)
	if err == nil {
		t.Error("Decode must produce error")
	}
	expectedErrorMessage := fmt.Sprintf("character %s does not exist in alphabet", invalidCharacter)
	if err.Error() != expectedErrorMessage {
		t.Error("Unexpected error message")
	}
}

func BenchmarkUint32Encode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Uint32Instance.Encode(uint32(n))
	}
}
