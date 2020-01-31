package routefusion

import "time"

// User represents the changeable details pertaining to a user.
// TODO: QUESTION- Is this a multipart or marshalled http body?
type User struct {
	UserName string
	Password string
	UserData
}

// UserData is a representation of the base data that is common for all users.
type UserData struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Country     string
	CompanyName string
}

// AdminUpdateableUser is a representation of data updateable by an admin.
type AdminUpdateableUser struct {
	UserData

	//Optional
	PostalCode string

	//Optional
	City string

	//Optional
	Street string
}

// BeneficiaryInput is a representation of data accompanying a request to
// create a beneficiary.
type BeneficiaryInput struct {
	Type string

	// Optional for types that are not Personal.
	FirstNameOnAccount string

	// Optional for types that are not Personal.
	LastNameOnAccount string

	// Optional for types that are not business.
	CompanyName string

	BankCountry       string
	BankName          string
	AccountNumber     string
	Currency          string
	Address1          string
	Country           string
	City              string
	PostalCode        string
	RoutingNumber     string
	SwiftBic          string
	BsbNumber         string
	Cpfcnpj           string
	StateProvince     string
	PhoneNumber       string
	BranchName        string
	BankCity          string
	BankStateProvince string
	Clabe             string
	BankCode          string
	TaxNumber         string
	BranchCode        string
}

// UpdateBeneficiaryInput represents a set of alterable fields for a benificiary.
type UpdateBeneficiaryInput struct {
	BeneficiaryInput
	Email          string
	Address2       string
	AccountType    string
	BankCity       string
	BankAddress1   string
	BankAddress2   string
	BankCountry    string
	BankPostalCode string
}

// QuoteInput denotes the input structure to create a quote.
type QuoteInput struct {
	SourceAmount        int64
	SourceCurrency      string
	DestinationCurrency string
	// Format "YYYY/MM/DD or MM/DD/YYYY"
	PaymentDate string
}

// TransferInput
type TransferInput struct {
	BeneficiaryID     int
	SourceAmount      int64
	DestinationAmount int64
	Reference         string
	QuoteUUID         string
	AutoComplete      bool
}

// TransferState represents the current state and date of any transaction.
type TransferState struct {
	State     string
	CreatedAt time.Time
}
