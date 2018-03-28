package parse

import (
	"fmt"
	"unicode/utf8"

	"github.com/faiface/funky/parse/parseinfo"
)

type Error struct {
	SourceInfo *parseinfo.Source
	Msg        string
}

func (err *Error) Error() string {
	return fmt.Sprintf("%v: %v", err.SourceInfo, err.Msg)
}

var SpecialRunes = []rune{'(', ')', '[', ']', '{', '}', ',', ';', '\\', 'λ', '#'}

func IsSpecialRune(r rune) bool {
	for _, special := range SpecialRunes {
		if r == special {
			return true
		}
	}
	return false
}

func IsSpecial(s string) bool {
	if utf8.RuneCountInString(s) != 1 {
		return false
	}
	return IsSpecialRune([]rune(s)[0])
}
