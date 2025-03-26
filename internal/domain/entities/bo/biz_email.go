package bo

type Email struct {
	To       []string
	From     string
	Subject  string
	Text     []byte
	HTML     []byte
	WithLogo bool
}
