package email

type Email struct {
	To   string
	From string
	Body string
}

func NewEmail() *Email {
	return &Email{}
}

func SendMessage() {

}
