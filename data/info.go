package data

type Brands struct {
	BrandId      int    `json:"brand_id"`
	BrandName    string `json:"brand_name"`
	BrandProduct []int  `json:"brand_product"`
}

type Products struct {
	ProductId     int    `json:"product_id"`
	ProductName   string `json:"product_name"`
	ProductBrand  Brands `json:"product_brand"`
	ProductPrice  int    `json:"product_price"`
	ProductStatus bool   `json:"product_status"`
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type brandDB map[int]Brands
type gearDB map[int]Products
type CredsDB map[string]string

var ProductList gearDB
var BrandList brandDB
var CredList CredsDB

func Productgen() {
	ProductList[1] = Products{
		ProductId:     1,
		ProductName:   "laptop",
		ProductBrand:  BrandList[1],
		ProductPrice:  100000,
		ProductStatus: true,
	}
	ProductList[2] = Products{
		ProductId:     2,
		ProductName:   "mobile",
		ProductBrand:  BrandList[1],
		ProductPrice:  10000,
		ProductStatus: true,
	}
	ProductList[3] = Products{
		ProductId:     3,
		ProductName:   "laptop",
		ProductBrand:  BrandList[2],
		ProductPrice:  200000,
		ProductStatus: true,
	}
	ProductList[4] = Products{
		ProductId:     4,
		ProductName:   "powerbank",
		ProductBrand:  BrandList[2],
		ProductPrice:  100000,
		ProductStatus: true,
	}
}

func Brandgen() {
	BrandList[1] = Brands{
		BrandId:      1,
		BrandName:    "apple",
		BrandProduct: []int{1, 2},
	}
	BrandList[2] = Brands{
		BrandId:      2,
		BrandName:    "samsung",
		BrandProduct: []int{3, 4},
	}
}

func GenCreds() {

	CredList["Ashraful"] = "12345"
	CredList["Neaj"] = "12345"
}

func Generator() {
	Brandgen()
	Productgen()
	GenCreds()
}

func init() {
	ProductList = make(gearDB)
	BrandList = make(brandDB)
	CredList = make(CredsDB)

}
