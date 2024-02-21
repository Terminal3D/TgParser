package data

const (
	UpdateStatusNotAvailable = "not_available"
	UpdateStatusLowerPrice   = "lower_price"
	UpdateStatusHigherPrice  = "higher_price"
)

type UpdateInfo struct {
	Status        string
	Item          *ProductData
	PreviousPrice float64
	CurrentPrice  float64
}
