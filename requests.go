package routefusion

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

	BankCountry   string
	BankName      string
	AccountNumber string
	Currency      string
	Address1      string
	Country       string
	City          string
	PostalCode    string
}
