package payment

type PaymentGateway interface {
	Create(payment Payment) (Payment, error)
	Check(payment Payment) (Payment, error)
	Cancel(payment Payment) (Payment, error)
	Refund(payment Payment, amount int64) (Payment, error)
	Webhook(data any)
	Name() string
}

type PaymentStatus string

const (
	PaymentStatusCreated  PaymentStatus = "created"
	PaymentStatusPaied    PaymentStatus = "paied"
	PaymentStatusCanceled PaymentStatus = "canceled"
	PaymentStatusFailed   PaymentStatus = "failed"
	PaymentStatusExpired  PaymentStatus = "expired"
	PaymentStatusRefunded PaymentStatus = "refunded"
)

type Payment struct {
	ID             string
	Amount         int64
	Refunded       int64
	Status         PaymentStatus
	Url            string
	Currency       string
	PaymentGateway string
}
