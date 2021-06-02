# REST APIs in GOlang to scrape an Amazon product web page given its URL, and store product details in MongoDB

## Description

This repo contains ready to use RESTfull APIs in GOlang.

 * REST_API_1 (POST) takes a Amazon product URL in payload and scrapes its webpage to find the details like name, image url, description, price and number of reviews related with that product, calls another SEPARATE API to store retrieved details in MongoDB.

 * REST_API_2 (POST) takes a JSON structure of product detail in request payload and store it along with timestamp in MongoDB (Database: "dummyStore", Collection:"productsDetails") to persist the information.

 * REST_API_1 uses [Colly](https://github.com/gocolly/colly) library for scraping.
 
 * Each service is dockerized as separate image.


### Project structure
```
.
├── api1
|   ├── REST_API_1.go
|   ├── Dockerfile
├── api2
|   ├── REST_API_2.go
|   └── Dockerfile
├── docker-compose.yml
└── README.md
```

## Installation

### Install Docker & Docker Compose

```bash
$ curl -sSL https://get.docker.com/ | sh
$ sudo pip install docker-compose
```

* See [Docker installation details](https://docs.docker.com/engine/install/).
* See [Docker Compose installation details](https://docs.docker.com/compose/install/).



### Execution and consuming services

* Download this repo in local system.
* Go inside the project folder where 'docker-compose.yml' file is present.
* Open in terminal.
* Run the services using `docker-compose up`.
* Use 'POST' on endpoint http://localhost:10005/getProductsDetailsFromURL and pass an Amazon product url in request payload. 

Example of request payload -
```
{
"url":"https://www.amazon.in/dp/B084YC1FHT/ref=s9_acsd_al_bw_c2_x_1_i?pf_rd_m=A1K21FY43GMZF8&pf_rd_s=merchandised-search-4&pf_rd_r=Q04XFWSNF4CCMZT4B38D&pf_rd_t=101&pf_rd_p=9614dae4-965a-4441-b222-9db02c06115b&pf_rd_i=22938665031"
}
```

Example of response -
```
{
    "url": "https://www.amazon.in/dp/B084YC1FHT/ref=s9_acsd_al_bw_c2_x_1_i?pf_rd_m=A1K21FY43GMZF8&pf_rd_s=merchandised-search-4&pf_rd_r=Q04XFWSNF4CCMZT4B38D&pf_rd_t=101&pf_rd_p=9614dae4-965a-4441-b222-9db02c06115b&pf_rd_i=22938665031",
    "product": {
        "name": "LG 190 L 4 Star Inverter Direct-Cool Single Door Refrigerator (GL-D201ASCY, Scarlet Charm, Base stand with Drawer)",
        "imageURL": "https://images-eu.ssl-images-amazon.com/images/I/41TzBxb3MrL._SX342_SY445_QL70_ML2_.jpg",
        "description": "Direct-cool refrigerator: Economical and Cooling without fluctuation\nCapacity 190 litres: Suitable for families with 2 to 3 members and bachelors\nEnergy Rating 4 Star: High energy efficiency\nManufacturer warranty: 1 year on product 10 years on compressor *T&C\nSmart inverter compressor: Unmatched performance great savings and super silent operation\nShelf type: Spill proof toughened glass\nInside box: 1 unit Refrigerator & 1 Unit user manual\nSpl. Features : Moist ‘n’ Fresh is an innovative lattice-patterned box cover which maintains the moisture at the optimal level | anti-bacterial gasket | Fastest Ice Making | Anti rat bite | Vegetable basket with 12.6 litres capacity | Tray egg | 2+3 Door basket (full size/half size)| Lock | Works without stabilizer: 90v ~ 310v | base stand with drawer | Solar Smart*\n",
        "price": "₹ 15,890.00",
        "totalReviews": 730
    }
}
```
* REST_API_2 will listen on port 10006 (endpoint - http://localhost:10006/createProductsDetails).


