package cmd

import (
	"bufio"
	"encoding/csv"
	"github.com/spf13/cobra"
	"io"
	"os"
	"strings"
)

var (
	defaultDelimiter  = ","
	markdownDelimiter = "|"
)

var convertCmd = &cobra.Command{
	Use:   "convert example.md",
	Short: "Convert markdown table to csv file",
	Args:  cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath := args[0]
		delimiter := defaultDelimiter

		if len(args) == 2 {
			delimiter = args[1]
		}

		if err := convert(filePath, delimiter); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(convertCmd)
}

func convert(filePath string, delimiter string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	csvData, err := getCsvData(file)
	if err != nil {
		return err
	}

	return createCsvFile(file.Name(), delimiter, csvData)
}

func createCsvFile(fileName string, delimiter string, csvData [][]string) error {
	csvFile, err := os.Create(fileName + ".csv")
	if err != nil {
		return err
	}

	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	writer.Comma = rune(delimiter[0])
	defer writer.Flush()

	for _, value := range csvData {
		err := writer.Write(value)
		if err != nil {
			return err
		}
	}
	return nil
}

func getCsvData(r io.Reader) ([][]string, error) {
	var csvData [][]string

	scanner := bufio.NewScanner(r)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i == 1 || len(line) == 0 {
			i++
			continue
		}

		line = cleanLine(line)
		columns := getColumns(line)
		newRow := cleanColumns(columns)

		csvData = append(csvData, newRow)
		i++
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return csvData, nil
}

func cleanColumns(columns []string) []string {
	var newRows []string
	for i := 0; i < len(columns); i++ {
		newRow := strings.TrimSpace(columns[i])
		newRows = append(newRows, newRow)
	}
	return newRows
}

func getColumns(line string) []string {
	return strings.Split(line, "|")
}

func cleanLine(line string) string {
	firstCharacter := line[:1]
	lastCharacter := line[len(line)-1:]

	if firstCharacter == markdownDelimiter {
		line = line[1:]
	}

	if lastCharacter == markdownDelimiter {
		return line[:len(line)-1]
	}

	return line
}
