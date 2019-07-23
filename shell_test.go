package main

import "testing"

func TestPrintShellIntegration(t *testing.T) {
	want := `Hello there
This is a test.`
	got := PrintShellIntegration(want)
	if want != got {
		t.Errorf("Got: %s, want: %s", got, want)
	}
}
