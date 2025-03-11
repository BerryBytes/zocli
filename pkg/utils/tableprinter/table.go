package table

import (
	"fmt"
	"io"
	"strings"

	"github.com/berrybytes/zocli/pkg/utils/factory"
)

type ttyTablePrinter struct {
	out        io.Writer
	f          *factory.Factory
	WithHeader bool

	rows     [][]tablefield
	maxwidth int
}

type TablePrinter interface {
	AddField(string, ...fieldOptions)
	Print() error
	HeaderRow(func(string) string, ...string)
	EndRow()
	Separator()
}

type tablefield struct {
	text      string
	width     int
	colorFunc func(string) string
}

type TableType string

func New(f *factory.Factory, maxwidth int) TablePrinter {
	w := f.IO.Out
	if maxwidth == 0 {
		maxwidth = f.IO.TerminalWidth()
	}
	opts := ttyTablePrinter{
		out:      w,
		f:        f,
		maxwidth: maxwidth,
	}

	return &opts
}

func (t *ttyTablePrinter) Separator() {
	fmt.Println("--------------------")
}

func (t *ttyTablePrinter) EndRow() {
	t.rows = append(t.rows, []tablefield{})
}

func (t *ttyTablePrinter) HeaderRow(color func(string) string, columns ...string) {
	for _, col := range columns {
		opts := func(f *tablefield) {
			f.colorFunc = color
		}
		t.AddField(strings.ToUpper(col), opts)
	}
	t.EndRow()
}

type fieldOptions func(*tablefield)

func (t *ttyTablePrinter) AddField(s string, opts ...fieldOptions) {
	if t.rows == nil {
		t.rows = make([][]tablefield, 1)
	}

	rowNo := len(t.rows) - 1

	field := tablefield{
		text:  s,
		width: len(s),
	}

	for _, opt := range opts {
		opt(&field)
	}

	t.rows[rowNo] = append(t.rows[rowNo], field)
}

func (t *ttyTablePrinter) Print() error {
	if len(t.rows) == 0 {
		return nil
	}

	delim := "  "
	numCols := len(t.rows[0])
	colWidths := t.calculateColumnWidths(len(delim))

	for _, row := range t.rows {
		for col, field := range row {
			if col > 0 {
				_, err := fmt.Fprint(t.out, delim)
				if err != nil {
					return err
				}
			}
			truncVal := field.text
			if col < numCols-1 {
				// pad value with spaces on the right
				if padWidth := colWidths[col] - DisplayWidth(field.text); padWidth > 0 {
					truncVal += strings.Repeat(" ", padWidth)
				}
			}
			if field.colorFunc != nil {
				truncVal = field.colorFunc(truncVal)
			}
			_, err := fmt.Fprint(t.out, truncVal)
			if err != nil {
				return err
			}
		}
		if len(row) > 0 {
			_, err := fmt.Fprint(t.out, "\n")
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *ttyTablePrinter) calculateColumnWidths(delimSize int) []int {
	numCols := len(t.rows[0])
	maxColWidths := make([]int, numCols)
	colWidths := make([]int, numCols)

	for _, row := range t.rows {
		for col, field := range row {
			w := DisplayWidth(field.text)
			if w > maxColWidths[col] {
				maxColWidths[col] = w
			}
		}
	}

	availWidth := func() int {
		setWidths := 0
		for col := 0; col < numCols; col++ {
			setWidths += colWidths[col]
		}
		return t.maxwidth - delimSize*(numCols-1) - setWidths
	}
	numFixedCols := func() int {
		fixedCols := 0
		for col := 0; col < numCols; col++ {
			if colWidths[col] > 0 {
				fixedCols++
			}
		}
		return fixedCols
	}

	// set the widths of short columns
	if w := availWidth(); w > 0 {
		if numFlexColumns := numCols - numFixedCols(); numFlexColumns > 0 {
			perColumn := w / numFlexColumns
			for col := 0; col < numCols; col++ {
				if max := maxColWidths[col]; max < perColumn {
					colWidths[col] = max
				}
			}
		}
	}

	// truncate long columns to the remaining available width
	if numFlexColumns := numCols - numFixedCols(); numFlexColumns > 0 {
		perColumn := availWidth() / numFlexColumns
		for col := 0; col < numCols; col++ {
			if colWidths[col] == 0 {
				if max := maxColWidths[col]; max < perColumn {
					colWidths[col] = max
				} else if perColumn > 0 {
					colWidths[col] = perColumn
				}
			}
		}
	}

	// add the remainder to truncated columns
	if w := availWidth(); w > 0 {
		for col := 0; col < numCols; col++ {
			d := maxColWidths[col] - colWidths[col]
			toAdd := w
			if d < toAdd {
				toAdd = d
			}
			colWidths[col] += toAdd
			w -= toAdd
			if w <= 0 {
				break
			}
		}
	}

	return colWidths
}
