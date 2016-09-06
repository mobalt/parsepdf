package parsepdf1

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var (
	cleanup = strings.NewReplacer(
		//old, new
		"&amp;", "&",
		"&lt", "<",
		"&gt", ">",
		"<b>", "",
		"</b", "",
		"<i>", "",
		"</i>", "",
	)
	textElement *regexp.Regexp = regexp.MustCompile("^<text\\D*(\\d+)\\D+(\\d+)\\D+(\\d+)\\D+(\\d+)\\D+(\\d+)\"> *(.+?) *</text>")
)

func internal_stripper(r rune) rune {
	// could have stopped at 126 but the additional
	// accented chars might be used for foreign names?
	// maybe? also greek letters
	// also, 127 == delete
	if r >= 32 && r <= 248 && r != 127 {
		return r
	}
	return -1
}
func stripWeirdChars(str string) string {
	return strings.Map(internal_stripper, str)
}

func ReadFile(filename string) *Pages {
	file, ok := os.Open(filename)
	if e != nil {
		panic(e)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		raw := scanner.Text()
		switch raw[:6] {
		case "</page":
			// process page

			continue
		case "<text ":
			// process text
		}

		raw = cleanup.Replace(raw)
		raw = stripWeirdChars(raw)
		match := textElement.FindStringSubmatch(raw)
		if len(match) == 0 {
			continue
		}
		text := match[6]
		x, _ := match[2]
		y, _ := match[1]
		w, _ := match[3]
		h, _ := match[4]

	}

}
