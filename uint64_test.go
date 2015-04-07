package uint2str

import (
	"fmt"
	"testing"
)

var Uint64Instance = NewUint64Codec(ALPHABET)

var uint64EncDecTestData = map[uint64]string{
	0:          "a",
	100:        "bM",
	3466876560: "dWMO72",
	4294967295: "eQPpmd",
}

func TestUint64EncodeDecodeEquals(t *testing.T) {
	for i := uint64(100000000000); i < 100000100500; i++ {
		num := uint64(i)
		encoded := Uint64Instance.Encode(num)
		decoded, err := Uint64Instance.Decode(encoded)
		if err != nil {
			t.Error(err)
		}
		if num != decoded {
			t.Errorf("Error while decoding %d", num)
		}
	}
}

func TestUint64CustomAlphabet(t *testing.T) {
	alphabet := "abc"
	codec := NewUint64Codec(alphabet)
	num := uint64(4294967295)
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

func TestUint64Encode(t *testing.T) {
	for num, expected := range uint64EncDecTestData {
		encoded := Uint64Instance.Encode(num)
		if encoded != expected {
			t.Errorf("Error while encoding %d", num)
		}
	}
}

func TestUint64Decode(t *testing.T) {
	for expected, encoded := range uint64EncDecTestData {
		decoded, err := Uint64Instance.Decode(encoded)
		if err != nil {
			t.Error(err)
		}
		if decoded != expected {
			t.Errorf("Error while decoding %s", encoded)
		}
	}
}

func TestUint64DecodeError(t *testing.T) {
	invalidCharacter := "-"
	_, err := Uint64Instance.Decode(invalidCharacter)
	if err == nil {
		t.Error("Decode must produce error")
	}
	expectedErrorMessage := fmt.Sprintf("character %s does not exist in alphabet", invalidCharacter)
	if err.Error() != expectedErrorMessage {
		t.Error("Unexpected error message")
	}
}

func BenchmarkUint64Encode(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Uint64Instance.Encode(uint64(n))
	}
}
