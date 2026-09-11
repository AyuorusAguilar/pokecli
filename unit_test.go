package main
import "testing"

func TestCleanInput(te *testing.T)  {
	type test struct{
		input string
		expectedOutput []string
	}

	cases := []test{
		{input: "Ismael Manzanero", expectedOutput: []string{"ismael", "manzanero"}},
		{input: "", expectedOutput: []string{}},
		{input: "    Ismael Manzanero     ", expectedOutput: []string{"ismael", "manzanero"}},
		{input: "IsmaelManzanero", expectedOutput: []string{"ismaelmanzanero"}},
		{input: "I s m a e l", expectedOutput: []string{"i", "s", "m", "a", "e", "l"}},
	}

	for _, c := range cases {
		actualOutput := cleanInput(c.input)

		if len(actualOutput) != len(c.expectedOutput) {
			te.Errorf("\nError:\n  Expected:\n    %v\n  Got:\n    %v", c.expectedOutput, actualOutput)
			continue
		}

		for i, actualWord := range actualOutput {
			if actualWord != c.expectedOutput[i] {
				te.Errorf("\nError:\n  Expected:\n    %v\n  Got:\n    %v", c.expectedOutput, actualOutput)
			}
		}

	}
}