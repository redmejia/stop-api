package models

type Product struct {
	Name  string
	Sizes []int
	Price float64
	Image string
}

type Pants struct {
	Type    string
	PProduc []Product
}
