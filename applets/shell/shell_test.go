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
		"1 22 333":                      {"1", "22", "333"},
		"\"111 11 11\" 222 \"333 333\"": {"111 11 11", "222", "333 333"},
		"\"111 \\\" 1\\11\" 222":        {"111 \" 1\\11", "222"},
	}
	for sin, sout := range lines {
		res, e := parameterize(sin)
		if e != nil {
			t.Errorf("Splitting <%s>: Expected %s, got an error: %s", sin, stringifyStringArray(sout), e.String())
		}
		if !compareStringArrays(sout, res) {
			t.Errorf("Splitting <%s>: Expected %s, got %s", sin, stringifyStringArray(sout), stringifyStringArray(res))
		}
	}
}

func TestCleanMatch(t *testing.T) {
	lines := map[string]string{
		"abc":        "abc",
		"a\\\"b":     "a\"b",
		"a\\\\\\\"b": "a\\\"b",
		"a\\\\b":     "a\\b",
		"a\\t":       "a\\t",
	}

	for sin, sout := range lines {
		res := cleanMatch(sin)
		if res != sout {
			t.Errorf("Cleaning <%s>: Expected <%s>, got <%s>", sin, sout, res)
		}
	}
}

func TestCommentRecognition(t *testing.T) {
	positives := []string{
		"#",
		"#!/bin/bash",
		"     #",
		"\t     #",
		"\t \t#",
	}
	for i, line := range positives {
		if !isComment(line) {
			t.Errorf("<%s> (#%d) not recognized as comment", line, i)
		}
	}

	negatives := []string{
		"asdf",
		"    asd",
		"        a#",
		"      !#",
		"     \a#",
	}
	for i, line := range negatives {
		if isComment(line) {
			t.Errorf("<%s> (#%d) recognized as comment", line, i)
		}
	}
}

func compareStringArrays(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func stringifyStringArray(a []string) (s string) {
	s = "{"
	for i := range a {
		s += "\"" + a[i] + "\""
		if i+1 != len(a) {
			s += ", "
		}
	}
	s += "}"
	return
}
