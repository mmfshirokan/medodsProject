package mail

import log "github.com/sirupsen/logrus"

// Returns Mock for MailService
func New() *Ms {
	return &Ms{}
}

type Ms struct {
}

func (m *Ms) SendMail(mail string) error {
	log.Info("Recived mail: ", mail)

	// Some logic for sending mail here

	return nil
}
