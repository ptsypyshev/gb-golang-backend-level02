package models

const (
	ProjectGroup = iota
	Organization
	CorpGroup
	Community
)

type Group struct {
	ID int
	GroupName string
	Kind int
	Users []*User
}
