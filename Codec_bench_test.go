package uint2str

import (
	"math"
	"testing"
)

func BenchmarkEncodeInt8(b *testing.B) {
	tv := int64(math.MaxInt8)
	for n := 0; n < b.N; n++ {
		_, err := inst.Encode(tv)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEncodeUint8(b *testing.B) {
	tv := uint64(math.MaxInt8)
	for n := 0; n < b.N; n++ {
		_, err := inst.Encode(tv)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEncodeInt64(b *testing.B) {
	tv := int64(math.MaxInt64)
	for n := 0; n < b.N; n++ {
		_, err := inst.Encode(tv)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEncodeUint64(b *testing.B) {
	tv := uint64(math.MaxUint64)
	for n := 0; n < b.N; n++ {
		_, err := inst.Encode(tv)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkDecodeInt8(b *testing.B) {
	b.StopTimer()
	var decoded int8
	maxInt8 := int8(math.MaxInt8)
	encoded, err := inst.Encode(maxInt8)
	if err != nil {
		b.Error(err)
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		err := inst.Decode(encoded, &decoded)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkDecodeUint8(b *testing.B) {
	b.StopTimer()
	var decoded uint8
	maxUint8 := uint8(math.MaxUint8)
	encoded, err := inst.Encode(maxUint8)
	if err != nil {
		b.Error(err)
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		err := inst.Decode(encoded, &decoded)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkDecodeInt64(b *testing.B) {
	b.StopTimer()
	var decoded int64
	maxInt64 := int64(math.MaxInt64)
	encoded, err := inst.Encode(maxInt64)
	if err != nil {
		b.Error(err)
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		err := inst.Decode(encoded, &decoded)
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkDecodeUint64(b *testing.B) {
	b.StopTimer()
	var decoded uint64
	maxUint64 := uint64(math.MaxUint64)
	encoded, err := inst.Encode(maxUint64)
	if err != nil {
		b.Error(err)
	}
	b.StartTimer()
	for n := 0; n < b.N; n++ {
		err := inst.Decode(encoded, &decoded)
		if err != nil {
			b.Error(err)
		}
	}
}
