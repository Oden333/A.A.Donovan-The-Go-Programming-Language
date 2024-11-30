package types

type Action string

var (
	Create Action = "create"
	Read   Action = "read"
	Update Action = "update"
	Delete Action = "delete"
)
