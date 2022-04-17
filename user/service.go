package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisteruserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInnput) (bool, error)
	SaveAvatar(id int, filelocation string) (User, error)
	GetUserById(id int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisteruserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email

	userTest, err := s.repository.FindByEmail(input.Email)

	if err != nil {
		return user, err
	}

	if userTest.ID != 0 {
		return user, err
	}

	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInnput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) SaveAvatar(id int, filelocation string) (User, error) {
	// dapatkan user berdasarkan ID
	// update attribute avatar filename
	// simpan perubahan avatar filename

	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = filelocation
	updatedUser, err := s.repository.Update(user)

	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserById(id int) (User, error) {
	user, err := s.repository.FindById(id)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("No user with that ID")
	}

	return user, nil
}
