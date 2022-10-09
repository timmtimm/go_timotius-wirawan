package usecase_test

import (
	"belajar-go-echo/model"
	_userMock "belajar-go-echo/repository/mocks"
	usecase "belajar-go-echo/usecase"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	userRepository _userMock.UserRepository
	userUseCase    usecase.UserUsecase
)

func TestMain(m *testing.M) {
	userUseCase = usecase.NewUserUsecase(&userRepository)

	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("Register | Valid", func(t *testing.T) {
		UserInput := model.UserInput{
			Email:    "budi@gmail.com",
			Password: "asd123",
		}
		User := model.User{}
		var err error

		userRepository.On("Register", UserInput).Return(User, err).Once()

		result, _ := userUseCase.Register(UserInput)

		assert.NotNil(t, result)
	})

	t.Run("Register | Invalid password", func(t *testing.T) {
		UserInput := model.UserInput{
			Email:    "budi@gmail.com",
			Password: "",
		}
		User := model.User{}
		var err error

		userRepository.On("Register", UserInput).Return(User, err).Once()

		_, err = userUseCase.Register(UserInput)

		assert.NotNil(t, err)
	})

	t.Run("Register | Invalid email", func(t *testing.T) {
		UserInput := model.UserInput{
			Email:    "",
			Password: "asd123",
		}
		User := model.User{}
		var err error

		userRepository.On("Register", UserInput).Return(User, err).Once()

		_, err = userUseCase.Register(UserInput)

		assert.NotNil(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		UserInput := model.UserInput{
			Email:    "budi@gmail.com",
			Password: "asd123",
		}
		User := model.User{}
		var err error

		userRepository.On("Login", UserInput).Return(User, err).Once()
		result2 := userUseCase.Login(UserInput)

		assert.NotNil(t, result2)
	})

	t.Run("Login | Invalid password", func(t *testing.T) {
		UserInput := model.UserInput{
			Email:    "budi@gmail.com",
			Password: "",
		}
		User := model.User{}
		var err error

		userRepository.On("Login", UserInput).Return(User, err).Once()

		result := userUseCase.Login(UserInput)

		assert.Empty(t, result)
	})

	t.Run("Login | Invalid email", func(t *testing.T) {
		UserInput := model.UserInput{
			Email:    "",
			Password: "asd123",
		}
		User := model.User{}
		var err error

		userRepository.On("Login", UserInput).Return(User, err).Once()

		result := userUseCase.Login(UserInput)

		assert.Empty(t, result)
	})
}

func TestGetAllUsers(t *testing.T) {
	t.Run("Get All Users | Valid", func(t *testing.T) {
		User := []model.User{}
		var err error

		userRepository.On("GetAllUsers").Return(User, err).Once()
		result, _ := userUseCase.GetAllUsers()

		assert.NotNil(t, result)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("Get All Users | Valid", func(t *testing.T) {
		UserInput := model.User{
			Email:    "budi@gmail.com",
			Password: "asd123",
		}

		var err error

		userRepository.On("CreateUser", UserInput).Return(err).Once()
		result := userUseCase.CreateUser(UserInput)

		assert.Nil(t, result)
	})
}
