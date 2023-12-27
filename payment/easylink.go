package payment

type EasyLinkPaymentGateway struct {
	Pin         string
	SecPin      string
	ScretKey    string
	ChannelID   string
	ChannelURL  string
	ChannelType string
	CallbackUrl string
}

func (p *EasyLinkPaymentGateway) Create(payment Payment) (Payment, error) {
	return payment, nil
}

func (p *EasyLinkPaymentGateway) Check(payment Payment) (Payment, error) {
	return payment, nil
}

func (p *EasyLinkPaymentGateway) Cancel(payment Payment) (Payment, error) {
	return payment, nil
}

func (p *EasyLinkPaymentGateway) Refund(payment Payment, amount int64) (Payment, error) {
	return payment, nil
}

func (p *EasyLinkPaymentGateway) Webhook(data any) {
}

func (p *EasyLinkPaymentGateway) Name() string {
	return "easylink"
}
