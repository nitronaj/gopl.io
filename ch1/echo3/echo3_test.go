package main

import (
	"testing"
)

func BenchmarkJoinArgs(b *testing.B) {
	args := [10]string{"Hi","Nader","How", "are", "you", "I", "am", "awesome", "bye", "bye"}
    JoinArgs(args[:])
}
