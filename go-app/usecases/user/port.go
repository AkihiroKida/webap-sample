package user

import (
	"context"

	"go-app/entities"
)

type UserInputPort interface {
	GetUserByID(ctx context.Context, userID string)
}

type UserOutputPort interface {
	Render(*entities.User)
	RenderError(error)
}

type UserRepository interface {
	GetUserByID(ctx context.Context, UserID string) (*entities.User, error)
}
