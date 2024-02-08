package data

type ProductData struct {
	Name      string   `json:"name"`
	Available bool     `json:"available"`
	Price     float32  `json:"price"`
	Sizes     []string `json:"sizes"`
}
