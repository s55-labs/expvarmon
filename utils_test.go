package main

import "testing"

func TestUtils(t *testing.T) {
	def := "memstats.Alloc,memstats.Sys"
	extra := ""

	vars, err := ParseVars(def, extra)
	if err != nil {
		t.Fatalf("Err not nil: %v", err)
	}

	if len(vars) != 2 {
		t.Fatalf("vars should contain 2 elements, but has %d", len(vars))
	}

	def = "memstats.Alloc,memstats.Sys"
	extra = "goroutines,counter.A"

	vars, err = ParseVars(def, extra)
	if err != nil {
		t.Fatalf("Err not nil: %v", err)
	}

	if len(vars) != 4 {
		t.Fatalf("vars should contain 4 elements, but has %d", len(vars))
	}

	strs := dot2slice("memstats.Alloc")
	if len(strs) != 2 || strs[0] != "memstats" || strs[1] != "Alloc" {
		t.Fatalf("dot2slice wrong err: %v", strs)
	}
}
