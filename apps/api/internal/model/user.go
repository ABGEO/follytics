package model

type UserType string

const (
	UserTypeRegular   UserType = "REGULAR"
	UserTypeReference UserType = "REFERENCE"
)

type User struct {
	Base

	Type      UserType `gorm:"type:user_type;default:REGULAR"`
	GHID      int64    `gorm:"uniqueIndex"`
	Username  string   `gorm:"uniqueIndex"`
	Name      string
	Email     string
	Avatar    string
	Followers []*User `gorm:"many2many:user_followers"`
}
