package routefusion

import "io"

// Client specifies the abstraction for Routefusion APIs.
type Client interface {
	Users
	Beneficiaries
	Quotes
	Transfers
	BatchTransfers
	Transactions
	Account
	Webhooks
	KYC
	CurrencyCoverage
	WireInstructions
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

// Transfers specifies the operations that can be performed around transfers.
type Transfers interface {
	CreateTransfer(*TransferInput) (*TransferResponse, error)
	GetTransfer(id string) (*TransferResponse, error)
	CancelTransfer(uuid string) (cancelledID string, err error)
	CreateTransferMaster(uuid string) (*TransferState, error)
	GetTransferMaster(subUserID, transferID string) (*TransferResponse, error)
	GetTransferStatusMaster(subUserID, transferID string) (*TransferState, error)
	CancelTransferMaster(subUserID, transferID string) (cancelledID string, err error)
}

// BatchTransfers specifies the operations that can be performed on batch
// transfers.
type BatchTransfers interface {
	CreateBatchPayment(payload io.ReadSeeker) (*BatchTransferStatus, error)
	GetBatchPayment(batchID string) (*BatchTransferStatus, error)
}

// Transactions is an interface that specifies operations concerning reading
// of transactions.
type Transactions interface {
	GetTransactions() ([]TransactionResponse, error)
}

// Account dictates an interface for retrieving account reports.
type Account interface {
	GetBalance() (*BalanceResponse, error)
}

// Webhooks is an interface for webhook based operation.
type Webhooks interface {
	GetWebhook(id string) (*WebhookResponse, error)
	UpdateWebhook(id string, updateInput WebhookUpdateInput) (*WebhookResponse, error)
	IndexWebhooks()
	CreateWebhook(createInput WebhookUpdateInput) (*WebhookResponse, error)
	DeleteWebhook(id string) error
}

// KYC is an interface for KYC based CRUD operations.
type KYC interface {
	// TODO: Solve this problem, struct vs io.ReadSeeker
	CreateKYC(subUserID string, kycBody KYCBody) error
	ShowKYC(subUserID string) (*KYCDetails, error)
	UpdateUserKYC(subUserID string, kycBody KYCBody) error
	DeleteKYC(subUserID string) error
}

// CurrencyCoverage is an interface for currency based transactions.
type CurrencyCoverage interface {
	// TODO: Interface{} is because I cant make sense of the response to
	// this endpoint. Need to fix.
	GetCurrencies() (interface{}, error)
}

// WireInstructions is an interface for wireinstruction based operations.
type WireInstructions interface {
	GetWireInstructions(currencyCode string) ([]PaymentInstructions, error)
}
