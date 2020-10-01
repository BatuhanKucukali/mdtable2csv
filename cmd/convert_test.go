package cmd

import (
	"strings"
	"testing"
)

var (
	markdown = `| Tables        | Are           | Cool  |
				| ------------- |:-------------:| -----:|
				| col 3 is      | right-aligned | $1600 |
				| col 2 is      | centered      |   $12 |
				| zebra stripes | are neat      |    $1 |`

	markdown2 = `Tables        | Are           | Cool  
				------------- |:--------------:| -----:
				col 3 is      | right-aligned  | $1600 
				col 2 is      | centered       |   $12 
				zebra stripes | are neat       |    $1 `
)

func TestGetColumns(t *testing.T) {
	result := getColumns("col 3 is | right-aligned | $1600")

	if len(result) != 3 {
		t.Errorf("Expected 3 but actual %d", len(result))
	}
}

func TestGetCleanColumns(t *testing.T) {
	columns := []string{" col 3 is", " right-aligned    ", "   $1600"}
	result := cleanColumns(columns)

	if len(result) != 3 {
		t.Errorf("Expected 3 but actual %d", len(result))
	}

	if result[0] != "col 3 is" {
		t.Errorf("Expected 'col 3 is' but actual '%s'", result[0])
	}

	if result[1] != "right-aligned" {
		t.Errorf("Expected 'right-aligned but actual '%s'", result[1])
	}

	if result[2] != "$1600" {
		t.Errorf("Expected '$1600' but actual '%s'", result[2])
	}

}

func TestClineLineWhenLineStartWithPipe(t *testing.T) {
	result := cleanLine("|col 3 is | $1600|")

	if result != "col 3 is | $1600" {
		t.Errorf("Expected 'col 3 is | $1600' but actual '%s'", result)
	}
}

func TestCleanLine(t *testing.T) {
	result := cleanLine("col 3 is | $1600")

	if result != "col 3 is | $1600" {
		t.Errorf("Expected 'col 3 is | $1600' but actual '%s'", result)
	}
}

func TestGetCsvDataWhenLineStartWithPipe(t *testing.T) {
	result, err := getCsvData(strings.NewReader(markdown))

	if err != nil {
		t.Errorf("Error = %v", err)
	}

	if len(result) != 4 {
		t.Errorf("Expected 3 but actual %d", len(result))
	}
}

func TestGetCsvDataWhenLineStartWithoutPipe(t *testing.T) {
	result, err := getCsvData(strings.NewReader(markdown2))

	if err != nil {
		t.Errorf("Error = %v", err)
	}

	if len(result) != 4 {
		t.Errorf("Expected 3 but actual %d", len(result))
	}

}
