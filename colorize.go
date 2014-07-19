package chalk

import "fmt"

type Color struct {
	value int
}

type Style interface {
	Foreground(Color)
	Background(Color)
	Style(string) string
	WithBackground(Color) Style
	WithForeground(Color) Style
}

type style struct {
	foreground Color
	background Color
	// TODO(ttacon): add styles at some point (when we care enough about them)
}

func (c Color) Value() int {
	return c.value
}

func (c Color) Color(val string) string {
	return fmt.Sprintf("%s%s%s", c, val, ResetColor)
}

func (c Color) String() string {
	return fmt.Sprintf("\u001b[%dm", 30+c.value)
}

func (c Color) NewStyle() Style {
	return &style{foreground: c}
}

func (s *style) WithBackground(col Color) Style {
	s.Background(col)
	return s
}

func (s *style) WithForeground(col Color) Style {
	s.Foreground(col)
	return s
}

func (s *style) String() string {
	var toReturn string
	toReturn = fmt.Sprintf("\u001b[%dm", 40+s.background.Value())
	return toReturn + fmt.Sprintf("\u001b[%dm", 30+s.foreground.Value())
}

func (s *style) Style(val string) string {
	return fmt.Sprintf("%s%s%s", s, val, Reset)
}

func (s *style) Foreground(col Color) {
	s.foreground = col
}

func (s *style) Background(col Color) {
	s.background = col
}

var (
	nine = 9

	Black      = Color{0}
	Red        = Color{1}
	Green      = Color{2}
	Yellow     = Color{3}
	Blue       = Color{4}
	Magenta    = Color{5}
	Cyan       = Color{6}
	White      = Color{7}
	ResetColor = Color{9}

	Reset = &style{ResetColor, ResetColor}
)
