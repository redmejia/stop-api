package api

import (
	"encoding/json"
	"net/http"

	"github.com/redmejia/stop/models"
)

// test data for product pants and shirts
var productPants = []models.Product{
	models.Product{
		PID:   1,
		Name:  "Pant1",
		Sizes: []interface{}{1, 2, 3, 41, 2, 3, 4},
		Price: 12.53,
		Image: "images/pants-unsplash.jpg",
	},
	models.Product{
		PID:   2,
		Name:  "Pant2",
		Sizes: []interface{}{1, 2, 3, 41, 2, 3, 4},
		Price: 12.53,
		Image: "images/pants-unsplash.jpg",
	},
	models.Product{
		PID:   3,
		Name:  "Pant3",
		Sizes: []interface{}{1, 2, 3, 41, 2, 3, 4},
		Price: 12.53,
		Image: "images/pants-unsplash.jpg",
	},
}

var productShirts = []models.Product{
	models.Product{
		PID:   1,
		Name:  "shirt1",
		Sizes: []interface{}{"xs", "s", "m"},
		Price: 12.53,
		Image: "images/seven-unsplash.jpg",
	},
	models.Product{
		PID:   2,
		Name:  "shirt2",
		Sizes: []interface{}{"xs", "s", "m"},
		Price: 12.53,
		Image: "images/seven-unsplash.jpg",
	},
	models.Product{
		PID:   3,
		Name:  "shirt3",
		Sizes: []interface{}{"xs", "s", "m"},
		Price: 12.53,
		Image: "images/seven-unsplash.jpg",
	},
}

// test data for new arrivals
var newArrivals = []models.NewArrivals{
	models.NewArrivals{
		Key: 1,
		Src: "http://127.0.0.1:8080/images/unplash_one.jpg",
	},
	models.NewArrivals{
		Key: 2,
		Src: "http://127.0.0.1:8080/images/unsplash_two.jpg",
	},
}

type product map[string][]models.Product
type arrivalsProduct map[string][]models.NewArrivals

func (a *Application) ProductPants(w http.ResponseWriter, r *http.Request) {
	var pants = make(product)
	pants["pants"] = productPants

	b, err := json.Marshal(&pants)
	if err != nil {
		a.Error.Fatal(err)
		return
	}

	a.Success.Println(string(b))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (a *Application) ProductShirt(w http.ResponseWriter, r *http.Request) {
	var shirts = make(product)
	shirts["shirts"] = productShirts

	b, err := json.Marshal(&shirts)
	if err != nil {
		a.Error.Fatal(err)
		return
	}

	a.Success.Println(string(b))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

func (a *Application) ProductArrivals(w http.ResponseWriter, r *http.Request) {
	var arrivals = make(arrivalsProduct)
	arrivals["arrivals"] = newArrivals

	b, err := json.Marshal(&arrivals)
	if err != nil {
		a.Error.Fatal(err)
		return
	}

	a.Success.Println(string(b))

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}
