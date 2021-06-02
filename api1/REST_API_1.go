package main

import (
    "fmt"
    "log"
    "bytes"
    "io/ioutil"
    "strings"
    "strconv"
    "regexp"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/gocolly/colly"
)

type Product struct {
    Name string `json:"name"`
    ImageURL    string `json:"imageURL"`
    Description string `json:"description"`
    Price   string `json:"price"`
    TotalReviews      int `json:"totalReviews"`
}

type ProductDetails struct {
    Url string `json:"url"`   
    Product Product `json:"product"`
}

type  URLType struct {
    Url string `json:"Url"`
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage API 1 !")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/getProductsDetailsFromURL", getProductsDetails).Methods("POST")
    log.Fatal(http.ListenAndServe(":10005", myRouter))
}

func getProductsDetails(w http.ResponseWriter, r *http.Request) {
 
    var ProductDetail ProductDetails  
    fmt.Println("Endpoint Hit: getProductsDetails\n")

    reqBody, _ := ioutil.ReadAll(r.Body)
    var data URLType
    json.Unmarshal(reqBody, &data)

    fmt.Println(" URL in req body : ",data.Url,"\n")

    var Name,Price,RatingsNumber,Image_url, ProductDescription string

    c := colly.NewCollector()

    c.OnHTML("#productTitle", func(e *colly.HTMLElement) {
        Name = e.Text
        Name = strings.Trim(Name, "\n")
    })

    c.OnHTML("#acrCustomerReviewText", func(e *colly.HTMLElement) {
        RatingsNumber = e.Text
        RatingsNumber = strings.ReplaceAll(RatingsNumber, ",", "")

        re := regexp.MustCompile("[0-9]+")
        RatingsNumber = re.FindAllString(RatingsNumber, -1)[0]
    })

    c.OnHTML("#priceblock_ourprice", func(e *colly.HTMLElement) {
        Price = e.Text
        // fmt.Println("Price found --------->",Price)
    })

    c.OnHTML("#landingImage", func(e *colly.HTMLElement) {
        Image_url = e.Attr("src")
    
    })

    c.OnHTML("#feature-bullets", func(e *colly.HTMLElement) {
        var desc string
        ProductDescription = ""

        e.ForEach("li", func(_ int, e *colly.HTMLElement) {
            id := e.Attr("id")
    
            if id != "replacementPartsFitmentBullet"  {

                desc = e.ChildText("span.a-list-item")
                desc = strings.ReplaceAll(desc, ",", "")
                ProductDescription = ProductDescription + desc+"\n"
            }
        })
    })

    c.OnRequest(func(r *colly.Request) {
        fmt.Println("Visiting", r.URL.String(),"\n")
    })

    c.Visit(data.Url)

    fmt.Println(" Product name found usind COLLY : ", Name)
    fmt.Println(" Description found usind COLLY : ", ProductDescription)
    fmt.Println(" Price found usind COLLY : ",Price)
    fmt.Println(" Number of ratings found usind COLLY : ",RatingsNumber)
    fmt.Println(" Product image url found usind COLLY: ",Image_url)

    ProductDetail.Url = data.Url
    ProductDetail.Product.Name = Name
    ProductDetail.Product.ImageURL = Image_url
    ProductDetail.Product.Description = ProductDescription
    ProductDetail.Product.Price = Price
    num, err := strconv.Atoi(RatingsNumber)
    if err != nil {
        // handle error
        // fmt.Println(err)
    }

    ProductDetail.Product.TotalReviews = num
    fmt.Println(" Product details found in API 1 after SCRAPPING : ",ProductDetail,"\n")

    json.NewEncoder(w).Encode(ProductDetail)
    // --------------- add product to mongodb by calling API 2 ----------

    // url := "http://localhost:10006/createProductsDetails"
    url := "http://go-api-2:10006/createProductsDetails"


    payloadBuf := new(bytes.Buffer)
    json.NewEncoder(payloadBuf).Encode(ProductDetail)

    req, _ := http.NewRequest("POST", url, payloadBuf )
    req.Header.Add("Content-Type", "application/json")
    response, err := http.DefaultClient.Do(req)

    if err != nil {
    fmt.Println("HTTP call failed:", err)
    }

    defer response.Body.Close()

    statusOK := response.StatusCode >= 200 && response.StatusCode < 300
    if !statusOK {
        fmt.Println("Non-OK HTTP status:", response.StatusCode)
    }

}

func main() {
    fmt.Println("SERVER STARTED :: SELLER APP API 1\n")

    handleRequests()
}