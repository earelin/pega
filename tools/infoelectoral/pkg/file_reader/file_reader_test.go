/*
 * This program is free software: you can redistribute it and/or modify it under
 * the terms of the GNU General Public License as published by the Free Software
 * Foundation, either version 3 of the License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT ANY
 * WARRANTY; without even the implied warranty of MERCHANTABILITY or
 * FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License
 * for more details.
 *
 * You should have received a copy of the GNU General Public License along with
 * this program. If not, see <https://www.gnu.org/licenses/>.
 */

package file_reader

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	f, err := os.Open("testdata/columns_test.txt")
	if err != nil {
		t.Error(err)
	}

	var entries = []TestStruct{
		{
			Year:    2023,
			Month:   8,
			Day:     21,
			WeekDay: "Monday",
			Happy:   true,
		}, {
			Year:    2019,
			Month:   8,
			Day:     22,
			WeekDay: "Tuesday",
			Happy:   false,
		},
	}

	fileReader, err := NewFileReader[TestStruct](f)
	assert.Nil(t, err, "Should not have error on open file")

	for i := 0; i < 2; i++ {
		var entry, err = fileReader.Read()
		assert.Nil(t, err, "Should not have error on reading line")
		assert.Equal(t, entries[i], entry)
	}

	err = f.Close()
	if err != nil {
		t.Error(err)
	}
}

func Test_unMarshaling(t *testing.T) {
	tStruct, err := unMarshaling[TestStruct](firstLine, columnsContent)
	assert.Nil(t, err)
	assert.Equal(t, TestStruct{
		Year:    2023,
		Month:   8,
		Day:     21,
		WeekDay: "Monday",
		Happy:   true,
	}, tStruct)
}

func Test_extractColumns(t *testing.T) {
	var tStruct TestStruct

	columns, err := extractColumns(tStruct)
	assert.Nil(t, err, "Should not return error on valid columns")
	assert.ElementsMatch(t, columnsContent, columns)
}

func Test_calculateLineLength(t *testing.T) {
	var columns []Column

	assert.Equal(t, 0, calculateLineLength(columns))

	columns = []Column{
		{
			length: 2,
		},
		{
			length: 2,
		},
		{
			length: 4,
		},
	}
	assert.Equal(t, 8, calculateLineLength(columns))
}

func Test_sortColumns(t *testing.T) {
	columns := []Column{
		{
			position: 4,
		},
		{
			position: 0,
		},
		{
			position: 2,
		},
	}

	sortColumns(columns)

	assert.Equal(t, []Column{
		{
			position: 0,
		},
		{
			position: 2,
		},
		{
			position: 4,
		},
	}, columns)
}

var firstLine = []byte("20230821Monday  1")

type TestStruct struct {
	Year    int    `position:"0" length:"4"`
	Month   int    `position:"4" length:"2"`
	Day     int    `position:"6" length:"2"`
	WeekDay string `position:"8" length:"8"`
	Happy   bool   `position:"16" length:"1"`
}

var columnsContent = []Column{
	{
		name:       "Year",
		position:   0,
		length:     4,
		columnType: "int",
	},
	{
		name:       "Month",
		position:   4,
		length:     2,
		columnType: "int",
	},
	{
		name:       "Day",
		position:   6,
		length:     2,
		columnType: "int",
	},
	{
		name:       "WeekDay",
		position:   8,
		length:     8,
		columnType: "string",
	},
	{
		name:       "Happy",
		position:   16,
		length:     1,
		columnType: "bool",
	},
}
