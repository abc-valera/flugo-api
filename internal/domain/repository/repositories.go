package repository

type Repositories struct {
	UserRepo    UserRepository
	JokeRepo    JokeRepository
	LikeRepo    LikeRepository
	CommentRepo CommentRepository
}
