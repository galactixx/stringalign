package stringalign

import (
	"log"
	"slices"
	"strings"
	"unicode"

	"github.com/galactixx/ansiwalker"
	"github.com/galactixx/stringwrap"
	"github.com/mattn/go-runewidth"
)

// distributeSpaces divides the total number of spaces evenly across
// the specified number of gaps, returning a slice with the count of
// spaces assigned to each gap.
func distributeSpaces(gaps int, spaces int) []int {
	base, rem := spaces/gaps, spaces%gaps
	spacesPerGap := make([]int, gaps)
	for i := range spacesPerGap {
		spacesPerGap[i] = base
		if i < rem {
			spacesPerGap[i]++
		}
	}
	return spacesPerGap
}

type spaceGaps struct {
	indices   []int
	lastSpace bool
	gapIndex  int
}

// addGap records the index of a space gap in the line.
func (g *spaceGaps) addGap(index int) { g.indices = append(g.indices, index) }

// clearGap removes the first recorded gap and advances the gap index
// counter.
func (g *spaceGaps) clearGap() {
	g.indices = slices.Delete(g.indices, 0, 1)
	g.gapIndex += 1
}

// nextGap returns the index of the next available space gap, or -1 if
// none remain.
func (g *spaceGaps) nextGap() int {
	if len(g.indices) > 0 {
		return g.indices[0]
	}
	return -1
}

// alignFactory returns the alignment function corresponding to the
// given mode.
func alignFactory(align string) func(
	line string, spaces int, meta stringwrap.WrappedString,
) string {
	switch align {
	case "right":
		return rightAlign
	case "left":
		return leftAlign
	case "center":
		return centerAlign
	case "justify":
		return justify
	default:
		log.Fatalf("%s is not a valid align mode", align)
		return nil
	}
}

// rightAlign prepends spaces to shift the text to the right within the limit.
func rightAlign(line string, spaces int, _ stringwrap.WrappedString) string {
	prepend := strings.Repeat(" ", spaces)
	return prepend + line
}

// leftAlign appends spaces to shift the text to the left within the limit.
func leftAlign(line string, spaces int, _ stringwrap.WrappedString) string {
	postpend := strings.Repeat(" ", spaces)
	return line + postpend
}

// centerAlign evenly distributes spaces on both sides to center the text.
func centerAlign(line string, spaces int, _ stringwrap.WrappedString) string {
	// partition the difference and assign some number to be
	// spaces appended to the left and right of the text
	leftNumSpaces := spaces / 2
	rightNumSpaces := spaces - leftNumSpaces

	leftSpaces := strings.Repeat(" ", leftNumSpaces)
	rightSpaces := strings.Repeat(" ", rightNumSpaces)
	return leftSpaces + line + rightSpaces
}

// justify inserts additional spaces between words to fully justify the line,
// unless it is the last segment or no extra spaces are needed.
func justify(line string, spaces int, meta stringwrap.WrappedString) string {
	if meta.LastSegmentInOrig || spaces == 0 {
		return line
	}

	gaps := spaceGaps{}
	idx := 0
	for idx < len(line) {
		r, rSize, next, _ := ansiwalker.ANSIWalk(line, idx)
		idx = next - rSize

		if r == ' ' && !gaps.lastSpace {
			gaps.addGap(idx)
			gaps.lastSpace = true
		} else {
			gaps.lastSpace = false
		}
		idx = next
	}

	// based on the number of gaps and extra spaces, determine
	// the number of spaces per gap
	spacesPerGap := distributeSpaces(len(gaps.indices), spaces)

	// iterate through the line one again and compile a new string
	// where spaces are evenly distributed across all viewable gaps
	buf := strings.Builder{}
	curGapIdx := gaps.nextGap()
	for idx, r := range line {
		if idx == curGapIdx {
			spacesToAdd := spacesPerGap[gaps.gapIndex] + 1
			buf.WriteString(strings.Repeat(" ", spacesToAdd))
			gaps.clearGap()
			curGapIdx = gaps.nextGap()
		} else {
			buf.WriteRune(r)
		}
	}

	return buf.String()
}

// alignString wraps and then aligns each line of the input text according
// to the mode.
func alignString(str string, limit int, align string) (string, error) {
	stringAligner := alignFactory(align)
	strToAlign, metadata, err := stringwrap.StringWrap(str, limit, 4)
	metaLines := metadata.WrappedLines

	var alignedLines []string
	for idx, line := range strings.Split(strToAlign, "\n") {
		// right trim each line to ensure that there is no trailing
		// whitespace, since this is important to ensure that all text
		// is appropriately aligned.
		trimLine := strings.TrimRightFunc(line, unicode.IsSpace)

		// calculate the width after trimming in order to determine the
		// number of spaces from limit
		trimWidth := runewidth.StringWidth(trimLine)
		numSpaces := limit - trimWidth

		// call string aligner function and add to lines slice
		alignedLine := stringAligner(trimLine, numSpaces, metaLines[idx])
		alignedLines = append(alignedLines, alignedLine)
	}
	return strings.Join(alignedLines, "\n"), err
}

// LeftAlign wraps and left-aligns the input text within the given limit.
func LeftAlign(str string, limit int) (string, error) {
	return alignString(str, limit, "left")
}

// RightAlign wraps and right-aligns the input text within the given limit.
func RightAlign(str string, limit int) (string, error) {
	return alignString(str, limit, "right")
}

// CenterAlign wraps and center-aligns the input text within the given limit.
func CenterAlign(str string, limit int) (string, error) {
	return alignString(str, limit, "center")
}

// Justify wraps and justifies the input text within the given limit.
func Justify(str string, limit int) (string, error) {
	return alignString(str, limit, "justify")
}
