/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package api

import (
	"encoding/json"
	"fmt"
	"github.com/AshrafulHaqueToni/httpAPIserver/data"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwa"
	"net/http"
	"strconv"
	"time"
)

var jwtkey = []byte("ThisIsSecretKey")
var tokenAuth *jwtauth.JWTAuth
var tokenString string
var token jwt.Token

func ShowAllBrands(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data.BrandList)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func GetBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	brandid := chi.URLParam(r, "brand_id")

	id, _ := strconv.Atoi(brandid)

	if _, ok := data.BrandList[id]; !ok {
		w.WriteHeader(404)
		return
	}
	err := json.NewEncoder(w).Encode(data.BrandList[id])

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func DeleteBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	brandid := chi.URLParam(r, "brand_id")

	id, _ := strconv.Atoi(brandid)

	if len(data.BrandList[id].BrandProduct) > 0 {
		w.Write([]byte("This brand has product"))
		return
	}
	delete(data.BrandList, id)

	err := json.NewEncoder(w).Encode(data.BrandList)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func AddBrands(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newBrand data.Brands

	err := json.NewDecoder(r.Body).Decode(&newBrand)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	newBrand.BrandId = GetBrandId()
	data.BrandList[newBrand.BrandId] = newBrand

	err = json.NewEncoder(w).Encode(data.BrandList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func Remove(tmp []int, val int) []int {

	for i := 0; i < len(tmp); i++ {
		if tmp[i] == val {
			tmp = append(tmp[:i], tmp[i+1:]...)
			break
		}
	}

	return tmp
}

func ShowAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data.ProductList)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productid := chi.URLParam(r, "product_id")

	id, _ := strconv.Atoi(productid)

	if _, ok := data.ProductList[id]; !ok {
		w.WriteHeader(404)
		return
	}
	err := json.NewEncoder(w).Encode(data.ProductList[id])

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productid := chi.URLParam(r, "product_id")
	id, _ := strconv.Atoi(productid)

	if _, ok := data.ProductList[id]; !ok {
		w.WriteHeader(404)
		return
	}

	tmpBrand := data.ProductList[id].ProductBrand
	tmpBrand.BrandProduct = Remove(tmpBrand.BrandProduct, id)

	delete(data.ProductList, id)

	err := json.NewEncoder(w).Encode(data.ProductList)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func GetProductId() int {
	avai := 1
	for value, _ := range data.ProductList {
		//fmt.Println(value)
		if value > avai {
			break
		}
		avai++
	}
	return avai
}

func GetBrandId() int {
	avai := 1
	for value, _ := range data.BrandList {
		if value > avai {
			break
		}
		avai++
	}
	return avai
}

/*
		{
	    "product_id": 2,
	    "product_name": "keyboard",
	    "product_brand": {
	        "brand_id": 1,
	        "brand_name": "apple",
	        "brand_product": [
	            1,
	            2
	        ]
	    },
	    "product_price": 1000,
	    "product_status": true
		}
*/
func AddProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newProduct data.Products

	err := json.NewDecoder(r.Body).Decode(&newProduct)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	tmpbrand := newProduct.ProductBrand

	newProduct.ProductId = GetProductId()

	//fmt.Println(newProduct.ProductId)

	f := false

	/*for key, value := range data.BrandList {
		if tmpbrand == value {
			f = true
			data.BrandList[key].BrandProduct = append(data.BrandList[key].BrandProduct, newProduct.ProductId)
			newProduct.ProductBrand = data.BrandList[key]
			break
		}
	}*/

	if f == false {
		newBrand := data.Brands{
			BrandId:      1,
			BrandName:    tmpbrand.BrandName,
			BrandProduct: []int{newProduct.ProductId},
		}
		newBrand.BrandId = GetBrandId()
		data.BrandList[newBrand.BrandId] = newBrand
		newProduct.ProductBrand = newBrand
	}

	data.ProductList[newProduct.ProductId] = newProduct
	err = json.NewEncoder(w).Encode(data.ProductList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return

	}

	w.WriteHeader(http.StatusOK)

}

/*
{
    "product_id": 1,
    "product_name": "laptop1",
    "product_brand": {
        "brand_id": 1,
        "brand_name": "apple",
        "brand_product": [
            1,
            2
        ]
    },
    "product_price": 100000,
    "product_status": true
}
*/

func UpdateProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	productid := chi.URLParam(r, "product_id")

	id, _ := strconv.Atoi(productid)

	if _, ok := data.ProductList[id]; !ok {
		w.WriteHeader(404)
		return
	}
	newProduct := data.ProductList[id]
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	data.ProductList[id] = newProduct
	err = json.NewEncoder(w).Encode(data.ProductList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func Login(w http.ResponseWriter, r *http.Request) {

	var creds data.Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	correctPassword, ok := data.CredList[creds.Username]

	if !ok || creds.Password != correctPassword {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expireTime := time.Now().Add(30 * time.Minute)

	//fmt.Println(expireTime)

	_, tokenString, err := tokenAuth.Encode(map[string]interface{}{
		"aud": "Ashraful",
		"exp": expireTime.Unix(),
	})

	//fmt.Println("check")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   tokenString,
		Expires: expireTime,
	})
	w.WriteHeader(http.StatusOK)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Expires: time.Now(),
	})
	w.WriteHeader(http.StatusOK)
}

func StartAPI(Port string) {

	//fmt.Println("check")
	data.Generator()
	tokenAuth = jwtauth.New(string(jwa.HS256), jwtkey, nil)
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Post("/login", Login)

	r.Group(func(r chi.Router) {
		//fmt.Println("check")
		// Details - https://github.com/go-chi/jwtauth/blob/master/jwtauth.go
		// Seek, verify and validate JWT tokens
		r.Use(jwtauth.Verifier(tokenAuth))
		// Handle valid / invalid tokens. In this example, we use
		// the provided authenticator middleware, but you can write your
		// own very easily, look at the Authenticator method in jwtauth.go
		// and tweak it, its not scary.
		r.Use(jwtauth.Authenticator)

		//fmt.Println("check")

		r.Route("/products", func(r chi.Router) {
			r.Get("/", ShowAllProducts)
			r.Get("/{product_id}", GetProduct)
			r.Delete("/delete/{product_id}", DeleteProduct)
			r.Post("/addproduct", AddProducts)
			r.Post("/update/{product_id}", UpdateProducts)
		})
		r.Route("/brands", func(r chi.Router) {
			r.Get("/", ShowAllBrands)
			r.Get("/{brand_id}", GetBrand)
			r.Delete("/delete/{brand_id}", DeleteBrand)
			r.Post("/addbrand", AddBrands)
		})

		r.Post("/logout", Logout)
	})
	port := ":" + Port

	fmt.Println(port)

	http.ListenAndServe(port, r)
}
