package uint2str

import (
	"reflect"
	"testing"
)

var baseCodecRevTestData = map[string]string{
	"The quick brown 狐 jumped over the lazy 犬": "犬 yzal eht revo depmuj 狐 nworb kciuq ehT",
	"Test string":                              "gnirts tseT",
	"тестовая строка":                          "акортс яавотсет",
}
var BaseInstance = newBaseCodec(ALPHABET)

func TestSetAlphabet(t *testing.T) {
	alphabet := "犬abc"
	codec := newBaseCodec(alphabet)
	if string(codec.getAlphabet()) != alphabet || !reflect.DeepEqual(codec.getAlphabet(), []rune(alphabet)) {
		t.Error("Error while setting up alphabet")
	}
}

func TestReverse(t *testing.T) {
	for str, expected := range baseCodecRevTestData {
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

func newBaseCodec(alphabet string) *BaseCodec {
	sh := new(BaseCodec)
	sh.setAlphabet(alphabet)
	return sh
}
