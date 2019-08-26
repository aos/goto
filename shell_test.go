package main

func ExamplePrintShellIntegration() {
	want := `Hello there
This is a test.`
	PrintShellIntegration(want)

	// Output:
	// Hello there
	// This is a test.
}
