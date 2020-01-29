package routefusion

// Client specifies the abstraction for Routefusion APIs.
// TODO: Question - Should the client be a unified interface or can I split
// it up into sub interfaces like Users, Beneficiaries and embded them in
// the unified interface? This adds a bit of overhead to the client while
// calling which is why I only want to do it if absolutely needed.
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
	GetSubUserBeneficiariesMaster(subuserID string) ([]Beneficiary, error)
	GetSubUserBeneficiaryMaster(subuserID string, beneficiaryID string) (*BeneficiaryBase, error)
	CreateSubUserBeneficiaryMaster(subUserID string) (*BeneficiaryBase, error)
	UpdateSubUserBeneficiaryMaster(subUserID string, beneficiaryID string) (*BeneficiaryBase, error)
}
