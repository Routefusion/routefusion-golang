package routefusion

import "time"

// UserDetails is a representation of the base struct that user based route
// options respond with.
type UserDetails struct {
	UUID                  string    `json:"uuid"`
	FirstName             string    `json:"first_name"`
	LastName              string    `json:"last_name"`
	Email                 string    `json:"email"`
	PhoneNumber           string    `json:"phone_number"`
	Country               string    `json:"country"`
	Verified              bool      `json:"verified"`
	Type                  string    `json:"type"`
	VerificationSubmitted bool      `json:"verification_submitted"`
	CompanyName           string    `json:"company_name"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
	MasterUser            bool      `json:"master_user"`
}

// UpdatedUserDetails are a representation of the values sent after a user is
// updated.
type UpdatedUserDetails struct {
	UserDetails
	MasterUserUUID string `json:"master_user_uuid"`
}

// AllUserDetails are the details received by the master call.
type AllUserDetails struct {
	UserDetails
	CompanyName string `json:"company_name"`
	City        string `json:"city"`
	Street      string `json:"street"`
	PostalCode  string `json:"postal_code"`
	Admin       bool   `json:"admin"`
}

// A BeneficiaryBase is a representation of someone you can send money to.
type BeneficiaryBase struct {
	ID                 int         `json:"id"`
	UUID               string      `json:"uuid"`
	UserID             int         `json:"user_id"`
	CompanyName        string      `json:"company_name"`
	FirstNameOnAccount string      `json:"first_name_on_account"`
	LastNameOnAccount  string      `json:"last_name_on_account"`
	Type               string      `json:"type"`
	BankName           string      `json:"bank_name"`
	BranchName         interface{} `json:"branch_name"`
	BankCity           string      `json:"bank_city"`
	BankCode           interface{} `json:"bank_code"`
	BranchCode         interface{} `json:"branch_code"`
	AccountType        string      `json:"account_type"`
	AccountNumber      string      `json:"account_number"`
	RoutingNumber      string      `json:"routing_number"`
	Clabe              interface{} `json:"clabe"`
	TaxNumber          interface{} `json:"tax_number"`
	Email              string      `json:"email"`
	PhoneNumber        interface{} `json:"phone_number"`
	Country            string      `json:"country"`
	City               string      `json:"city"`
	BankStateProvince  string      `json:"bank_state_province"`
	Verified           bool        `json:"verified"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	Currency           string      `json:"currency"`
	Cpfcnpj            interface{} `json:"cpfcnpj"`
	SwiftBic           string      `json:"swift_bic"`
	BankAddress1       string      `json:"bank_address1"`
	BankAddress2       interface{} `json:"bank_address2"`
	BankCountry        string      `json:"bank_country"`
	BankPostalCode     string      `json:"bank_postal_code"`
	Address1           string      `json:"address1"`
	Address2           interface{} `json:"address2"`
	StateProvince      string      `json:"state_province"`
	PostalCode         string      `json:"postal_code"`
	BsbNumber          interface{} `json:"bsb_number"`
}

// Beneficiary is a representation of a single complete Beneficiary.
type Beneficiary struct {
	BeneficiaryBase
	StatusHistory []struct {
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"status_history"`
	Status string `json:"status"`
}

// QuoteResponse is the standard response for a created quote.
type QuoteResponse struct {
	UUID                string    `json:"uuid"`
	SourceCurrency      string    `json:"source_currency"`
	DestinationCurrency string    `json:"destination_currency"`
	Rate                string    `json:"rate"` // why not float?
	InvertedRate        string    `json:"inverted_rate"`
	DateOfPayment       time.Time `json:"date_of_payment"`
	ExpiresAt           time.Time `json:"expires_at"`
	CreatedAt           time.Time `json:"created_at"`
}
