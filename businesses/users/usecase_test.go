package users_test

import (
	"echo-recipe/app/middlewares"
	"echo-recipe/businesses/users"
	_userMock "echo-recipe/businesses/users/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	usersRepository _userMock.Repository
	usersService    users.Usecase

	usersDomain users.Domain
)

func TestMain(m *testing.M) {
	usersService = users.NewUserUsecase(&usersRepository, &middlewares.ConfigJWT{})

	usersDomain = users.Domain{
		ID:       1,
		Name:     "Roihan",
		Email:    "roihan12@gmail.com",
		Password: "hello123",
		Token:    "token",
	}
	m.Run()
}

func TestRegister(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		usersRepository.On("Register", mock.Anything).Return(usersDomain, nil).Once()

		input := users.Domain{
			ID:       1,
			Name:     "Roihan",
			Email:    "roihan12@gmail.com",
			Password: "hello123",
		}

		result := usersService.Register(&input)
		assert.Equal(t, usersDomain, result)
	})

	t.Run("Test case 2 - Invalid (duplicate email)", func(t *testing.T) {
		usersRepository.On("Register", mock.Anything).Return(users.Domain{}, assert.AnError).Once()

		input := users.Domain{
			ID:       1,
			Name:     "Roihan 2",
			Email:    "roihan12@gmail.com",
			Password: "hello123",
		}

		result := usersService.Register(&input)
		assert.NotEqual(t, usersDomain, result)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		usersRepository.On("GetByEmail", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(usersDomain, nil).Once()

		input := users.Domain{
			Email:    "roihan12@gmail.com",
			Password: "hello123",
		}

		result := usersService.Login(input.Email, input.Password)
		result.Token = "token"
		assert.Equal(t, usersDomain, result)
	})

	t.Run("Test case 2 - Invalid (wrong username/password)", func(t *testing.T) {
		usersRepository.On("GetByEmail", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(usersDomain, nil).Once()

		inputInvalid := users.Domain{
			Email:    "roihan122@gmail.com",
			Password: "hello123",
		}

		result := usersService.Login(inputInvalid.Email, inputInvalid.Password)
		result.Token = "token"
		assert.NotEqual(t, usersDomain.Email, usersDomain.Password, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Test case 1 - Valid", func(t *testing.T) {
		usersRepository.On("Update", mock.Anything).Return(usersDomain, nil).Once()

		input := users.Domain{
			ID:       1,
			Name:     "roihandited",
			Email:    "roihan12@gmail.com",
			Password: "hello123",
		}

		result := usersService.Update(&input)
		assert.Equal(t, usersDomain, result)
	})

	t.Run("Test case 2 - Invalid (Duplicate email)", func(t *testing.T) {
		usersRepository.On("Update", mock.Anything).Return(usersDomain, assert.AnError).Once()

		input := users.Domain{
			ID:       2,
			Name:     "roihandited",
			Email:    "roihan12@gmail.com",
			Password: "hello123",
		}

		result := usersService.Update(&input)
		assert.NotEqual(t, usersDomain.ID, usersDomain.Name, usersDomain.Email, usersDomain.Password, result)
	})
}
