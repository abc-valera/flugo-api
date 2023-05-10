package mock

import "github.com/abc-valera/flugo-api/internal/infrastructure/repository"

// Repositories contains mock repository structs
type Repositories struct {
	UserRepository    *UserRepository
	JokeRepository    *JokeRepository
	LikeRepository    *LikeRepository
	CommentRepository *CommentRepository
}

// Returns two structs:
//   - repos with interfaces which is used to init services
//   - repos with mock structs which is used to references mocks
func NewRepositories() (*repository.Repositories, *Repositories) {
	ur := new(UserRepository)
	jr := new(JokeRepository)
	lr := new(LikeRepository)
	cr := new(CommentRepository)
	return &repository.Repositories{
			UserRepository:    ur,
			JokeRepository:    jr,
			LikeRepository:    lr,
			CommentRepository: cr,
		}, &Repositories{
			UserRepository:    ur,
			JokeRepository:    jr,
			LikeRepository:    lr,
			CommentRepository: cr,
		}
}
