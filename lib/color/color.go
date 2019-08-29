package color

import "fmt"

const (
	esc = "\x1b"
	endFormat = "m"

)

const (
	resetColor = iota
	bold
	faint
	italic
	underline
	slowBlink
	fastBlink
	reverse
	conceal
	strikeThrough
	resetFont
	fontAlt1
	fontAlt2
	fontAlt3
	fontAlt4
	fontAlt5
	fontAlt6
	fontAlt7
	fontAlt8
	fontAlt9
	fraktur
	boldOff
	resetIntensity
	resetItalic
	restUnderline
	resetBlink
	__unused1
	resetInvert
	resetConceal
	resetStrikeThrough

	// Normal FG Colors
	fgBlack
	fgRed
	fgGreen
	fgYellow
	fgBlue
	fgMagenta
	fgCyan
	fgWhite

	// Other FG Color

	fgCustom
	resetFg

	// Normal BG Colors

	bgBlack
	bgRed
	bgGreen
	bgYello
	bgBlue
	bgMagenta
	bgCyan
	bgWhite

	// Other BG Color

	bgCustom
	resetBg

	__unused2

	framed
	encircled
	overlined
	resetFrame
	resetOverline
)

// Ideograms
const (
	ideogramUnderline = 60 + iota
	ideogramDoubleUnderline
	ideogramOverline
	ideogramDoubleOverline
	ideogramStressMarking
	resetIdeogram
)

// Bright colors FG
const (
	fgBrightBlack = iota + 90
	fgBrightRed
	fgBrightGreen
	fgBrightYellow
	fgBrightBlue
	fgBrightMagenta
	fgBrightCyan
	fgBrightWhite
)

// Bright Colors BG
const (
	bgBrightBlack = iota + 100
	bgBrightRed
	bgBrightGreen
	bgBrightYellow
	bgBrightBlue
	bgBrightMagenta
	bgBrighCyan
	bgBrightWhite
)

// Misc Aliases
const (
	doubleUnderline = boldOff
	resetEncircle = resetFrame
	rightSideLine = ideogramUnderline
	rightSideDoubleLine = ideogramDoubleUnderline
	leftSideLine = ideogramOverline
	leftSideDoubleLine = ideogramDoubleOverline
)

func DarkRed(in string) string {
	return fgColor(in, fgRed)
}

func Red(in string) string {
	return fgColor(in, fgBrightRed)
}

func DarkCyan(in string) string {
	return fgColor(in, fgCyan)
}

func Cyan(in string) string {
	return fgColor(in, fgBrightCyan)
}

func Blue(in string) string {
	return fgColor(in, fgBrightBlue)
}

func DarkGray(in string) string {
	return fgColor(in, fgBrightBlack)
}

func Gray(in string) string {
	return fgColor(in, fgWhite)
}

func fgColor(in string, color int) string {
	return fmt.Sprintf("%s[%dm%s%s[%dm", esc, color, in, esc, resetFg)
}