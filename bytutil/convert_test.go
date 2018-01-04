package bytutil

import "testing"

func BenchmarkConvertStrAndBytes(b *testing.B) {
	str := "////////////pattern////////////"
	for i := 0; i < b.N; i++ {
		byt := StringToBytes(str)
		str = BytesToString(byt)
	}
}

func BenchmarkConvertStrAndBytes2(b *testing.B) {
	str := "////////////pattern////////////"
	for i := 0; i < b.N; i++ {
		byt := []byte(str)
		str = string(byt)
	}
}
