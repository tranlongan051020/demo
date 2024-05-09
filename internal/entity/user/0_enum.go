package user

type UserRole uint8

const (
	Admin UserRole = iota
	Manager
	Viewer
)
