package Utils

import (
	"fmt"
	"log"
)

// Contract templates (stored as constants, files, or in the database)
var englishTemplate string = `CONTRACT AGREEMENT

This agreement is made on {{date_created}} between:
- **Service Provider (Tradesman)**:
  Name: {{tradesman_first_name}} {{tradesman_last_name}}
  Phone Number: {{tradesman_phone}}
  Address: {{tradesman_location}} ({{tradesman_loc_details}})

- **Client**:
  Name: {{client_first_name}} {{client_last_name}}
  Phone Number: {{client_phone}}
  Address: {{client_location}} ({{client_loc_details}})

**Service Details**:
- Service Type: {{listing_type}} ({{listing_title}})
- Description: {{listing_description}}
- Location of Service: {{listing_city}}, {{listing_country}}

**Transaction Terms**:
- Price: {{transaction_price}} {{transaction_currency}}
- Job Start Date: {{job_start_date}}
- Job End Date: {{job_end_date}}
- Details from Tradesman: {{details_from_offering}}
- Details from Client: {{details_from_offered}}

**Acknowledgments**:
Both parties agree to the terms outlined in this document. Disputes will be resolved as per the local laws of Lebanon.

Signature of Tradesman: ______________________
Signature of Client: ______________________
` // English contract template
var arabicTemplate string = `...` // Arabic contract template

// Fetch data from the database (using the query above)
var contractData = ContractData{
	TradesmanFirstName:  "John",
	TradesmanLastName:   "Doe",
	TradesmanPhone:      "+961-12345678",
	TradesmanLocation:   "Beirut",
	TradesmanLocDetails: "Hamra Street",
	ClientFirstName:     "Jane",
	ClientLastName:      "Smith",
	ClientPhone:         "+961-87654321",
	ClientLocation:      "Tripoli",
	ClientLocDetails:    "Mina",
	ListingType:         "Offer",
	ListingTitle:        "Plumbing Services",
	ListingDescription:  "Fixing water pipes and leaks.",
	ListingCity:         "Beirut",
	ListingCountry:      "Lebanon",
	TransactionPrice:    200,
	TransactionCurrency: "USD",
	JobStartDate:        "2024-12-20",
	JobEndDate:          "2024-12-25",
	DateCreated:         "2024-12-15",
	DetailsFromOffering: "Requires tools to be provided.",
	DetailsFromOffered:  "Ensure no damage to tiles.",
}

func main() {
	// Generate the contract
	finalContract, err := GenerateContract(englishTemplate, contractData)
	if err != nil {
		log.Fatalf("Error generating contract: %v", err)
	}

	fmt.Println(finalContract)
}
