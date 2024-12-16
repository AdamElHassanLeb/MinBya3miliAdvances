package Utils

import (
	"bytes"
	"text/template"
)

type ContractData struct {
	TradesmanFirstName  string
	TradesmanLastName   string
	TradesmanPhone      string
	TradesmanLocation   string
	TradesmanLocDetails string
	ClientFirstName     string
	ClientLastName      string
	ClientPhone         string
	ClientLocation      string
	ClientLocDetails    string
	ListingType         string
	ListingTitle        string
	ListingDescription  string
	ListingCity         string
	ListingCountry      string
	TransactionPrice    float64
	TransactionCurrency string
	JobStartDate        string
	JobEndDate          string
	DateCreated         string
	DetailsFromOffering string
	DetailsFromOffered  string
}

func GenerateContract(templateStr string, contractData ContractData) (string, error) {
	tmpl, err := template.New("contract").Parse(templateStr)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, contractData)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
