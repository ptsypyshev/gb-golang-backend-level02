package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/ptsypyshev/gb-golang-backend-level02/lesson01/internal/mappers"
	"github.com/ptsypyshev/gb-golang-backend-level02/lesson01/internal/models"
)

const DefaultGroupMembersSize = 10

func main() {
	fmt.Println("It's ready! But storage and methods are not implemented now :-(")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	users := mappers.NewUserMapper(nil)
	groups := mappers.NewGroupMapper(nil)

	test1User := models.User{
		FirstName: "User1",
		LastName:  "Test1",
	}

	test2User := models.User{
		FirstName: "User2",
		LastName:  "Test2",
	}

	test1Group := models.Group{
		GroupName: "Test1",
		Kind:      models.Organization,
		Users:     make([]*models.User, 0, DefaultGroupMembersSize),
	}

	test2Group := models.Group{
		GroupName: "Test2",
		Kind:      models.ProjectGroup,
		Users:     make([]*models.User, 0, DefaultGroupMembersSize),
	}

	createdUser1, _ := users.Create(ctx, test1User)
	createdUser2, _ := users.Create(ctx, test2User)
	createdGroup1, _ := groups.Create(ctx, test1Group)
	createdGroup2, _ := groups.Create(ctx, test2Group)

	_ = users.JoinGroup(ctx, createdUser1, createdGroup1)
	_ = users.JoinGroup(ctx, createdUser1, createdGroup2)
	_ = users.JoinGroup(ctx, createdUser2, createdGroup2)

	_ = users.LeaveGroup(ctx, createdUser1, createdGroup2)

	_ = groups.AddMember(ctx, createdUser1, createdGroup2)
	_ = groups.AddMember(ctx, createdUser2, createdGroup1)

	_ = groups.RemoveMember(ctx, createdUser1, createdGroup2)
	_ = groups.RemoveMember(ctx, createdUser2, createdGroup1)

	_, _ = users.SearchUsersByName(ctx, "Test1")
	_, _ = users.SearchUsersByGroupName(ctx, "Test2")

	_, _ = groups.SearchGroupByName(ctx, "Test1")
	_, _ = groups.SearchGroupByMemberName(ctx, "Test2")

	<-ctx.Done()
	cancel()
}
