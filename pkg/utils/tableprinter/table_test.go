package table

import (
	"strings"
	"testing"

	"github.com/berrybytes/zocli/pkg/utils/factory"
	mock_factory "github.com/berrybytes/zocli/pkg/utils/factory/mock"
	"github.com/berrybytes/zocli/pkg/utils/iostreams"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		factory *factory.Factory
		width   int
	}{
		{
			"width value zero",
			mock_factory.NewFactory(),
			0,
		},
		{
			"width value random",
			mock_factory.NewFactory(),
			100,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tp := New(test.factory, test.width)

			assert.NotNil(t, tp, "Table Printer must not be nil")
		})
	}
}

func TestAddField(t *testing.T) {
	f := mock_factory.NewFactory()
	ttPrinter := New(f, 0).(*ttyTablePrinter)

	ttPrinter.AddField("TestField")
	assert.Equal(t,
		"TestField",
		ttPrinter.rows[0][0].text,
		" Expected first field to be 'TestField'",
	)

	ttPrinter.AddField("TestField2")
	assert.Equal(t,
		"TestField2",
		ttPrinter.rows[0][1].text,
		" Expected second field to be 'TestField2'",
	)
}

func TestEndRow(t *testing.T) {
	f := mock_factory.NewFactory()
	ttPrinter := New(f, 80).(*ttyTablePrinter)

	ttPrinter.AddField("TestField")
	ttPrinter.AddField("TestField2")

	ttPrinter.EndRow()

	ttPrinter.AddField("TestField3")

	assert.Equal(t,
		2,
		len(ttPrinter.rows),
		"Expected table to have 2 rows",
	)
	assert.Equal(t,
		"TestField3",
		ttPrinter.rows[1][0].text,
		"Expected first field of second row to be 'TestField'",
	)
}

func TestHeaderRow(t *testing.T) {
	f := mock_factory.NewFactory()
	ttPrinter := New(f, 80).(*ttyTablePrinter)

	ttPrinter.HeaderRow(strings.ToUpper, "header1", "header2")

	assert.Equal(t,
		"HEADER1",
		ttPrinter.rows[0][0].text,
		"Expected first column to be 'HEADER1'",
	)
	assert.Equal(t,
		"HEADER2",
		ttPrinter.rows[0][1].text,
		"Expected second column to be 'HEADER2'",
	)
}
func TestPrint(t *testing.T) {
	io, _, out, _ := iostreams.Test()
	f := mock_factory.NewFactory()
	f.IO = io
	ttprinter := New(f, 80).(*ttyTablePrinter)

	ttprinter.HeaderRow(strings.ToUpper, "column_1", "column_2")
	ttprinter.EndRow()
	ttprinter.AddField("Value_1")
	ttprinter.AddField("Value_2")
	ttprinter.EndRow()

	err := ttprinter.Print()

	assert.NoError(t, err, "No error expected on print")

	outString := out.String()
	expectedOut := "COLUMN_1  COLUMN_2\nValue_1   Value_2\n"

	assert.Equal(t, expectedOut, outString, "Expected print output does not match")
}
