	"regexp"
	addLine    = "+%s%s"
	deleteLine = "-%s%s"
	equalLine  = " %s%s"
	noNewLine  = "\n\\ No newline at end of file\n"
		message += "\n"
	c.current = &hunk{ctxPrefix: strings.TrimSuffix(ctxPrefix, "\n")}
var splitLinesRE = regexp.MustCompile(`[^\n]*(\n|$)`)

	out := splitLinesRE.FindAllString(s, -1)
	var prefix, suffix string
	n := len(o.text)
	if n > 0 && o.text[n-1] != '\n' {
		suffix = noNewLine
	}
	return fmt.Sprintf(prefix, o.text, suffix)