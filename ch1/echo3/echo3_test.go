package main

import (
	"testing"
)

func BenchmarkJoinArgs(b *testing.B) {
	args := [2]string{"Hi","Nader"}
    JoinArgs(args[:])
}
