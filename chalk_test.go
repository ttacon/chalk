package chalk

import (
	"fmt"
	"os"
	"sync"
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

func TestSetEnabled(t *testing.T) {
	origEnabled := Enabled()
	defer SetEnabled(origEnabled)

	SetEnabled(false)
	if Enabled() {
		t.Errorf("expected Enabled() to be false after SetEnabled(false)")
	}

	// Colors should produce plain text or empty strings
	if got := Red.Color("test string"); got != "test string" {
		t.Errorf("expected Red.Color(\"test string\") to be \"test string\", was %q", got)
	}
	if got := Red.String(); got != "" {
		t.Errorf("expected Red.String() to be \"\", was %q", got)
	}
	if got := ResetColor.String(); got != "" {
		t.Errorf("expected ResetColor.String() to be \"\", was %q", got)
	}

	// TextStyles should produce plain text or empty strings
	if got := Bold.TextStyle("bold text"); got != "bold text" {
		t.Errorf("expected Bold.TextStyle(\"bold text\") to be \"bold text\", was %q", got)
	}
	if got := Bold.String(); got != "" {
		t.Errorf("expected Bold.String() to be \"\", was %q", got)
	}

	// Styles should produce plain text or empty strings
	st := Red.NewStyle().WithBackground(Blue).WithTextStyle(Underline)
	if got := st.Style("styled text"); got != "styled text" {
		t.Errorf("expected st.Style(\"styled text\") to be \"styled text\", was %q", got)
	}
	if got := st.String(); got != "" {
		t.Errorf("expected st.String() to be \"\", was %q", got)
	}
	if got := Reset.String(); got != "" {
		t.Errorf("expected Reset.String() to be \"\", was %q", got)
	}
	if got := Reset.Style("text"); got != "text" {
		t.Errorf("expected Reset.Style(\"text\") to be \"text\", was %q", got)
	}

	// Empty string handling when disabled
	if got := Red.Color(""); got != "" {
		t.Errorf("expected Red.Color(\"\") to be \"\", was %q", got)
	}
	if got := st.Style(""); got != "" {
		t.Errorf("expected st.Style(\"\") to be \"\", was %q", got)
	}

	// Re-enable and verify ANSI formatting is restored
	SetEnabled(true)
	if !Enabled() {
		t.Errorf("expected Enabled() to be true after SetEnabled(true)")
	}

	expectedColor := fmt.Sprintf("%shello%s", Red, ResetColor)
	if got := Red.Color("hello"); got != expectedColor {
		t.Errorf("expected Red.Color(\"hello\") to be %q, was %q", expectedColor, got)
	}
	if got := Red.String(); got != "\u001b[31m" {
		t.Errorf("expected Red.String() to be %q, was %q", "\u001b[31m", got)
	}
	if got := st.Style("hello"); got != "\u001b[44m\u001b[31m\u001b[4mhello\u001b[24m\u001b[49m\u001b[39m" {
		t.Errorf("expected st.Style(\"hello\") to contain ANSI escapes, was %q", got)
	}
}

func TestCheckNoColor(t *testing.T) {
	origEnabled := Enabled()
	defer SetEnabled(origEnabled)

	origEnv, hadEnv := os.LookupEnv("NO_COLOR")
	defer func() {
		if hadEnv {
			os.Setenv("NO_COLOR", origEnv)
		} else {
			os.Unsetenv("NO_COLOR")
		}
	}()

	// When NO_COLOR is not set
	os.Unsetenv("NO_COLOR")
	SetEnabled(true)
	if CheckNoColor() {
		t.Errorf("expected CheckNoColor() to return false when NO_COLOR is unset")
	}
	if !Enabled() {
		t.Errorf("expected Enabled() to remain true when NO_COLOR is unset")
	}

	// When NO_COLOR is set to an empty string (standard specifies non-empty)
	os.Setenv("NO_COLOR", "")
	SetEnabled(true)
	if CheckNoColor() {
		t.Errorf("expected CheckNoColor() to return false when NO_COLOR is empty")
	}
	if !Enabled() {
		t.Errorf("expected Enabled() to remain true when NO_COLOR is empty")
	}

	// When NO_COLOR is set to non-empty
	os.Setenv("NO_COLOR", "1")
	SetEnabled(true)
	if !CheckNoColor() {
		t.Errorf("expected CheckNoColor() to return true when NO_COLOR is set")
	}
	if Enabled() {
		t.Errorf("expected Enabled() to be false when NO_COLOR is set")
	}
	if got := Red.Color("plain text"); got != "plain text" {
		t.Errorf("expected Red.Color(\"plain text\") to be \"plain text\", was %q", got)
	}
}

func TestConcurrentEnabled(t *testing.T) {
	origEnabled := Enabled()
	defer SetEnabled(origEnabled)

	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				SetEnabled(j%2 == 0)
				_ = Enabled()
				_ = Red.Color("concurrent")
				_ = Bold.TextStyle("concurrent")
				_ = Red.NewStyle().Style("concurrent")
			}
		}(i)
	}
	wg.Wait()
}
