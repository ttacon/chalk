package colorize

import "fmt"

type Color interface {
	Foreground(*int) int
	Background(*int) int
	Color(string) string
	WithBackground(Color) Color
}

// TODO(ttacon): add styles at some point (when we care enough about them)

type color struct {
	value      int
	background *int
}

func (c *color) WithBackground(col Color) Color {
	bg := col.Foreground(nil)
	fmt.Println(bg)
	c.Background(&bg)
	return c
}

func (c *color) String() string {
	var toReturn string
	if c.background != nil {
		toReturn = fmt.Sprintf("\u001b[%dm", 40+*c.background)
	}
	return toReturn + fmt.Sprintf("\u001b[%dm", 30+c.value)
}

func (c *color) Color(val string) string {
	return fmt.Sprintf("%s%s%s", c, val, Reset)
}

func (c *color) Foreground(val *int) int {
	if val != nil {
		c.value = *val
	}
	return c.value
}

func (c *color) Background(val *int) int {
	if val != nil {
		temp := *val
		c.background = &temp
	}
	return *c.background
}

var (
	nine = 9

	Black   = &color{0, nil}
	Red     = &color{1, nil}
	Green   = &color{2, nil}
	Yellow  = &color{3, nil}
	Blue    = &color{4, nil}
	Magenta = &color{5, nil}
	Cyan    = &color{6, nil}
	White   = &color{7, nil}
	Reset   = &color{9, &nine}
)
