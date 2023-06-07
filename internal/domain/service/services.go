package service

type Services struct {
	Logger        Logger
	PasswordMaker PasswordMaker
	TokenMaker    TokenMaker
	EmailSender   EmailSender
}
