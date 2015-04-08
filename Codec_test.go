package uint2str

import (
	"errors"
	"math"
	"reflect"
	"testing"
)

var inst = NewCodec(ALPHABET)

var testEncodedStr = "bM"

var encExp = map[int]string{
	100: "bM",
}

var maxInts = map[interface{}]string{
	int8(math.MaxInt8):   "cd",
	int16(math.MaxInt16): "iGF",
	int32(math.MaxInt32): "cvuMLb",
	int64(math.MaxInt64): "k9viXaIfiWh",
}

var maxUints = map[interface{}]string{
	uint8(math.MaxUint8):   "eh",
	uint16(math.MaxUint16): "rdb",
	uint32(math.MaxUint32): "eQPpmd",
	uint64(math.MaxUint64): "v8QrKbgkrIp",
}

func TestGetAlphabet(t *testing.T) {
	if inst.GetAlphabet() != ALPHABET {
		t.Errorf("error in alphabet getter. expected: `%s`, given: `%s`", ALPHABET, inst.GetAlphabet())
	}
}

func TestZeroValEncode(t *testing.T) {
	zeroInt := int(0)
	zeroIntEncoded, err := inst.Encode(zeroInt)
	if err != nil {
		t.Error(err)

	}

	zeroUint := uint(0)
	zeroUintEncoded, err := inst.Encode(zeroUint)
	if err != nil {
		t.Error(err)
	}

	expString := string(ALPHABET[0])
	if zeroIntEncoded != zeroUintEncoded || zeroIntEncoded != expString {
		t.Errorf("zero encode failed. expected: %s", expString)
	}
}

func TestIntegerEncodeDecode(t *testing.T) {
	for num, expected := range maxInts {
		res, err := inst.Encode(num)
		if err != nil {
			t.Error(err)
		}
		if res != expected {
			t.Errorf("encode of %s failed. expected: `%s`, given: `%s`", reflect.TypeOf(num).String(), expected, res)
		}
		expectedNum := reflect.ValueOf(num).Int()
		num := expectedNum
		err = inst.Decode(res, &num)
		if err != nil {
			t.Error(err)
		}
		if expectedNum != num {
			t.Errorf("decode failed. expected: `%d`, given: `%d`", expectedNum, num)
		}
	}
}

func TestUnsignedIntegerEncodeDecode(t *testing.T) {
	for num, expected := range maxUints {
		res, err := inst.Encode(num)
		if err != nil {
			t.Error(err)
		}
		if res != expected {
			t.Errorf("encode of %s failed. expected: `%s`, given: `%s`", reflect.TypeOf(num).String(), expected, res)
		}
		expectedNum := reflect.ValueOf(num).Uint()
		num := expectedNum
		err = inst.Decode(res, &num)
		if err != nil {
			t.Error(err)
		}
		if expectedNum != num {
			t.Errorf("decode failed. expected: `%d`, given: `%d`", expectedNum, num)
		}
	}
}

func TestEncodeDecodeDeepPointers(t *testing.T) {
	expVar := 100
	intVar := expVar
	pIntVar := &intVar
	ppIntVar := &pIntVar
	pppIntVar := &ppIntVar
	tVar := &pppIntVar // ****int8

	encoded, err := inst.Encode(tVar)
	if err != nil {
		t.Error(err)
	}
	encExpected := encExp[100]
	if encoded != encExpected {
		t.Errorf("encode of %d failed. expected: %s, given: %s", tVar, encExpected, encoded)
	}

	err = inst.Decode(encoded, tVar)
	if err != nil {
		t.Error(err)
	}
}

func TestEncodeIncorrectType(t *testing.T) {
	var err error
	errExp := errors.New("Expected error")
	_, err = inst.Encode(nil)
	if err == nil {
		t.Error(errExp)
	}

	var tInterface interface{}
	_, err = inst.Encode(tInterface)
	if err == nil {
		t.Error(errExp)
	}

	var tComplex128 complex128
	_, err = inst.Encode(tComplex128)
	if err == nil {
		t.Error(errExp)
	}

	var tComplex64 complex64
	_, err = inst.Encode(tComplex64)
	if err == nil {
		t.Error(errExp)
	}

	var tStruct struct{}
	_, err = inst.Encode(tStruct)
	if err == nil {
		t.Error(errExp)
	}

	var tBool bool
	_, err = inst.Encode(tBool)
	if err == nil {
		t.Error(errExp)
	}

	var tSlice []int64
	_, err = inst.Encode(tSlice)
	if err == nil {
		t.Error(errExp)
	}
}

func TestDecodeIncorrectType(t *testing.T) {
	var err error
	errExp := errors.New("expected error")
	err = inst.Decode(testEncodedStr, nil)
	if err == nil {
		t.Error(errExp)
	}

	var tInterface interface{}
	err = inst.Decode(testEncodedStr, &tInterface)
	if err == nil {
		t.Error(errExp)
	}

	var tComplex128 complex128
	err = inst.Decode(testEncodedStr, &tComplex128)
	if err == nil {
		t.Error(errExp)
	}

	var tComplex64 complex64
	err = inst.Decode(testEncodedStr, &tComplex64)
	if err == nil {
		t.Error(errExp)
	}

	var tStruct struct{}
	err = inst.Decode(testEncodedStr, &tStruct)
	if err == nil {
		t.Error(errExp)
	}

	var tBool bool
	err = inst.Decode(testEncodedStr, &tBool)
	if err == nil {
		t.Error(errExp)
	}

	var tSlice []int64
	err = inst.Decode(testEncodedStr, &tSlice)
	if err == nil {
		t.Error(errExp)
	}
}

func TestDecodeAlphabetError(t *testing.T) {
	var iInt int
	err := inst.Decode("|", &iInt)
	if err == nil {
		t.Error("expected error")
	}

	var iUint uint
	err = inst.Decode("|", &iUint)
	if err == nil {
		t.Error("expected error")
	}
}

func TestDecodeNonPtr(t *testing.T) {
	var i int
	err := inst.Decode(testEncodedStr, i)
	if err == nil {
		t.Error("expected error")
	}
}

func TestDecodeEmptyStr(t *testing.T) {
	var i int
	err := inst.Decode("", i)
	if err == nil {
		t.Error("expected error")

	}
}
