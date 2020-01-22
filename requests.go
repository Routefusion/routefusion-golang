package routefusion

// User represents the changeable details pertaining to a user.
// TODO: QUESTION- Is this a multipart or marshalled http body?
type User struct {
	UserName    string
	Password    string
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Country     string
	CompanyName string
}
