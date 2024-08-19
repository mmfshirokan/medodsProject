package mail

// Returns Mock for MailService
func New() *Ms {
	return &Ms{}
}

type Ms struct {
}

func (m *Ms) SendMail(mail string) error {
	return nil
}
