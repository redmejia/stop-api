package api

import (
	"encoding/json"
	"net/http"

	"github.com/redmejia/stop/models"
)

// test data
var productPants = []models.Product{
	models.Product{
		Name:  "Pant1",
		Sizes: []interface{}{1, 2, 3, 41, 2, 3, 4},
		Price: 12.53,
		Image: "no image",
	},
	models.Product{
		Name:  "Pant2",
		Sizes: []interface{}{1, 2, 3, 41, 2, 3, 4},
		Price: 12.53,
		Image: "no image",
	},
	models.Product{
		Name:  "Pant3",
		Sizes: []interface{}{1, 2, 3, 41, 2, 3, 4},
		Price: 12.53,
		Image: "no image",
	},
}

var productShirts = []models.Product{
	models.Product{
		Name:  "shirt1",
		Sizes: []interface{}{"xs", "s", "m"},
		Price: 12.53,
		Image: "no image",
	},
	models.Product{
		Name:  "shirt2",
		Sizes: []interface{}{"xs", "s", "m"},
		Price: 12.53,
		Image: "no image",
	},
	models.Product{
		Name:  "shirt3",
		Sizes: []interface{}{"xs", "s", "m"},
		Price: 12.53,
		Image: "no image",
	},
}

type data map[string][]models.Product

func (a *Application) ProductPants(w http.ResponseWriter, r *http.Request) {
	var pants = make(data)
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
	var shirts = make(data)
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
