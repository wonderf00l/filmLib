package role

type Role uint8

const (
	Administrator Role = iota + 1
	RegularUser
)

var RoleMap = map[Role]string{
	Administrator: "Administrator",
	RegularUser:   "RegularUser",
}
