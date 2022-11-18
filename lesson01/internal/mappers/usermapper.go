package mappers

import (
	"context"

	"github.com/ptsypyshev/gb-golang-backend-level02/lesson01/internal/models"
)

type UserStore interface {
	Create(ctx context.Context, u models.User) (int, error)
	Read(ctx context.Context, id int) (*models.User, error)
	Update(ctx context.Context, id int, m map[string]any) error
	Delete(ctx context.Context, id int) error
	JoinGroup(ctx context.Context, groupId int) error
	LeaveGroup(ctx context.Context, groupId int) error
	SearchUsersByName(ctx context.Context, userName string) (chan models.User, error)
	SearchUsersByGroupName(ctx context.Context, groupName string) (chan models.User, error)
}

type UserMapper struct {
	ustore UserStore
}

func NewUserMapper(ustore UserStore) *UserMapper {
	return &UserMapper{
		ustore: ustore,
	}
}

func (um *UserMapper) Create(ctx context.Context, u models.User) (*models.User, error) {
	return &u, nil
}

func (um *UserMapper) Read(ctx context.Context, id int) (*models.User, error) {
	return &models.User{}, nil
}

func (um *UserMapper) Update(ctx context.Context, id int, m map[string]any) error {
	return nil
}

func (um *UserMapper) Delete(ctx context.Context, id int) error {
	return nil
}

func (um *UserMapper) JoinGroup(ctx context.Context, groupId int) error {
	return nil
}

func (um *UserMapper) LeaveGroup(ctx context.Context, groupId int) error {
	return nil
}

func (um *UserMapper) SearchUsersByName(ctx context.Context, userName string) (chan models.User, error) {
	return nil, nil
}

func (um *UserMapper) SearchUsersByGroupName(ctx context.Context, groupName string) (chan models.User, error) {
	return nil, nil
}