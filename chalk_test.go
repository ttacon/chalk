package chalk

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	col := Color{1}
	if col.Value() != 1 {
		t.Errorf("col.Value() != 1, was %d", col.Value())
	}

	expected := "\u001b[31m"
	if col.String() != expected {
		t.Errorf("expected col.String() to be %q, was %q", expected, col.String())
	}

	expected = fmt.Sprintf("%shello%s", expected, ResetColor)
	if col.Color("hello") != expected {
		t.Errorf("expected col.Color() to be %q, was %q", expected, col.Color("hello"))
	}
}

func TestTextStyle(t *testing.T) {
	text := TextStyle{1, 22}

	expected := "\u001b[1m\u001b[22m"
	if text.String() != expected {
		t.Errorf("expected text.String() to be %q, was %q", expected, text.String())
	}

	expected = "\u001b[1mhello\u001b[22m"
	if text.TextStyle("hello") != expected {
		t.Errorf("expected text.TextStyle() to be %q, was %q", expected, text.String())
	}

	empty := TextStyle{}
	if empty.TextStyle("hello") != "hello" {
		t.Errorf("expected empty.TextStyle() to be %q, was %q",
			expected,
			empty.String())
	}
}

func TestStyle(t *testing.T) {
	colStyle := Red.NewStyle()
	expected := "\u001b[40m\u001b[31m"
	if colStyle.String() != expected {
		t.Errorf("expected colStyle.String() to be %q, was %q",
			expected,
			colStyle.String())
	}

	textStyle := Bold.NewStyle()
	expected = "\u001b[40m\u001b[30m\u001b[1mhello\u001b[22m\u001b[49m\u001b[39m"
	if textStyle.Style("hello") != expected {
		t.Errorf("expected textStyle.Style(\"hello\") to be %q, was %q",
			expected,
			textStyle.Style("hello"))
	}

	// reset it
	colStyle.Foreground(ResetColor)
	colStyle.Background(ResetColor)
	expected = "\u001b[49m\u001b[39m"
	if colStyle.String() != expected {
		t.Errorf("expected colStyle.String() to be %q, was %q",
			expected,
			colStyle.String())
	}

	builtStyle := colStyle.
		WithForeground(Red).
		WithBackground(Blue).
		WithTextStyle(Underline)
	expected = "\u001b[44m\u001b[31m\u001b[4mhello\u001b[24m\u001b[49m\u001b[39m"
	if builtStyle.Style("hello") != expected {
		t.Errorf("expected builtStyle.Style() to be %q, was %q",
			expected,
			builtStyle.Style("hello"))
	}
}
