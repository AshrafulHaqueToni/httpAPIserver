## RESTful HTTP API server using [Go](https://github.com/golang), [Cobra CLI](https://github.com/spf13/cobra), [Go-chi](https://github.com/go-chi/chi)

### Description

This is a basic RESTful API server, build with Golang. In this API server I have implemented Cobra CLI for running the API from the CLI and also used go-chi instead of Go net/http.


[![Go Report Card](https://goreportcard.com/badge/github.com/AshrafulHaqueToni/httpAPIserver)](https://goreportcard.com/report/github.com/AshrafulHaqueToni/httpAPIserver)
------------ 

### Installation

- `git clone https://github.com/AshrafulHaqueToni/httpAPIserver.git`
- `cd httpAPIserver`
- `go install ./httpAPIserver`

---------------

### Run by CLI Commands

- start the API in default port : 8081 by `httpAPIserver start`
- start the API in your given port by `httpAPIserver start -p=8088`, give your port number in the place of 8088

--------------

### Run the API server in docker container using dockerfile

#### Create docker image from the dockerfile

- `docker build -t <image_name> .`
- or `docker build -t <docker_hub_username>/<image_name>:<tag> .` (if your do this then don't need to give tag before dockerhub push)

#### Run the API server from the docker image in docker container

- `docker run -p 8088:8080 <image_name>` (valid when used `CMD ["start", "-p", "8080"]` in Dockerfile)
- `docker run -p 8081:8088 <image_name> start -p "8088"` (valid for current version of Dockerfile, when did not used CMD in Dockerfile)

--------------

#### upload the image to [Docker Hub](https://hub.docker.com/)

- `docker login --username=<docker_hub_username>`
- `docker tag <id_of_the_created_image> <docker_hub_username>/<name_of_the_image>:<tag>`
- `docker push <docker_hub_username>/<name_of_the_image>:<tag>`

--------------

#### run using volume (valid for current version where did not gave .env file in docker image)


- `docker run -v <absolute_host_path/.env>:<container_path/.env> -p 8088:8089 <image_name> start -p 8089`


--------------

### The Endpoints of this REST API

| Endpoint                | Function        | Method | StatusCode                                    | Authentication |
|-------------------------|-----------------|--------|-----------------------------------------------|----------------|
| `/login`                | logIn           | GET    | StatusOK, StatusUnauthorized                  | Basic          |
| `/products`             | ShowAllProducts | GET    | StatusOK, StatusUnauthorized                  | JWT            |
| `/brands`               | ShowAllBrands   | GET    | StatusOK, StatusUnauthorized                  | JWT            |
| `/products/{id}`        | GetProducts     | GET    | StatusOK, StatusNoContent, StatusUnauthorized | JWT            |
| `/brands/{id}`          | GetBrands       | GET    | StatusOK, StatusNoContent, StatusUnauthorized | JWT            |
| `/products/{id}`        | DeleteProduct   | DELETE | StatusOK, StatusNoContent, StatusUnauthorized | JWT            |
| `/products/addproduct`  | AddProducts     | POST   | StatusOK, StatusUnauthorized                  | JWT            |
| `/products/update/{id}` | UpdateProducts  | POST   | StatusOK, StatusNoContent,StatusUnauthorized  | JWT            |
| `/brands/delete/{id}`   | DeleteBrand     | DELETE | StatusOK, StatusUnauthorized                  | JWT            |
| `/brands/addbrand`      | AddBrands       | POST   | StatusOK, StatusUnauthorized                  | JWT            |
| `/logout`               | Logout          | POST   | StatusOK, StatusUnauthorized                  | JWT            |

----------------

### Data Model

* Brand Model
```
    type Brands struct {
        BrandId      int    `json:"brand_id"`
        BrandName    string `json:"brand_name"`
        BrandProduct []int  `json:"brand_product"`
    }
```

* Product Model
```
    type Products struct {
        ProductId     int    `json:"product_id"`
        ProductName   string `json:"product_name"`
        ProductBrand  Brands `json:"product_brand"`
        ProductPrice  int    `json:"product_price"`
        ProductStatus bool   `json:"product_status"`
    }

```

* Credentials Model
```
   type Credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

```

----------------

### JWT Authentication

- implemented JWT authentication
- first of all user need to hit `/login` endpoint with basic authentication then a token will be given and with that token for specific time user can do other request
----------------

#### Run the API server

- `curl -X POST -H "Content-Type:application/json" -d '{"username":"Ashraful","password":"12345"}' http://localhost:8081/login`

#### Get all products

- `curl -X GET -H "Authorization: Bearer ${TOKEN}" http://localhost:8081/products`

#### Get single product

- `curl -X GET -H "Authorization: Bearer ${TOKEN}" http://localhost:8081/products/1`

#### Get all brands

- `curl -X GET -H "Authorization: Bearer ${TOKEN}" http://localhost:8081/brands`

#### Get single brand

- `curl -X GET -H "Authorization: Bearer ${TOKEN}" http://localhost:8081/brands/1`

#### add new product

```
curl -X POST -H "Content-Type:application/json" -d '{"product_id":5,"product_name":"keyboard","product_brand":{"brand_id":3,"brand_name":"A4tech","brand_product":[]}}' -H "Authorization: Bearer ${TOKEN}" http://localhost:8081/products/addproduct
```

#### add new brand

```
curl -X POST -H "Content-Type:application/json" -d '{"brand_id":4,"brand_name":"Dell","brand_product":[]}' -H "Authorization: Bearer ${TOKEN}" http://localhost:8081/brands/addbrand
```

#### Update any product

```
curl -X POST -H "Content-Type:application/json" -d '{"product_id":4,"product_name":"mobile","product_brand":{"brand_id":2,"brand_name":"apple","brand_product":[3,4]}}' -H "Authorization: Bearer ${TOKEN}" http://localhost:8081/products/update/4
```

#### Delete a product
```
curl -X DELETE -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://localhost:8081/products/delete/1`
```

#### Delete a brand
```
curl -X DELETE -H 'Accept: application/json' -H "Authorization: Bearer ${TOKEN}" http://localhost:8081/brands/delete/1`
```


----------------

### API Endpoints Testing

- Primarily tested the API endpoints by [Postman](https://github.com/postmanlabs)
- E2E Testing.
    - Checked response status code with our expected status code
