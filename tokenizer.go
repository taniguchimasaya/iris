package main

import (
	"regexp"
	"strings"
)

// TokenReader interface type is the interface for reading string with every token
type TokenReader interface {
	ReadToken() (string, error)
}

func concatMatcher(src ...string) *regexp.Regexp {
	return regexp.MustCompile("^(?:" + strings.Join(src, ")$|^(?:") + ")$")
}

var macro = strings.Join([]string{"#(?:[[:digit:]]+[aA]?)?", ",@?", "'", "`"}, "|")
var integer = strings.Join([]string{"[[:digit:]]+", "[+-][[:digit:]]*", "#(?:[bB][01]*)?", "#(?:[oO][0-7]*)?", "#(?:[xX][[:xdigit:]]*)?"}, "|")
var float = strings.Join([]string{"[[:digit:]]+(?:\\.?[[:digit:]]*(?:[eE](?:[-+]?[[:digit:]]*)?)?)?", "[+-](?:[[:digit:]]+(?:\\.?[[:digit:]]*(?:[eE](?:[-+]?[[:digit:]]*)?)?)?)?"}, "|")
var character = strings.Join([]string{"#(?:\\\\[[:alpha:]]*)?", "#(?:\\\\[[:graph:]]?)?"}, "|")
var str = strings.Join([]string{"\"(?:\\\\\"|[^\"])*\"?"}, "|")
var symbol = strings.Join([]string{"[<>/*=?_!$%[\\]^{}~a-zA-Z][<>/*=?_!$%[\\]^{}~0-9a-zA-Z]*", "\\|(?:\\\\\\||[^|])*\\|?"}, "|")
var parentheses = strings.Join([]string{"\\.", "\\(", "\\)"}, "|")

var token = concatMatcher(
	macro,
	integer,
	float,
	character,
	str,
	symbol,
	parentheses)

// ReadToken returns error or string as token
func (r *Reader) ReadToken() (string, error) {
	buf := ""
	for {
		if buf == "" {
			p, _, err := r.PeekRune()
			if err != nil {
				return "", err
			}
			if token.MatchString(string(p)) {
				buf = string(p)
			}
		} else {
			p, _, err := r.PeekRune()
			if err != nil {
				return buf, nil
			}
			if !token.MatchString(buf + string(p)) {
				return buf, nil
			}
			buf += string(p)
		}
		r.ReadRune()
	}
}
