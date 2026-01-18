package types

import "time"

type NewBooking struct {
	Amount           float64   `json:"amount" validate:"required,gt=0"`
	UserID           int32     `json:"user_id" validate:"required,gt=0"`
	BookingPlatform  string    `json:"booking_platform" validate:"required,max=255"`
	FreeCancelBefore time.Time `json:"free_cancel_before" validate:"required"`
	BookingStart     time.Time `json:"booking_start" validate:"required,ltecsfield=BookingEnd"`
	BookingEnd       time.Time `json:"booking_end" validate:"required,gtecsfield=BookingStart"`
	Description      string    `json:"description" validate:"required,max=1000"`
}
