package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

//Kontrak Service
type Service interface {
	Register(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmail) (bool, error)
	SaveAvatar(id int, fileLocation string) (User, error)
	GetUserById(id int) (User, error)
}

type service struct {
	repository Repository
}

//Instance Service
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(input RegisterUserInput) (User, error) {
	user := User{}

	//Binding ke Struct User
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	//Hashing Password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.Password = string(passwordHash)
	user.Role = "User"	
	
	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
	
}	

func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	//Passing data ke repository
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	//Jika tidak ada ID
	if user.ID == 0 {
		return user, errors.New("User Not Found")
	}

	//cocokin password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmail) (bool, error) {
	email := input.Email

	//Cek apakah sudah ada yang memakai email
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	//cek apakah ada user yang menggunakan email tersebut
	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func(s *service)  SaveAvatar(id int, fileLocation string) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.Avatar = fileLocation

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
		return user, errors.New("User Not Found with that id")
	}

	return user, nil
}