package service

type PasswordService interface {
	// HashPassword returns hash of the provided password.
	//
	// Returned codes:
	// 	- Internal
	HashPassword(password string) (string, error)

	// CheckPassword checks if provided password matches provided hash.
	// If matches returns nil.
	//
	// Returned codes:
	//  - InvalidArgument
	// 	- Internal
	CheckPassword(password, hashedPassword string) error
}
