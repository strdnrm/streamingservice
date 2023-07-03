package model

type Delivery struct {
	OrderID string `db:"order_id" json:"-"`
	Name    string `db:"name" json:"name" validate:"required"`
	Phone   string `db:"phone" json:"phone" validate:"required"`
	Zip     string `db:"zip" json:"zip" validate:"required,max=10"`
	City    string `db:"city" json:"city" validate:"required"`
	Address string `db:"address" json:"address" validate:"required"`
	Region  string `db:"region" json:"region" validate:"required"`
	Email   string `db:"email" json:"email" validate:"required"`
}
