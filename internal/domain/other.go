package domain

type Repositories struct {
	UserRepo    UserRepository
	JokeRepo    JokeRepository
	LikeRepo    LikeRepository
	CommentRepo CommentRepository
}

type Packages struct {
	PasswordPkg PasswordPackage
	TokenPkg    TokenPackage
	EmailPkg    EmailPackage
}
