package main

import (
	"bytes"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	type inputResult struct {
		c    Config
		file string
	}
	const testDataDir = "testdata/main_functionality/"
	testCases := []struct {
		name string
		r    inputResult
	}{
		{name: "EmptyString", r: inputResult{c: NewConfig(""), file: "empty.txt"}},
		{name: "NewLine", r: inputResult{c: NewConfig("\n"), file: "newLine.txt"}},
		{name: "Hello", r: inputResult{c: NewConfig("hello"), file: "hello.txt"}},
		{name: "HeL10", r: inputResult{c: NewConfig("HeL10"), file: "hel10.txt"}},
		{name: "HelloThere", r: inputResult{c: NewConfig("Hello There"), file: "helloThere.txt"}},
		{name: "1Hello2There", r: inputResult{c: NewConfig("1Hello 2There"), file: "1hello2there.txt"}},
		{name: "CurlyHelloThere", r: inputResult{c: NewConfig("{Hello There}"), file: "curly_hello_there.txt"}},
		{name: "HelloNewLine", r: inputResult{c: NewConfig("Hello\nThere"), file: "hello_new_line.txt"}},
		{name: "HellowNewLine", r: inputResult{c: NewConfig("Hello\n\nThere"), file: "hello_2new_line.txt"}},
		{name: "Curly", r: inputResult{c: NewConfig("{|}~"), file: "curly.txt"}},
		{name: "HelloS", r: inputResult{c: Config{text: "hello", template: "standard"}, file: "hello.txt"}},
		{name: "HelloWorldSh", r: inputResult{c: Config{text: "hello world", template: "shadow"}, file: "helloWorldShadow.txt"}},
		{name: "N2MYT", r: inputResult{c: Config{text: "nice 2 meet you", template: "thinkertoy"}, file: "N2MYT.txt"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.ReadFile(testDataDir + tc.r.file)
			expected := string(file[:])
			if err != nil {
				t.Fatal(err)
			}
			var res bytes.Buffer
			err = run(tc.r.c, &res)
			if err != nil {
				t.Fatal(err)
			}

			if res.String() != expected {
				t.Errorf("Expected %s, got %s instead.\n", expected, res.String())
			}

		})
	}
}
