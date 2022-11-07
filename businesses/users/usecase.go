package users

import "echo-recipe/app/middlewares"

type UserUsecase struct {
	userRepository Repository
	jwtAuth        *middlewares.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtAuth *middlewares.ConfigJWT) Usecase {
	return &UserUsecase{
		userRepository: ur,
		jwtAuth:        jwtAuth,
	}
}

func (uu *UserUsecase) Register(userDomain *Domain) Domain {
	return uu.userRepository.Register(userDomain)
}

func (uu *UserUsecase) Login(email string, password string) Domain {
	user := uu.userRepository.GetByEmail(email, password)

	if user.ID == 0 {
		return Domain{}
	}

	user.Token = uu.jwtAuth.GenerateToken(user.ID)

	return user
}

func (uu *UserUsecase) Update(userDomain *Domain) Domain {
	return uu.userRepository.Update(userDomain)
}
