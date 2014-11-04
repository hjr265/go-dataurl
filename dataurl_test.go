package dataurl

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	cases := [][]string{
		{"data:,", ""},
		{"data:text/x-csrc;base64,SGVsbG8sIHdvcmxkIQ==", "Hello, world!"},
	}

	for _, c := range cases {
		u, err := Parse(c[0])
		if err != nil {
			t.Fatal(err)
		}

		b := &bytes.Buffer{}
		_, err = io.Copy(b, u)
		if err != nil {
			t.Fatal(err)
		}

		if b.String() != c[1] {
			t.Fatalf("Expected `%s`, got `%s`", c[1], b.String())
		}
	}
}

func ExampleParse() {
	u, err := Parse("data:text/x-csrc;base64,SGVsbG8sIHdvcmxkIQ==")
	catch(err)

	_, err = io.Copy(os.Stdout, u)
	catch(err)

	// Output:
	// Hello, world!
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
