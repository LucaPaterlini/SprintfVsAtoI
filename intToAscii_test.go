package inttoascii

import "testing"

func BenchmarkIntegerToASCII(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntegerToASCII(i)
	}
}

func BenchmarkIntegerToAscii2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntegerToASCII2(i)
	}
}
