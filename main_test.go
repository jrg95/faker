package faker

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func readCSVFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(fmt.Sprintf("failed to open CSV data: %v", err))
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("failed to close CSV data: %v\n", err)
		}
	}()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(fmt.Sprintf("failed to parse CSV data: %v", err))
	}

	var names []string
	for _, record := range records {
		names = append(names, record[0])
	}
	return names
}

func TestParseCSV(t *testing.T) {
	names := parseCSV("test")
	assert.Equal(t, []string{"test"}, names)
}

func TestParseCSVFirstNamesData(t *testing.T) {
	names := parseCSV(firstNamesData)
	assert.Equal(t, len(names), 662)
	assert.Equal(t, "Vanessa", names[0])
	assert.Equal(t, "Yvonne", names[len(names)-1])
}

// CSV has "Last Names" as the first name
func TestParseCSVLastNamesData(t *testing.T) {
	names := parseCSV(lastNamesData)
	assert.Equal(t, len(names), 990)
	assert.Equal(t, "Last Names", names[0])
	assert.Equal(t, "Fleming", names[len(names)-1])
}

func TestParseCSVFirstNamesMatch(t *testing.T) {
	names := readCSVFile("./data/first_names.csv")
	firstNames := parseCSV(firstNamesData)
	assert.Equal(t, names, firstNames)
}

func TestParseCSVLastNamesMatch(t *testing.T) {
	names := readCSVFile("./data/last_names.csv")
	lastNames := parseCSV(lastNamesData)
	assert.Equal(t, names, lastNames)
}

