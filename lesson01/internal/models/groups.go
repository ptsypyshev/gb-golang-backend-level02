package models

const (
	ProjectGroup = iota
	Organization
	CorpGroup
	Community
)

type Group struct {
	ID int
	kind int
	users []User
}
