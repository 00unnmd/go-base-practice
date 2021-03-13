package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Louise"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Louise")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Louise") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}