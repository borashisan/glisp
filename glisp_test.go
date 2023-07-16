package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  []string
		err   error
	}{
		{name: "single quotation", input: "'a\n", want: []string{"'", "a"}, err: nil},
		{name: "single quotation without newline", input: "'a", want: []string{"'", "a"}, err: nil},
		{name: "single quotation with space", input: "' a\n", want: []string{"'", "a"}, err: nil},
		{name: "single quotation after token", input: "a'\n", want: nil, err: errors.New("*** - SYSTEM::READ-EVAL-PRINT: variable A has no value")},
	}

	for _, tc := range tests {
		got, _ := parser(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
