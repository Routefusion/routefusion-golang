package routefusion

// Client specifies the abstraction for Routefusion APIs.
type Client interface {
	GetUser() (*UserDetails, error)
	UpdateUser(*User) (*UpdatedUserDetails, error)
	GetUserMaster(subUserUUID string) (*AllUserDetails, error)
	// TODO: Check pagination
	ListUsersMaster() ([]AllUserDetails, error)
	// TODO: Check pagination
	ListBeneficiaries() ([]Beneficiary, error)
	GetBeneficiary(id string) (*BeneficiaryBase, error)
	CreateBeneficiary(*BeneficiaryInput) (*BeneficiaryBase, error)
	UpdateBeneficiary(id string, body *UpdateBeneficiaryInput) (*BeneficiaryBase, error)
}
