package routefusion

// UserDetails is a representation of the base struct that user based route
// options respond with.
type UserDetails struct {
	UUID                  string `json:"uuid"`
	FirstName             string `json:"first_name"`
	LastName              string `json:"last_name"`
	Email                 string `json:"email"`
	PhoneNumber           string `json:"phone_number"`
	Country               string `json:"country"`
	Verified              bool   `json:"verified"`
	Type                  string `json:"type"`
	VerificationSubmitted bool   `json:"verification_submitted"`
	CompanyName           string `json:"company_name"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
	MasterUser            bool   `json:"master_user"`
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
