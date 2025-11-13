package types

type CreateHouseholdRequest struct {
	Name    string  `json:"name" validate:"required,min=1,max=100"`
	Address *string `json:"address,omitempty" validate:"omitempty,min=1,max=100"` // Optional field
}
