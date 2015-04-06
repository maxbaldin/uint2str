package uint2str

import (
	"reflect"
	"testing"
)

var BaseInstance = newBaseCodec(ALPHABET)

func TestSetAlphabet(t *testing.T) {
	alphabet := "?abc"
	codec := newBaseCodec(alphabet)
	if string(codec.getAlphabet()) != alphabet || !reflect.DeepEqual(codec.getAlphabet(), []rune(alphabet)) {
		t.Error("Error while setting up alphabet")
	}
}

func TestReverse(t *testing.T) {
	for str, expected := range testReverseStrings {
		given := BaseInstance.reverse(str)
		if given != expected {
			t.Errorf("String `%s` expected after revert: `%s`, given: `%s`", str, expected, given)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BaseInstance.reverse("The quick brown ? jumped over the lazy ?")
	}
}
