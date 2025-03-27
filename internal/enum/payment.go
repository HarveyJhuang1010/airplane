package enum

//go:generate enumer -type=PaymentStatus -trimprefix=PaymentStatus -yaml -json -text -sql -transform=snake --output=zzz_enumer_PaymentStatus.go
type PaymentStatus int32

const (
	PaymentStatusPending PaymentStatus = iota
	PaymentStatusSuccess
	PaymentStatusFailed
	PaymentStatusCancelled
	PaymentStatusRefunding
	PaymentStatusRefunded
)

//go:generate enumer -type=PaymentProvider -trimprefix=PaymentProvider -yaml -json -text -sql -transform=snake --output=zzz_enumer_PaymentProvider.go
type PaymentProvider int32

const (
	PaymentProviderStripe PaymentProvider = iota
	PaymentProviderPaypal
	PaymentProviderLinePay
	PaymentProviderApplePay
	PaymentProviderGooglePay
)

//go:generate enumer -type=PaymentMethod -trimprefix=PaymentMethod -yaml -json -text -sql -transform=snake --output=zzz_enumer_PaymentMethod.go
type PaymentMethod int32

const (
	PaymentMethodCreditCard PaymentMethod = iota
	PaymentMethodDebitCard
	PaymentMethodBankTransfer
)
