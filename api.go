package routefusion

// Client specifies the abstraction for Routefusion APIs.
type Client interface {
	Users
	Beneficiaries
	Quotes
}

// Users specifies the user related tasks that can be performed using the sdk.
type Users interface {
	GetUser() (*UserDetails, error)
	UpdateUser(*User) (*UpdatedUserDetails, error)
	GetUserMaster(subUserUUID string) (*AllUserDetails, error)
	// TODO: Check pagination
	ListUsersMaster() ([]AllUserDetails, error)
}

// Beneficiaries specifies the operations that can be performed around benficiaries.
type Beneficiaries interface {
	// TODO: Check pagination
	ListBeneficiaries() ([]Beneficiary, error)
	GetBeneficiary(id string) (*BeneficiaryBase, error)
	CreateBeneficiary(*BeneficiaryInput) (*BeneficiaryBase, error)
	UpdateBeneficiary(id string, body *UpdateBeneficiaryInput) (*BeneficiaryBase, error)
	GetSubUserBeneficiariesMaster(subuserID string) ([]Beneficiary, error)
	GetSubUserBeneficiaryMaster(subuserID string, beneficiaryID string) (*BeneficiaryBase, error)
	CreateSubUserBeneficiaryMaster(subUserID string) (*BeneficiaryBase, error)
	UpdateSubUserBeneficiaryMaster(subUserID string, beneficiaryID string) (*BeneficiaryBase, error)
}

// Quotes specifies the operations that can be performed around quotes.
type Quotes interface {
	CreateQuote(*QuoteInput) (*QuoteResponse, error)
}
