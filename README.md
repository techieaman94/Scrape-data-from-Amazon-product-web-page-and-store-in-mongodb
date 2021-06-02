# REST APIs in Golang to scrape an Amazon product web page given its URL, and store product details in MongoDB

## Description

This repo contains ready to use RESTfull APIs in GOlang.

 * REST_API_1 (POST) takes a Amazon product URL in payload and scrapes its webpage to find the details like name, image url, description, price and number of reviews related with that product, calls another SEPARATE API to store retrieved details in MongoDB.

 * REST_API_2 (POST) takes a JSON structure of product detail in request payload and store it along with timestamp in MongoDB (Database: "dummyStore", Collection:"productsDetails") to persist the information.

## Getting Started

### Project structure
.
├── api1
|   ├── REST_API_1.go
|   ├── Dockerfile
├── api2
|   ├── REST_API_2.go
|   └── Dockerfile
├── docker-compose.yml
└── README

### Dependencies

* Describe any prerequisites, libraries, OS version, etc., needed before installing program.

## Installation

### Install Docker & Docker Compose

```bash
$ curl -sSL https://get.docker.com/ | sh
$ sudo pip install docker-compose
```

See [Docker installation details](https://docs.docker.com/engine/install/).
See [Docker Compose installation details](https://docs.docker.com/compose/install/).



### Execution and consuming services

* Download this repo in local system.
* Go inside the project folder where 'docker-compose.yml' file is present.
* Open in terminal.
* Run the application using `docker-compose up`.
* Use 'POST' on endpoint http://localhost:10005/getProductsDetailsFromURL and pass a Amazon product url in request payload. 

Example of request payload -
```
{
    "url":"https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC/"
}
```
* REST_API_2 will listen on port 10006 (endpoint - http://localhost:10006/createProductsDetails).


