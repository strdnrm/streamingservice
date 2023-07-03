package model

type Item struct {
	OrderID     string `db:"order_id" json:"-"`
	ChrtID      int    `db:"chrt_id" json:"chrt_id" validate:"required"`
	TrackNumber string `db:"item_track_number" json:"track_number" validate:"required"`
	Price       int    `db:"price" json:"price" validate:"required,gt=0"`
	RID         string `db:"rid" json:"rid" validate:"required"`
	Name        string `db:"item_name" json:"name" validate:"required"`
	Sale        int    `db:"sale" json:"sale" validate:"required"`
	Size        string `db:"size" json:"size" validate:"required"`
	TotalPrice  int    `db:"total_price" json:"total_price" validate:"required,gt=0"`
	NmID        int    `db:"nm_id" json:"nm_id" validate:"required"`
	Brand       string `db:"brand" json:"brand" validate:"required"`
	Status      int    `db:"status" json:"status" validate:"required"`
}
