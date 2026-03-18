package faker

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"math/rand"
	"strings"
)

//go:embed data/first_names.csv
var firstNamesData string

//go:embed data/last_names.csv
var lastNamesData string

//go:embed data/domains.csv
var domainsData string

var (
	firstNames []string
	lastNames  []string
	domains    []string
)

func init() {
	firstNames = parseCSV(firstNamesData)
	lastNames = parseCSV(lastNamesData)
	domains = parseCSV(domainsData)
}

func parseCSV(data string) []string {
	reader := csv.NewReader(strings.NewReader(data))
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

// GenerateFirstName returns a random first name.
func GenerateFirstName() string {
	return getRandomEl(firstNames)
}

// GenerateLastName returns a random last name.
func GenerateLastName() string {
	return getRandomEl(lastNames)
}

// GenerateFullName returns a random full name.
func GenerateFullName() string {
	return fmt.Sprintf("%s %s", GenerateFirstName(), GenerateLastName())
}

// GenerateDomain returns a random domain
func GenerateDomain() string {
	return getRandomEl(domains)
}

// GenerateEmailAddress returns a random email.
func GenerateEmailAddress() string {
	firstName := GenerateFirstName()
	lastName := GenerateLastName()
	domain := GenerateDomain()
	return fmt.Sprintf("%s.%s@%s", firstName, lastName, domain)
}

// GenerateEmailAddressByFullName returns an email with a random domain.
func GenerateEmailAddressByFullName(firstName, lastName string) string {
	domain := GenerateDomain()
	return fmt.Sprintf("%s.%s@%s", firstName, lastName, domain)
}

func getRandomEl(el []string) string {
	return el[rand.Intn(len(el))]
}
