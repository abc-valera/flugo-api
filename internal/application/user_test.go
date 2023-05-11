package application

// import (
// 	"context"
// 	"testing"

// 	"github.com/abc-valera/flugo-api/internal/domain"
// 	"github.com/abc-valera/flugo-api/internal/infrastructure/framework"
// 	FMock "github.com/abc-valera/flugo-api/internal/infrastructure/framework/mock"
// 	"github.com/abc-valera/flugo-api/internal/infrastructure/repository"
// 	RMock "github.com/abc-valera/flugo-api/internal/infrastructure/repository/mock"
// 	"github.com/stretchr/testify/suite"
// )

// type userServiceSuite struct {
// 	suite.Suite
// 	repos   *RMock.Repositories
// 	frs     *FMock.Frameworks
// 	service UserService
// }

// func (s *userServiceSuite) SetupSuite() {
// }

// func (s *userServiceSuite) SetupTest() {
// 	var r *repository.Repositories
// 	var f *framework.Frameworks
// 	r, s.repos = RMock.NewRepositories()
// 	f, s.frs = FMock.NewFrameworks()
// 	s.service = newUserService(r, f)
// }

// func (s *userServiceSuite) TearDownTest() {
// }

// func (s *userServiceSuite) TearDownSuite() {
// }

// func (s *userServiceSuite) TestSignUp() {
// 	password := "test"
// 	returnPassword := "1234"
// 	s.frs.PasswordFramework.EXPECT().HashPassword(password).Return(returnPassword, nil)

// 	createUser := &domain.User{
// 		Username:       "test",
// 		Email:          "test@gmail.com",
// 		HashedPassword: returnPassword,
// 	}
// 	s.repos.UserRepository.EXPECT().CreateUser(context.Background(), createUser).Return(nil)

// 	user := &domain.User{
// 		Username: "test",
// 		Email:    "test@gmail.com",
// 	}
// 	err := s.service.SignUp(context.Background(), user, password)
// 	s.NoError(err)
// }

// func (s *userServiceSuite) TestSignIn() {
// 	username := "test"
// 	email := "email@gmail.com"
// 	hashedPassword := "1234"
// 	password := "test"
// 	user := &domain.User{
// 		Username:       username,
// 		Email:          email,
// 		HashedPassword: hashedPassword,
// 	}

// 	s.repos.UserRepository.EXPECT().GetUserByEmail(context.Background(), email).Return(user, nil)
// 	s.frs.PasswordFramework.EXPECT().CheckPassword(password, hashedPassword).Return(nil)

// 	access, refresh := "access", "refresh"
// 	s.frs.TokenFramework.EXPECT().CreateAccessToken(username).Return(access, nil, nil)
// 	s.frs.TokenFramework.EXPECT().CreateRefreshToken(username).Return(refresh, nil, nil)

// 	rUser, rAccess, rRefresh, err := s.service.SignIn(context.Background(), email, password)
// 	s.NoError(err)
// 	s.Equal(user, rUser)
// 	s.Equal(access, rAccess)
// 	s.Equal(refresh, rRefresh)
// }

// func (s *userServiceSuite) TestUpdateUserPassword() {
// 	username := "test"
// 	email := "email@gmail.com"
// 	hashedPassword := "1234"
// 	oldPassword := "test"
// 	newPassword := "test_new"
// 	newHash := "5432"
// 	user := &domain.User{
// 		Username:       username,
// 		Email:          email,
// 		HashedPassword: hashedPassword,
// 	}

// 	s.repos.UserRepository.EXPECT().GetUserByUsername(context.Background(), username).Return(user, nil)
// 	s.frs.PasswordFramework.EXPECT().CheckPassword(oldPassword, hashedPassword).Return(nil)
// 	s.frs.PasswordFramework.EXPECT().HashPassword(newPassword).Return(newHash, nil)
// 	s.repos.UserRepository.EXPECT().UpdateUserHashedPassword(context.Background(), username, newHash).Return(nil)

// 	err := s.service.UpdateUserPassword(context.Background(), username, oldPassword, newPassword)
// 	s.NoError(err)
// }

// func (s *userServiceSuite) TestDeleteUser() {
// 	username := "test"
// 	email := "email@gmail.com"
// 	hashedPassword := "1234"
// 	password := "test"
// 	user := &domain.User{
// 		Username:       username,
// 		Email:          email,
// 		HashedPassword: hashedPassword,
// 	}

// 	s.repos.UserRepository.EXPECT().GetUserByUsername(context.Background(), username).Return(user, nil)
// 	s.frs.PasswordFramework.EXPECT().CheckPassword(password, hashedPassword).Return(nil)
// 	s.repos.UserRepository.EXPECT().DeleteUser(context.Background(), username).Return(nil)

// 	err := s.service.DeleteUser(context.Background(), username, password)
// 	s.NoError(err)
// }

// func TestUserServiceSuite(t *testing.T) {
// 	suite.Run(t, new(userServiceSuite))
// }
