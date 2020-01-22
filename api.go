package routefusion

// Client specifies the abstraction for Routefusion APIs.
type Client interface {
	GetUser() (*UserDetails, error)
	UpdateUser(*User) (*UpdatedUserDetails, error)
	GetUserAdmin(subUserUUID string) (*AllUserDetails, error)
}
