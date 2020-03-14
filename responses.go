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

// TransferResponse is the standard response to transfer operations.
type TransferResponse struct {
	UserID         int         `json:"user_id"`
	AccountID      interface{} `json:"account_id"`
	BeneficiaryID  int         `json:"beneficiary_id"`
	SourceAmount   string      `json:"source_amount"`
	ExchangeRate   string      `json:"exchange_rate"`
	Reference      string      `json:"reference"`
	Fee            string      `json:"fee"`
	CurrencyPairs  string      `json:"currency_pairs"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      interface{} `json:"updated_at"`
	UUID           string      `json:"uuid"`
	State          string      `json:"state"`
	AuthorizingIP  string      `json:"authorizing_ip"`
	TransferStates []struct {
		State     string    `json:"state"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"transfer_states"`
	SourceCurrency      string `json:"source_currency"`
	DestinationAmount   string `json:"destination_amount"`
	DestinationCurrency string `json:"destination_currency"`
	Deposit             bool   `json:"deposit"`
}

// TransferStatus is the status response to get transfer status.
type TransferStatus struct {
	State     string `json:"state"`
	CreatedAt string `json:"created_at"`
}

// BatchTransferStatus is the standard batch transfer status response.
type BatchTransferStatus struct {
	UUID      string `json:"uuid"`
	QuoteUUID string `json:"quote_uuid"`
	Status    string `json:"status"`
}

// TransactionResponse is a representation of data about transactions.
type TransactionResponse struct {
	UUID                string      `json:"uuid"`
	UserID              int         `json:"user_id"`
	AccountID           interface{} `json:"account_id"`
	BeneficiaryID       int         `json:"beneficiary_id"`
	CurrencyPairs       string      `json:"currency_pairs"`
	SourceCurrency      string      `json:"source_currency"`
	SourceAmount        string      `json:"source_amount"`
	DestinationAmount   string      `json:"destination_amount"`
	DestinationCurrency string      `json:"destination_currency"`
	ExchangeRate        string      `json:"exchange_rate"`
	AuthorizingIP       string      `json:"authorizing_ip"`
	State               string      `json:"state"`
	TransferStates      []struct {
		State     string    `json:"state"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"transfer_states"`
	CreatedAt time.Time `json:"created_at"`
}

// WebhookResponse is the representation of response data for webhook based
// operations.
type WebhookResponse struct {
	UUID        string    `json:"uuid"`
	URL         string    `json:"url"`
	Type        string    `json:"type"`
	Rfuuid      string    `json:"rfuuid"`
	FailedCount int       `json:"failed_count"`
	RetryCount  int       `json:"retry_count"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

// BalanceResponse is a representation of current balance.
type BalanceResponse struct {
	Currency string
	Balance  float64
}

// KYCDetails is a representation of details retained by KYC.
type KYCDetails struct {
	AgreedToTerms          bool   `json:"agreedToTerms"`
	AllowAccountManagement bool   `json:"allowAccountManagement"`
	Address                string `json:"address"`
	City                   string `json:"city"`
	CompanyName            string `json:"companyName"`
	Country                string `json:"country"`
	DateOfIncorporation    string `json:"dateOfIncorporation"`
	Dba                    bool   `json:"dba"`
	DbaName                string `json:"dbaName"`
	IncorporationNumber    string `json:"incorporationNumber"`
	Officers               []struct {
		Address     string `json:"address"`
		Citizenship string `json:"citizenship"`
		City        string `json:"city"`
		Dob         string `json:"dob"`
		FirstName   string `json:"firstName"`
		IDNumber    string `json:"idNumber"`
		IDType      string `json:"idType"`
		JobTitle    string `json:"jobTitle"`
		LastName    string `json:"lastName"`
		Owner       bool   `json:"owner"`
		Ownership   string `json:"ownership"`
		PostalCode  string `json:"postalCode"`
		State       string `json:"state"`
		Title       string `json:"title"`
	} `json:"officers"`
	Owners []struct {
		Address     string `json:"address"`
		Citizenship string `json:"citizenship"`
		City        string `json:"city"`
		Dob         string `json:"dob"`
		FirstName   string `json:"firstName"`
		IDNumber    string `json:"idNumber"`
		IDType      string `json:"idType"`
		JobTitle    string `json:"jobTitle"`
		LastName    string `json:"lastName"`
		Owner       bool   `json:"owner"`
		Ownership   string `json:"ownership"`
		PostalCode  string `json:"postalCode"`
		State       string `json:"state"`
		Title       string `json:"title"`
	} `json:"owners"`
	Payments struct {
		Countries []string `json:"countries"`
		Frequency string   `json:"frequency"`
		Purpose   string   `json:"purpose"`
		Volume    string   `json:"volume"`
	} `json:"payments"`
	Phone      string `json:"phone"`
	PostalCode string `json:"postalCode"`
	State      string `json:"state"`
	Structure  string `json:"structure"`
	Website    string `json:"website"`
}

// PaymentInstructions is a representation of payment instructions.
type PaymentInstructions struct {
	Currency            string `json:"Currency"`
	PaymentInstructions string `json:"PaymentInstructions"`
}
