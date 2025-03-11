package table

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIndent(t *testing.T) {
	tests := []struct {
		name       string
		textString string
		indent     string
		expected   string
	}{
		{
			"space",
			"This is a test string",
			" ",
			" This is a test string",
		},
		{
			"space with new line at first",
			"\nThis is a test string",
			" ",
			" \n This is a test string",
		},
		{
			"space with new line at last",
			"This is a test string\n",
			" ",
			" This is a test string\n ",
		},
		{
			"blank string",
			"",
			" ",
			"",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Indent(test.textString, test.indent)
			assert.Equal(t, test.expected, result)
		})
	}
}
func TestDisplayWidth(t *testing.T) {
	if DisplayWidth("test") != 4 {
		t.Errorf("Error in DisplayWidth")
	}
	if DisplayWidth(" test") != 5 {
		t.Errorf("Error in DisplayWidth")
	}
	if DisplayWidth(" test ") != 6 {
		t.Errorf("Error in DisplayWidth")
	}
}
func TestTruncate(t *testing.T) {
	tests := []struct {
		name        string
		testString  string
		manualWidth int
		expected    string
	}{
		{
			"const minWidthForEllipsis minus 1",
			strings.Repeat("a", minWidthForEllipsis-1),
			minWidthForEllipsis,
			strings.Repeat("a", minWidthForEllipsis-1),
		},
		{
			"custom width",
			"abcdefghijklmnopqrstuvwxyz",
			10,
			"abcdefg...",
		},
		{
			"const minWidthForEllipsis width",
			strings.Repeat("a", minWidthForEllipsis),
			minWidthForEllipsis,
			strings.Repeat("a", minWidthForEllipsis),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Truncate(test.manualWidth, test.testString)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestPluralize(t *testing.T) {
	tests := []struct {
		name       string
		testString string
		times      int
		expected   string
	}{
		{
			"single thing",
			"test",
			1,
			"1 test",
		},
		{
			"2 things",
			"test",
			2,
			"2 tests",
		},
		{
			"0 thing",
			"test",
			0,
			"0 tests",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Pluralize(test.times, test.testString)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestRelativeTimeAgo(t *testing.T) {
	tests := []struct {
		name            string
		checktimeFirst  time.Time
		checktimeSecond time.Time
		expected        string
	}{
		{
			"less than a minute ago",
			time.Now(),
			time.Now().Add(-50 * time.Second),
			"less than a minute ago",
		},
		{
			"minutes ago",
			time.Now(),
			time.Now().Add(-50 * time.Minute),
			"minutes ago",
		},
		{
			"about 1 hour ago",
			time.Now(),
			time.Now().Add(-61 * time.Minute),
			"about 1 hour ago",
		},
		{
			"about 1 day ago",
			time.Now(),
			time.Now().Add(-25 * time.Hour),
			"about 1 day ago",
		},
		{
			"about 2 months ago",
			time.Now(),
			time.Now().Add(-1460 * time.Hour),
			"about 2 months ago",
		},
		{
			"about 1 year ago",
			time.Now(),
			time.Now().Add(-8761 * time.Hour),
			"about 1 year ago",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RelativeTimeAgo(test.checktimeFirst, test.checktimeSecond)
			assert.Contains(t, result, test.expected)
		})
	}
}
