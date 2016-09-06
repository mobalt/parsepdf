package parsepdf1

import (
	"regexp"
)

var reKeyValue *regexp.Regexp = regexp.MustCompile("^ *([^:\\d]{0,38}\\d?[^\\d:]{1,38}?) *: *(.*?) *$")

type Word struct {
	x, y, x2, y2 int
	key, value   string
	parent       *Line
}
type Words []*Word

type WordMap map[int][]*Word
type WordList []*([]*Word)

func NewWord(x int, y int, w int, h int, text string) *Word {
	word := &Word{
		x:  x,
		y:  y,
		x2: x + w,
		y2: y + h,
	}

	match := reKeyValue.FindStringSubmatch(text)

	switch {

	// No Key found, entire match must be value
	case len(match) == 0:
		word.value = text

	// possible Key & Value combo
	case match[2] != "":
		if len(match[1]) > 40 {
			// Heuristics: The possible key is too long to be real key,
			// probably just a value with an intentional/typo hyphen
			// Therefore, don't seperate. Lump as a single value
			word.value = text
		} else {
			// combo success
			word.key = match[1]
			word.value = match[2]
		}

		// only key found
	default:
		word.key = match[1]
	}
	return word
}

func (word *Word) SetParent(parent *Line) {
	word.parent = parent
}
