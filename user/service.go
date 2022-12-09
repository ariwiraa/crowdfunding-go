package user

import "golang.org/x/crypto/bcrypt"

//Kontrak Service
type Service interface {
	Register(input RegisterUserInput) (User, error)
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
