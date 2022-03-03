package api

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/redmejia/stop/models"
)

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

	// postgresql array data type does not start at index 0 insted index 1
	// {'28', '30', '32'} = sizes[1], sizes[2], sizes[3]
	query := `
		SELECT 
			id, name, 
			sizes[1], sizes[2], 
			sizes[3], sizes[4], 
			price, image
		FROM 
			pants
			`

	scanRows := func(rows *sql.Rows) (models.Product, error) {

		var pants models.PantSizes

		// this data type will populate with the data from the database
		product := models.Product{
			PID:  0,
			Name: "",
			Sizes: []string{ // add more size
				pants.SizeOne,
				pants.SizeTwo,
				pants.SizeThree,
				pants.SizeFour,
			},
			Price: 0.0,
			Image: "",
		}

		// var product models.Product

		err := rows.Scan(
			&product.PID,
			&product.Name,
			&product.Sizes[0],
			&product.Sizes[1],
			&product.Sizes[2],
			&product.Sizes[3],
			&product.Price,
			&product.Image,
		)

		return product, err

	}

	products, err := a.DB.Get(query, scanRows)

	if err != nil {
		a.Error.Println(err)
	}

	a.Success.Println(products)

	var pants = make(product)
	pants["pants"] = products

	b, err := json.Marshal(&pants)
	if err != nil {
		a.Error.Println(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}

func (a *Application) ProductShirt(w http.ResponseWriter, r *http.Request) {

	// postgresql array data type does not start at index 0 insted index 1
	// {'28', '30', '32'} = sizes[1], sizes[2], sizes[3]
	query := `
		SELECT 
			id, name,
			sizes[1], sizes[2],
			sizes[3], price,
			image
		FROM
			shirts
	`
	scanRows := func(rows *sql.Rows) (models.Product, error) {

		var shirtSize models.ShirtSizes

		// this data type will populate with the data from the database
		product := models.Product{
			PID:  0,
			Name: "",
			Sizes: []string{ //add more size
				shirtSize.S,
				shirtSize.M,
				shirtSize.L,
			},
			Price: 0.0,
			Image: "",
		}

		err := rows.Scan(
			&product.PID,
			&product.Name,
			&product.Sizes[0],
			&product.Sizes[1],
			&product.Sizes[2],
			&product.Price,
			&product.Image,
		)

		return product, err

	}

	products, err := a.DB.Get(query, scanRows)
	if err != nil {
		a.Error.Println(err)
	}

	var shirts = make(product)
	shirts["shirts"] = products

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

// This method is use for the web stop applicatio
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
