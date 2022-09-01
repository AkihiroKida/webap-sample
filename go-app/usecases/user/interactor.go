package user

import "context"

type User struct {
	OutputPort UserOutputPort
	UserRepo   UserRepository
}

func NewUserInputPort(outputPort UserOutputPort, userRepository UserRepository) UserInputPort {
	return &User{
		OutputPort: outputPort,
		UserRepo:   userRepository,
	}
}

func (u *User) GetUserByID(ctx context.Context, userID string) {
	user, err := u.UserRepo.GetUserByID(ctx, userID)
	if err != nil {
		u.OutputPort.RenderError(err)
		return
	}
	u.OutputPort.Render(user)
}
