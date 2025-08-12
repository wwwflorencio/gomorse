package morse

import "testing"

func TestDecode(t *testing.T) {
	decoder := NewMorseDecoder()

	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Sentence with word separation",
			input:    ".... . -.--   .--- ..- -.. .",
			expected: "HEY JUDE",
		},
		{
			name:     "Sentence without word separation",
			input:    ".... . -.-- .--- ..- -.. .",
			expected: "HEYJUDE",
		},
		{
			name:     "Extra spaces at start and end",
			input:    "   .... . -.--   .--- ..- -.. .   ",
			expected: "HEY JUDE",
		},
		{
			name:     "Empty message",
			input:    "",
			expected: "",
		},
		{
			name:     "Invalid code",
			input:    "...-.-",
			expected: "?",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := decoder.Decode(testCase.input)
			if got != testCase.expected {
				t.Errorf("Decode(%q) = %q; want %q", testCase.input, got, testCase.expected)
			}
		})
	}
}
