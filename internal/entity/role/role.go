package role

type Role uint8

const (
	Administrator Role = iota + 1
	RegularUser
)
