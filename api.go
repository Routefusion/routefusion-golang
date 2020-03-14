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

	// GET v1/users/me
	GetUser() (*UserDetails, error)

	// PUT v1/users/me
	UpdateUser(*User) (*UpdatedUserDetails, error)

	// GET v1/users/me/users/{subUserID}
	GetUserMaster(subUserID string) (*AllUserDetails, error)

	// Get v1/users/me/users
	ListUsersMaster() ([]AllUserDetails, error)
}

// Beneficiaries specifies the operations that can be performed around benficiaries.
type Beneficiaries interface {

	// Get v1/beneficiaries
	ListBeneficiaries() ([]Beneficiary, error)

	// GET v1/beneficiaries/{id}
	GetBeneficiary(id string) (*BeneficiaryBase, error)

	// POST v1/beneficiaries
	CreateBeneficiary(*BeneficiaryInput) (*BeneficiaryBase, error)

	// PUT v1/beneficiaries/{id}
	UpdateBeneficiary(id string, body *UpdateBeneficiaryInput) (*BeneficiaryBase, error)

	// GET v1/users/{subUserID}
	GetSubUserBeneficiariesMaster(subUserID string) ([]Beneficiary, error)

	// GET v1/users/{subUserID}/beneficiaries/{beneficiary}
	GetSubUserBeneficiaryMaster(subUserID string, beneficiaryID string) (*BeneficiaryBase, error)

	// POST v1/users/{subuserID}/beneficiaries
	CreateSubUserBeneficiaryMaster(subUserID string) (*BeneficiaryBase, error)

	// PUT v1/users/{subUserID}/beneficiaries/{beneficiaryID}
	UpdateSubUserBeneficiaryMaster(subUserID string, beneficiaryID string) (*BeneficiaryBase, error)
}

// Quotes specifies the operations that can be performed around quotes.
type Quotes interface {
	// POST v1/quotes
	CreateQuote(*QuoteInput) (*QuoteResponse, error)
}

// Transfers specifies the operations that can be performed around transfers.
type Transfers interface {

	// POST v1/transfers
	CreateTransfer(*TransferInput) (*TransferResponse, error)

	// GET v1/transfers/{id}
	GetTransfer(id string) (*TransferResponse, error)

	// DELETE v1/transfers/{uuid}/cancel
	CancelTransfer(uuid string) (cancelledID string, err error)

	// GET v1/transfers/{uuid}/status
	GetTransferStatus(uuid string) (*TransferStatus, error)

	// POST v1/users/{uuid}/transfers
	CreateTransferMaster(uuid string) (*TransferState, error)

	// GET v1/users/subUserID/transfers/{transferID}
	GetTransferMaster(subUserID, transferID string) (*TransferResponse, error)

	// GET v1/users/{subUserID}/transfers/{transferID}/status
	GetTransferStatusMaster(subUserID, transferID string) (*TransferState, error)

	// DELETE v1/users/{subUserID}/transfers/{transferID}
	CancelTransferMaster(subUserID, transferID string) (cancelledID string, err error)
}

// BatchTransfers specifies the operations that can be performed on batch
// transfers.
type BatchTransfers interface {
	// POST v1/batches
	CreateBatchPayment(payload io.ReadSeeker) (*BatchTransferStatus, error)

	// GET v1/batches/{batchID}
	GetBatchPayment(batchID string) (*BatchTransferStatus, error)
}

// Transactions is an interface that specifies operations concerning reading
// of transactions.
type Transactions interface {
	// GET v1/transfers
	GetTransactions() ([]TransactionResponse, error)
}

// Account dictates an interface for retrieving account reports.
type Account interface {
	// GET v1/balance
	GetBalance() (*BalanceResponse, error)
}

// Webhooks is an interface for webhook based operation.
type Webhooks interface {
	// GET /v1/webhooks/{id}
	GetWebhook(id string) (*WebhookResponse, error)

	// PUT v1/webhooks/{id}
	UpdateWebhook(id string, updateInput WebhookUpdateInput) (*WebhookResponse, error)

	// GET v1/webhooks
	IndexWebhooks() ([]WebhookResponse, error)

	// POST v1/webhooks
	CreateWebhook(createInput WebhookUpdateInput) (*WebhookResponse, error)

	// DELETE v1/webhooks/{id}
	DeleteWebhook(id string) error
}

// KYC is an interface for KYC based CRUD operations.
type KYC interface {
	// POST v1/users/{subUserID}/verify
	CreateKYC(subUserID string, kycBody KYCBody) error

	// GET v1/users/{subUserID}/verify
	ShowKYC(subUserID string) (*KYCDetails, error)

	// PUT v1/users/{subUserID}/verify
	UpdateUserKYC(subUserID string, kycBody KYCBody) error

	// DELETE v1/users/{subUserID}/verify
	DeleteKYC(subUserID string) error
}

// CurrencyCoverage is an interface for currency based transactions.
type CurrencyCoverage interface {
	// Ignore
	// TODO: Interface{} is because I cant make sense of the response to
	// this endpoint. Need to fix.
	GetCurrencies() (interface{}, error)
}

// WireInstructions is an interface for wireinstruction based operations.
type WireInstructions interface {
	// GET v1/instructions?currency_code={currencyCode}
	GetWireInstructions(currencyCode string) ([]PaymentInstructions, error)
}
