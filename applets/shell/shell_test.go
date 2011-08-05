package shell

import (
	"testing"
)

func TestParameterizeEmptyString(t *testing.T) {
	emptyline := ""
	params, e := parameterize(emptyline)
	if params == nil {
		t.Fatal("Expected empty slice (len = 0), got nil")
	}
	l := len(params)
	if l != 0 {
		t.Fatalf("Expected empty slice, got len = %d", l)
	}
	if e != nil {
		t.Fatalf("Expected empty slice, got an error: %s", e.String())
	}
}

func TestParameterizeStrings(t *testing.T) {
	lines := map[string][]string{
		"1 22 333":                      []string{"1", "22", "333"},
		"\"111 11 11\" 222 \"333 333\"": []string{"111 11 11", "222", "333 333"},
		"\"111 \\\" 111\" 222":          []string{"111 \" 111", "222"},
	}
	for sin, sout := range lines {
		params, e := parameterize(sin)
		if e != nil {
			t.Errorf("Splitting <%s>: Expected %s, got an error: %s", sin, stringifyStringArray(sout), e.String())
		}
		if !compareStringArrays(sout, params) {
			t.Errorf("Splitting <%s>: Expected %s, got %s", sin, stringifyStringArray(sout), stringifyStringArray(params))
		}
	}
}

func compareStringArrays(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func stringifyStringArray(a []string) (s string) {
	s = "{"
	for i, _ := range a {
		s += "\"" + a[i] + "\""
		if i+1 != len(a) {
			s += ", "
		}
	}
	s += "}"
	return
}
