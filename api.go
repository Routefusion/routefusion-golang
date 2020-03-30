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

	// Get v1/users/me
	GetUser() (*UserDetails, error)

	// Put v1/users/me
	UpdateUser(user *User) (*UpdatedUserDetails, error)

	// Ignore
	// Get v1/users/me/users/{subUserID}
	GetUserMaster(subUserID string) (*AllUserDetails, error)

	// Get v1/users/me/users
	ListUsersMaster() ([]AllUserDetails, error)
}

// Beneficiaries specifies the operations that can be performed around benficiaries.
type Beneficiaries interface {

	// Get v1/beneficiaries
	ListBeneficiaries() ([]Beneficiary, error)

	// Get v1/beneficiaries/{id}
	GetBeneficiary(id string) (*BeneficiaryBase, error)

	// Post v1/beneficiaries
	CreateBeneficiary(beneficiary *BeneficiaryInput) (*BeneficiaryBase, error)

	// Ignore
	// Put v1/beneficiaries/{id}
	UpdateBeneficiary(id string, body *UpdateBeneficiaryInput) (*BeneficiaryBase, error)

	// Ignore
	// Get v1/users/{subUserID}
	GetSubUserBeneficiariesMaster(subUserID string) ([]Beneficiary, error)

	// Ignore
	// Get v1/users/{subUserID}/beneficiaries/{beneficiary}
	GetSubUserBeneficiaryMaster(subUserID string, beneficiaryID string) (*BeneficiaryBase, error)

	// Ignore
	// Post v1/users/{subuserID}/beneficiaries
	CreateSubUserBeneficiaryMaster(subUserID string) (*BeneficiaryBase, error)

	// Ignore
	// Put v1/users/{subUserID}/beneficiaries/{beneficiaryID}
	UpdateSubUserBeneficiaryMaster(subUserID string, beneficiaryID string) (*BeneficiaryBase, error)
}

// Quotes specifies the operations that can be performed around quotes.
type Quotes interface {

	// Post v1/quotes
	CreateQuote(*QuoteInput) (*QuoteResponse, error)
}

// Transfers specifies the operations that can be performed around transfers.
type Transfers interface {

	// Post v1/transfers
	CreateTransfer(input *TransferInput) (*TransferResponse, error)

	// Ignore
	// Get v1/transfers/{id}
	GetTransfer(id string) (*TransferResponse, error)

	// Ignore
	// Delete v1/transfers/{uuid}/cancel
	CancelTransfer(uuid string) (cancelledID string, err error)

	// Ignore
	// Get v1/transfers/{uuid}/status
	GetTransferStatus(uuid string) (*TransferStatus, error)

	// Ignore
	// Post v1/users/{uuid}/transfers
	CreateTransferMaster(uuid string) (*TransferState, error)

	// Ignore
	// Get v1/users/subUserID/transfers/{transferID}
	GetTransferMaster(subUserID, transferID string) (*TransferResponse, error)

	// Ignore
	// Get v1/users/{subUserID}/transfers/{transferID}/status
	GetTransferStatusMaster(subUserID, transferID string) (*TransferState, error)

	// Ignore
	// Delete v1/users/{subUserID}/transfers/{transferID}
	CancelTransferMaster(subUserID, transferID string) (cancelledID string, err error)
}

// BatchTransfers specifies the operations that can be performed on batch
// transfers.
type BatchTransfers interface {

	// Post v1/batches
	CreateBatchPayment(payload io.ReadSeeker) (*BatchTransferStatus, error)

	// Ignore
	// Get v1/batches/{batchID}
	GetBatchPayment(batchID string) (*BatchTransferStatus, error)
}

// Transactions is an interface that specifies operations concerning reading
// of transactions.
type Transactions interface {

	// Get v1/transfers
	GetTransactions() ([]TransactionResponse, error)
}

// Account dictates an interface for retrieving account reports.
type Account interface {

	// Get v1/balance
	GetBalance() (*BalanceResponse, error)
}

// Webhooks is an interface for webhook based operation.
type Webhooks interface {

	// Ignore
	// Get /v1/webhooks/{id}
	GetWebhook(id string) (*WebhookResponse, error)

	// Ignore
	// Put v1/webhooks/{id}
	UpdateWebhook(id string, updateInput WebhookUpdateInput) (*WebhookResponse, error)

	// Get v1/webhooks
	IndexWebhooks() ([]WebhookResponse, error)

	// Post v1/webhooks
	CreateWebhook(createInput WebhookUpdateInput) (*WebhookResponse, error)

	// Ignore
	// Delete v1/webhooks/{id}
	DeleteWebhook(id string) error
}

// KYC is an interface for KYC based CRUD operations.
type KYC interface {

	// Ignore
	// Post v1/users/{subUserID}/verify
	CreateKYC(subUserID string, kycBody KYCBody) error

	// Ignore
	// Get v1/users/{subUserID}/verify
	ShowKYC(subUserID string) (*KYCDetails, error)

	// Ignore
	// Put v1/users/{subUserID}/verify
	UpdateUserKYC(subUserID string, kycBody KYCBody) error

	// Ignore
	// Delete v1/users/{subUserID}/verify
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

	// Ignore
	// Get v1/instructions?currency_code={currencyCode}
	GetWireInstructions(currencyCode string) ([]PaymentInstructions, error)
}
