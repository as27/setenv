package setenv

import (
	"reflect"
	"testing"
)

var inFile = `VAR1=123
# Comment
MY_VAR=ABCDE
TEST=ab123

`

func mockReadFile(filename string) ([]byte, error) {
	return []byte(inFile), nil
}

func TestLoadLines(t *testing.T) {
	readFile = mockReadFile
	lines, _ := loadLines("")
	expect := []string{
		"VAR1=123",
		"MY_VAR=ABCDE",
		"TEST=ab123",
	}
	if !reflect.DeepEqual(lines, expect) {
		t.Errorf("-----\nGot: %#v\nExp: %#v", lines, expect)
	}
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		input  string
		expect Env
		err    bool
	}{
		{
			"VAR1=123",
			Env{"VAR1", "123"},
			false,
		},
		{
			"MY_VAR  =Aab123  ",
			Env{"MY_VAR", "Aab123"},
			false,
		},
		{
			"MY_VAR    ",
			Env{},
			true,
		},
		{
			"MY_VAR  =Aab123 = 123 ",
			Env{},
			true,
		},
	}
	for _, test := range tests {
		got, err := parseLine(test.input)
		if test.err && err == nil {
			t.Errorf("Expect an error for input: %s", test.input)
		}
		if got != test.expect {
			t.Errorf("----\nGot: %s\nExp: %s", got, test.expect)
		}
	}
}
func TestParseFile(t *testing.T) {
	expect := []Env{
		Env{"VAR1", "123"},
		Env{"MY_VAR", "ABCDE"},
		Env{"TEST", "ab123"},
	}
	got, _ := ParseFile("")
	if !reflect.DeepEqual(got, expect) {
		t.Errorf("-----\nGot: %#v\nExp: %#v", got, expect)
	}
}
