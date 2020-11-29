package interfaces

type ICaptchaRepo interface {
	SetCaptcha(phone, code string) (err error)
	GetCaptcha(phone string) (code string, err error)
	DelCaptcha(phone string) (code int64, err error)
}
