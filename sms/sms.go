package sms

// A struct represent phone number with Area code
type PhoneNumber struct {
	AreaCode string
	Number   string
}

type SmsSender interface {
	SendWithTemplate(to PhoneNumber, templateId string, vars map[string]string) error
}
