# Intro RESTful API

## Resume Materi

### API

Application Programming Interface (API) is a set functions and procedures allowing the creation of applications that access the features or data of an operating system, application, or other service.

### REST

REpresentational State Transfer (REST)

Use: HTTP Protocol

Request & Response Format:
- JSON
- XML
- SOAP

HTTP request method:
- GET
- POST
- PUT
- DELETE
- HEAD
- OPTION
- PATCH

### Package Go

2 Package go for consuming API:
1. net/http
    Provides HTTP Client and Server implementations
2. encoding/json
    Contains many functions for JSON operations:
    - Decode JSON to object struct
    - Decode JSON to map[string]interface{} & interface{}
    - Decode array JSON to array object
    - Encode object to JSON string

Example with consuming public API
```go
//  library/consumeAPI.go
package main
import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// A response struct to map the entire response
type StarWarsPeople struct {
    Name        string      `json:"name"`
    Height      string      `json:"height"`
    Mass        string      `json:"mass"`
    HairColor   string      `json:"hair_color"`
    SkinColor   string      `json:"skin_color"`
    EyeColor    string      `json:"eye_color"`
    BirthYear   string      `json:"birth_year"`
    Gender      string      `json:"string"`
    Films       []string    `json:"films"`
}

func main() {
    response, _ := http.Get("https://swapi.co/api/people/1")
    
    responseData, _ := ioutil.ReadAll(response.body)
    defer response.Body.Close()

    var LukeSkyWalker StarWarsPeople
    json.Unmarshal(responseData, &LukeSkyWalker)

    fmt.Println("---Print Field---")
    fmt.Println(LukeSkyWalker.Name)
    fmt.Println(LukeSkyWalker.Height)
    fmt.Println(LukeSkyWalker.Mass)
    fmt.Println(LukeSkyWalker.HairColor)
    fmt.Println(LukeSkyWalker.Films[0])
}
```