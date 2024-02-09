package data

type ProductData struct {
	Name      string     `json:"name"`
	Brand     string     `json:"brand"`
	Available bool       `json:"available"`
	Price     float64    `json:"price"`
	Sizes     []SizeData `json:"sizes"`
}

type SizeData struct {
	Size     string `json:"size"`
	Quantity int    `json:"qty"`
}
