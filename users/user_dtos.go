package users

type ApiGender string

const (
	Male   ApiGender = "male"
	Female ApiGender = "female"
)

func (g ApiGender) IsValid() bool {
	switch g {
	case Male, Female:
		return true
	default:
		return false
	}
}

type RegisterUserDto struct {
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Password  string    `json:"password" validate:"required,min=8,max=100"`
	BirthDate string    `json:"birth_date" validate:"required,datetime=2006-01-02"`
	Gender    ApiGender `json:"gender" validate:"required,oneof=male female"`
}

type UpdateUserDto struct {
	Name      string    `json:"name" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	BirthDate string    `json:"birth_date" validate:"required,datetime=2006-01-02"`
	Gender    ApiGender `json:"gender" validate:"required,oneof=male female"`
}

type LoginUserDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type LoginResponseDto struct {
	Token string `json:"token"`
}

type UserDto struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	BirthDate string    `json:"birth_date"`
	Gender    ApiGender `json:"gender"`
}
