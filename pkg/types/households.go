package types

type CreateHouseholdRequest struct {
	Name    string  `json:"name"`
	Address *string `json:"address,omitempty"` // Optional field
}
