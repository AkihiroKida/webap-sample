package gateways

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"go-app/entities"
	"go-app/usecases/user"
)

type UserRepository struct {
	conn *sql.DB
}

func NewUserRepository(conn *sql.DB) user.UserRepository {
	return &UserRepository{
		conn: conn,
	}
}

func (u *UserRepository) GetUserByID(ctx context.Context, userID string) (*entities.User, error) {
	conn := u.GetDBConn()
	row := conn.QueryRowContext(ctx, "SELECT * FROM `user` WHERE id=?", userID)
	user := entities.User{}
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User Not Found. UserID = %s", userID)
		}
		log.Println(err)
		return nil, errors.New("Internal Server Error. adapter/gateway/GetUserByID")
	}
	return &user, nil
}

func (u *UserRepository) GetDBConn() *sql.DB {
	return u.conn
}
