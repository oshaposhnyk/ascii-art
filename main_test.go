package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	type inputResult struct {
		input      string
		resultFile string
	}
	testCases := []struct {
		name   string
		result inputResult
	}{
		{name: "EmptyString", result: inputResult{input: "", resultFile: "testdata/main_functionality/empty.txt"}},
		{name: "NewLine", result: inputResult{input: "\n", resultFile: "testdata/main_functionality/newLine.txt"}},
		{name: "Hello", result: inputResult{input: "hello", resultFile: "testdata/main_functionality/hello.txt"}},
		{name: "HeL10", result: inputResult{input: "HeL10", resultFile: "testdata/main_functionality/hel10.txt"}},
		{name: "HelloThere", result: inputResult{input: "Hello There", resultFile: "testdata/main_functionality/helloThere.txt"}},
		{name: "1Hello2There", result: inputResult{input: "1Hello 2There", resultFile: "testdata/main_functionality/1hello2there.txt"}},
		{name: "CurlyHelloThere", result: inputResult{input: "{Hello There}", resultFile: "testdata/main_functionality/curly_hello_there.txt"}},
		{name: "HelloNewLine", result: inputResult{input: "Hello\nThere", resultFile: "testdata/main_functionality/hello_new_line.txt"}},
		{name: "HellowNewLine", result: inputResult{input: "Hello\n\nThere", resultFile: "testdata/main_functionality/hello_2new_line.txt"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.ReadFile(tc.result.resultFile)
			expected := string(file[:])
			if err != nil {
				t.Fatal(err)
			}
			input := []string{tc.result.input}
			result, err := run(input)
			if err != nil {
				t.Fatal(err)
			}

			if result != expected {
				t.Errorf("Expected %s, got %s instead.\n", expected, result)
			}

		})
	}
}
