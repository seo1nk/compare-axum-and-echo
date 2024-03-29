package user

type SaveUserUseCase struct {
	userRepo userDomain.UserRepository
}

func NewSaveUserUseCase(
	userRepo userDomain.UserRepository,
) *SaveUserUseCase {
	return &SaveUserUseCase{
		userRepo: userRepo,
	}
}

type SaveUserDto struct {
	firstname string
	lastname  string
}
