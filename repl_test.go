package main
import ("testing")

func TestCleanInput(t *testing.T) {
    cases := []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	{
		input:    "jello humble       world  ",
		expected: []string{"jello", "humble", "world"},
	},
	{
		input:    "ThinK I LIKE poKemon",
		expected: []string{"think", "i", "like", "pokemon"},
	},
	{
		input:    "ThISisOnELonGWord",
		expected: []string{"thisisonelongword"},
	},
	{
		input:    "ThISisOnELonGWordwithlotsofwhitespace                       ",
		expected: []string{"thisisonelongwordwithlotsofwhitespace"},
	},
}
for _, c := range cases {
	actual := cleanInput(c.input)
	// Check the length of the actual slice against the expected slice
	// if they don't match, use t.Errorf to print an error message
	// and fail the test
	if len(actual) != len(c.expected) {
    		t.Fatalf("length mismatch: expected %d, got %d", len(c.expected), len(actual))
		}
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if word != expectedWord {
			t.Fatalf("word mismatch: expected %s, got %s", expectedWord, word)
		}
	}
}
}