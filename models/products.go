package models

type Product struct {
	PID   int      `json:"id"`
	Name  string   `json:"name"`
	Sizes []string `json:"sizes"`
	Price float64  `json:"price"`
	Image string   `json:"image"`
}

// Update with the products sizes
// Pants sizes
type PantSizes struct {
	SizeOne   string
	SizeTwo   string
	SizeThree string
	SizeFour  string
}

// Shirts sizes
type ShirtSizes struct {
	S string
	M string
	L string
}

// Arrivals for carousel display
type NewArrivals struct {
	Key int    `json:"key"`
	Src string `json:"src"`
}
