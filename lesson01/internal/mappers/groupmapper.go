package mappers

import (
	"context"

	"github.com/ptsypyshev/gb-golang-backend-level02/lesson01/internal/models"
)

type GroupStore interface {
	Create(ctx context.Context, g models.Group) (*models.Group, error)
	Read(ctx context.Context, id int) (*models.Group, error)
	Update(ctx context.Context, id int, m map[string]any) error
	Delete(ctx context.Context, id int) error
	AddMember(ctx context.Context, user *models.User, group *models.Group) error
	RemoveMember(ctx context.Context, user *models.User, group *models.Group) error
	SearchGroupByName(ctx context.Context, groupName string) (chan models.Group, error)
	SearchGroupByMemberName(ctx context.Context, userName string) (chan models.Group, error)
}

type GroupMapper struct {
	gstore GroupStore
}

func NewGroupMapper(gstore GroupStore) *GroupMapper {
	return &GroupMapper{
		gstore: gstore,
	}
}

func (gm *GroupMapper) Create(ctx context.Context, g models.Group) (*models.Group, error) {
	return &g, nil
}

func (gm *GroupMapper) Read(ctx context.Context, id int) (*models.Group, error) {
	return &models.Group{}, nil
}

func (gm *GroupMapper) Update(ctx context.Context, id int, m map[string]any) error {
	return nil
}

func (gm *GroupMapper) Delete(ctx context.Context, id int) error {
	return nil
}

func (gm *GroupMapper) AddMember(ctx context.Context, user *models.User, group *models.Group) error {
	return nil
}

func (gm *GroupMapper) RemoveMember(ctx context.Context, user *models.User, group *models.Group) error {
	return nil
}

func (gm *GroupMapper) SearchGroupByName(ctx context.Context, groupName string) (chan models.Group, error) {
	return nil, nil
}

func (gm *GroupMapper) SearchGroupByMemberName(ctx context.Context, userName string) (chan models.Group, error) {
	return nil, nil
}
