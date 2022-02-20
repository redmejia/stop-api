package models

type Product struct {
	PID   int           `json:"id"`
	Name  string        `json:"name"`
	Sizes []interface{} `json:"sizes"`
	Price float64       `json:"price"`
	Image string        `json:"image"`
}

// Arrivals for carousel display
type NewArrivals struct {
	Key int    `json:"key"`
	Src string `json:"src"`
}
