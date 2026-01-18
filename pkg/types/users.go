package types

type NewUser struct {
	Email       string  `json:"email" validate:"required,email"`
	Username    string  `json:"username" validate:"required,min=1,max=100"`
	Lastname    string  `json:"lastname" validate:"required,min=1,max=100"`
	Gender      string  `json:"gender" validate:"required,oneof=M F m f"`
	Role        *string `json:"role,omitempty" validate:"omitempty,oneof=admin user ADMIN USER"`
	HouseholdId int32   `json:"household_id" validate:"required,min=1"`
}
