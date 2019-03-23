package jsontagchecker

import (
	"fmt"
	"strings"
	"testing"
)

func TestIsJsonTag(t *testing.T) {
	cases := map[string]struct {
		tag    string
		expect bool
	}{
		"true": {
			tag:    `json:"abcdefg"`,
			expect: true,
		},
		"false": {
			tag:    `hogehoge`,
			expect: false,
		},
	}

	for n, tc := range cases {
		tc := tc
		t.Run(n, func(t *testing.T) {
			if IsJsonTag(tc.tag) != tc.expect {
				t.Errorf("want %v but actual %v\n", tc.expect, IsJsonTag(tc.tag))
			}
		})
	}
}

func TestChecker(t *testing.T) {
	cases := map[string]struct {
		field  string
		tag    string
		expect bool
	}{
		"success": {
			field:  "AbcAbc",
			tag:    `json:"abc_abc"`,
			expect: true,
		},
		"fail": {
			field:  "AbcAbc",
			tag:    `json:"abc_cba"`,
			expect: false,
		},
	}

	for n, tc := range cases {
		t.Run(n, func(t *testing.T) {
			fmt.Println(strings.Split(tc.tag, "\""))
			if Checker(tc.field, tc.tag) != tc.expect {
				t.Errorf("want %v, but actual %v\n", tc.expect, Checker(tc.field, tc.tag))
			}
		})
	}
}
