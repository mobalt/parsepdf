package parsepdf1

type Line struct {
	y int
	words Words
	parent *Page
}
type Lines []*Line
type LineMap map[int]Lines

func (linemap *LineMap) Add(text string) {
	word := &NewWord(text)
	y := word.y
	line, ok := (*linemap)[y]
	if ok {
		line.words = append(line.words, word)
	} else {
		(*linemap)[y] = Line{
			y:y,
			words: Words{word}
		}
	}
}

func (linemap *LineMap) ToPage(minVerticalDistance int) (page *Page) {
	// create new page for output
	page = &Page{}

	// get all y keys into a list
	// then sort by ascending value
	yList := make([]int, 0, len(*wm))
	for y := range *wm {
		yList = append(yList, y)
	}
	sort.Ints(yList)

	var (
		maxY int = -1
		line *Line
		lines Lines
	)

	for y := range yList[1:] {
		if y >= maxY {
			maxY = y + minVerticalDistance
			// line is a pointer to the next leading line
			// all other lines that are vertically close
			// will be merged into it. And by not holding on to
			// their pointers, those merged lines will be
			// garbage collected
			line = (*linemap)[y]
			line.parent = page
			lines = append(lines, line)
		} else {
			// the line specified by "y" is too close to 
			// prior leading line, therefore just merge words
			line.words = append(line.words, (*linemap)[y].words...)
		}
	}
	page.lines = lines
	return page
}
