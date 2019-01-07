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
		{"5", "5\n"},
	}
	for _, test := range tests {
		w := new(bytes.Buffer)
		mock := REPL{
			Writer: w,
		}
		mock.printResult(test.input)
		actual := w.String()
		if actual != test.expected {
			t.Errorf("unexpected result: got %s, but expected %s\n", actual, test.expected)
		}
	}
}
