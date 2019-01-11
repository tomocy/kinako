package repl

import (
	"bytes"
	"testing"
)

func TestPrintResult(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"5;", "5\n"},
		{"-6;", "-6\n"},
		{"7 + 8 - 9 * 10 / 11;", "15\n"},
		{"(12 + 13) / 14;", "1\n"},
		{"var x int = 15;", "15\n"},
		{"var x int = 16; x;", "16\n"},
		{"15; 16", "failed to find semicolon\n"},
		{"5 / 0;", "divided by zero\n"},
	}
	for _, test := range tests {
		w := new(bytes.Buffer)
		mock := REPL{
			Writer: w,
		}
		mock.printResult(test.input)
		actual := w.String()
		if actual != test.expected {
			t.Errorf("unexpected result: got %s, but expected %s\n", actual, test.expected) // TODO: use %q as formatter
		}
	}
}
