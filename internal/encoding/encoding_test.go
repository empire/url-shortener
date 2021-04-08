package encoding

import "testing"

func TestEncoding(t *testing.T) {
	tests := []struct {
		number   int64
		expected string
	}{
		{1, "b"},
		{10 * 62 * 62 * 62, "aaak"},
		{3*62*62 + 2*62 + 1, "bcd"},
	}

	for _, sample := range tests {
		actual := Encode(sample.number)
		if actual != sample.expected {
			t.Fatalf("Encode(%d) wants: %s, got: %s", sample.number, sample.expected, actual)
		}
	}

}
