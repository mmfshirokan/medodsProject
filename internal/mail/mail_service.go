package mail

type MailService interface {
	SendMail(mail string) error
}

// Returns Mock for MailService
func New() MailService {
	return &ms{}
}

type ms struct {
}

func (m *ms) SendMail(mail string) error {
	return nil
}
