package main

import (
    "fmt"
    "log"
    "context"
    "io/ioutil"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "time"

    // "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

)

type Product struct {
    Name string `json:"name"`
    ImageURL    string `json:"imageURL"`
    Description string `json:"description"`
    Price   string `json:"price"`
    TotalReviews      int `json:"totalReviews"`
}

type ProductDetails struct {
    Url   string `json:"url"`   
    Product Product `json:"product"`
    CreatedDate string `json:"createdDate"`
}

var Products []ProductDetails

// var HOST = "mongodb://localhost:27017"
var HOST = "mongodb://mongodb:27017"

// AddProduct adds a Product in the mongo DB
func AddProduct(product ProductDetails) bool {
 
    fmt.Println("AddProduct Called ::",product ,"\n")

    clientOptions := options.Client().ApplyURI(HOST)

    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
        return false
    }

    // Check the connection
    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
        return false
    }

    fmt.Println("Connected to MongoDB!")

    collection := client.Database("dummyStore").Collection("productsDetails")

    insertResult, err := collection.InsertOne(context.TODO(), product)
    if err != nil {
        log.Fatal(err)
        return false
    }

    fmt.Println("Inserted a single document RECORD ID : ", insertResult.InsertedID)
    return true
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage API 2 !")
    fmt.Println("Endpoint Hit: homePage")
}

// func returnAllProducts(w http.ResponseWriter, r *http.Request){
//     fmt.Println("Endpoint Hit: returnAllProducts")

//     clientOptions := options.Client().ApplyURI(HOST)

//     // Connect to MongoDB
//     client, err := mongo.Connect(context.TODO(), clientOptions)

//     if err != nil {
//         log.Fatal(err)
//         // return false
//     }

//     // Check the connection
//     err = client.Ping(context.TODO(), nil)

//     if err != nil {
//         log.Fatal(err)
//         // return false
//     }

//     fmt.Println("Connected to MongoDB!")

//     collection := client.Database("dummyStore").Collection("productsDetails")

//     cursor, err := collection.Find(context, bson.M{})
//     if err != nil {
//         log.Fatal(err)
//     }
//     var Products []bson.M
//     if err = cursor.All(context, &Products); err != nil {
//         log.Fatal(err)
//     }
//     fmt.Println(Products)
//     json.NewEncoder(w).Encode(Products)
// }

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    // myRouter.HandleFunc("/products", returnAllProducts)
    myRouter.HandleFunc("/createProductsDetails", createProductsDetails).Methods("POST")

    log.Fatal(http.ListenAndServe(":10006", myRouter))
}

func createProductsDetails(w http.ResponseWriter, r *http.Request) {
 
    var ProductDetail ProductDetails  
    fmt.Println("Endpoint Hit: createProductsDetails\n")

    reqBody, _ := ioutil.ReadAll(r.Body)
    json.Unmarshal(reqBody, &ProductDetail)

    fmt.Println(" Product details found in API 2 to store in MongoDB: ",ProductDetail,"\n")

    currentTime := time.Now()
 
    ProductDetail.CreatedDate = currentTime.String()

    success := AddProduct(ProductDetail) // FUNCTION CALL to add the product to the DB
    if !success {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    return

}

func main() {

    fmt.Println("SERVER STARTED :: SELLER APP API 2\n")
    handleRequests()

}