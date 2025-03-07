package response

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `binding:"required"     example:"01955908-d43b-7900-8f5c-5faa67dab4d3"            json:"id"`
	CreatedAt time.Time `binding:"required"     example:"1970-01-01T00:00:00.000+04:00"                   json:"createdAt"`
	Username  string    `binding:"required"     example:"john.doe"                                        json:"username"`
	Name      string    `binding:"required"     example:"John Doe"                                        json:"name"`
	Email     string    `binding:"required"     example:"john.doe@example.com"                            json:"email"`
	Avatar    string    `binding:"required,url" example:"https://avatars.githubusercontent.com/u/123456?" json:"avatar"`
}

type UserForEventReference struct {
	ID       uuid.UUID `binding:"required"     example:"01955908-d43b-7900-8f5c-5faa67dab4d3"            json:"id"`
	Username string    `binding:"required"     example:"john.doe"                                        json:"username"`
	Name     string    `binding:"required"     example:"John Doe"                                        json:"name"`
	Avatar   string    `binding:"required,url" example:"https://avatars.githubusercontent.com/u/123456?" json:"avatar"`
}
