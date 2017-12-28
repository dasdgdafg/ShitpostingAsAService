package main

import (
	"testing"
)

// exact rolling

func BenchmarkExact1d20(b *testing.B) {
	sides := int64(20)
	dice := int64(1)
	for i := 0; i < b.N; i++ {
		rollExact(sides, dice)
	}
}

func BenchmarkExact1000d20(b *testing.B) {
	sides := int64(20)
	dice := int64(1000)
	for i := 0; i < b.N; i++ {
		rollExact(sides, dice)
	}
}

func BenchmarkExact1000000d20(b *testing.B) {
	sides := int64(20)
	dice := int64(1000000)
	for i := 0; i < b.N; i++ {
		rollExact(sides, dice)
	}
}

func BenchmarkExact1d1000000(b *testing.B) {
	sides := int64(1000000)
	dice := int64(1)
	for i := 0; i < b.N; i++ {
		rollExact(sides, dice)
	}
}

func BenchmarkExact1000d1000000(b *testing.B) {
	sides := int64(1000000)
	dice := int64(1000)
	for i := 0; i < b.N; i++ {
		rollExact(sides, dice)
	}
}

func BenchmarkExact1000000d1000000(b *testing.B) {
	sides := int64(1000000)
	dice := int64(1000000)
	for i := 0; i < b.N; i++ {
		rollExact(sides, dice)
	}
}

func BenchmarkExact1d9223372036854775807(b *testing.B) {
	sides := int64(9223372036854775807)
	dice := int64(1)
	for i := 0; i < b.N; i++ {
		rollExact(sides, dice)
	}
}

func BenchmarkExact1000d9223372036854775807(b *testing.B) {
	sides := int64(9223372036854775807)
	dice := int64(1000)
	for i := 0; i < b.N; i++ {
		rollExact(sides, dice)
	}
}

func BenchmarkExact1000000d9223372036854775807(b *testing.B) {
	sides := int64(9223372036854775807)
	dice := int64(1000000)
	for i := 0; i < b.N; i++ {
		rollExact(sides, dice)
	}
}

// approximations (should take the same time regardless of the numbers)

func BenchmarkApprox1000000d20(b *testing.B) {
	sides := int64(20)
	dice := int64(1000000)
	for i := 0; i < b.N; i++ {
		rollApprox(sides, dice)
	}
}

func BenchmarkApprox9223372036854775807d20(b *testing.B) {
	sides := int64(20)
	dice := int64(9223372036854775807)
	for i := 0; i < b.N; i++ {
		rollApprox(sides, dice)
	}
}

func BenchmarkApprox1000000d1000000(b *testing.B) {
	sides := int64(1000000)
	dice := int64(1000000)
	for i := 0; i < b.N; i++ {
		rollApprox(sides, dice)
	}
}

func BenchmarkApprox9223372036854775807d1000000(b *testing.B) {
	sides := int64(1000000)
	dice := int64(9223372036854775807)
	for i := 0; i < b.N; i++ {
		rollApprox(sides, dice)
	}
}

func BenchmarkApprox1000000d9223372036854775807(b *testing.B) {
	sides := int64(9223372036854775807)
	dice := int64(1000000)
	for i := 0; i < b.N; i++ {
		rollApprox(sides, dice)
	}
}

func BenchmarkApprox9223372036854775807d9223372036854775807(b *testing.B) {
	sides := int64(9223372036854775807)
	dice := int64(9223372036854775807)
	for i := 0; i < b.N; i++ {
		rollApprox(sides, dice)
	}
}
