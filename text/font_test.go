package text

import (
	"testing"
)

func TestLoadFont(t *testing.T) {
	tests := []struct {
		path        string
		size        float64
		dpi         float64
		expectError bool
	}{
		{"testdata/valid.ttf", 24, 72, false},
		{"testdata/nonexistent.ttf", 24, 72, true},
		{"testdata/invalid.ttf", 24, 72, true},
	}

	for _, test := range tests {
		_, err := LoadFont(test.path, test.size, test.dpi)
		if err != nil && !test.expectError {
			t.Errorf("Expected no error, but got: %v", err)
		} else if err == nil && test.expectError {
			t.Errorf("Expected an error, but got nil for path: %s", test.path)
		}
	}
}
