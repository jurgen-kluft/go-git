	"github.com/go-git/go-git/v5/plumbing"
	addLine    = "%s+%s%s%s"
	deleteLine = "%s-%s%s%s"
	equalLine  = "%s %s%s%s"
	// colorConfig is the color configuration. The default is no color.
	color ColorConfig

// SetColor sets e's color configuration and returns e.
func (e *UnifiedEncoder) SetColor(colorConfig ColorConfig) *UnifiedEncoder {
	e.color = colorConfig
	return e
}

			c.WriteTo(&e.buf, e.color)
		e.buf.WriteString(e.color[Meta])


		e.buf.WriteString(e.color.Reset())
		e.buf.WriteString(e.color[Meta])
		e.buf.WriteString(e.color.Reset())
		e.buf.WriteString(e.color[Meta])
		e.buf.WriteString(e.color.Reset())
func (c *hunk) WriteTo(buf *bytes.Buffer, color ColorConfig) {
	buf.WriteString(color[Frag])
	buf.WriteString(color.Reset())
		buf.WriteString(d.String(color))
func (o *op) String(color ColorConfig) string {
	var setColor, prefix, suffix string
		setColor = color[New]
		setColor = color[Old]
		setColor = color[Context]
	return fmt.Sprintf(prefix, setColor, o.text, color.Reset(), suffix)