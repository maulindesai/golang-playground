package main

import "testing"

func BenchmarkStringConcatWithoutJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringConcatWithoutJoin()
	}
}

func BenchmarkStringConcatWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringConcatWithJoin()
	}
}
