# Aircall API wrapper for Go

This Go package provides a wrapper to interact with the [Aircall API](http://developer.aircall.io/)

## How to install

```
go get -u github.com/leoht/aircall
```

## Quick start

```go
package main

import (
    "fmt"
    "github.com/leoht/aircall"
)

func main() {
    // Create the client with your API id and secret
    client := aircall.NewClient("<your app ID>", "<your app secret")

    // Get current company info
    // Returns a CompanyResponse struct
    res, err := client.Company()

    if err != nil {
      panic(err)
    }

    company = res.Company
    fmt.Println("Company name", company.Name)

    // Get users list
    res, err := client.Users(aircall.Paginate{
        PerPage: 10
    })

    if err != nil {
      panic(err)
    }

    for _, u := range res.Users {
        fmt.Println("User", u.ID, u.Name)

        // Get specific user
        userRes, err := client.User(user.ID)
    }
    
    // Get numbers
    res, err := client.Numbers(aircall.Paginate{
        PerPage: 10,
        Page: 2
    })

    if err != nil {
      panic(err)
    }

    for _, num := range res.Numbers {
        fmt.Println("Number", num.ID, num.Name, num.Digits)
    }
}

```

## API Calls

Each API call function returns a specific struct type to hold the response data.
See [messages.go](messages.go) to view all possible response types, and [types.go](types.go) to view the possible entity types.

For all calls returning a lis of entities (e.g `Users()` or `Contacts()`), an attribute `Meta` will be present 
in the response struct. This attribute is a struct containing `Count`, `Total`, `CurrentPage`, `PerPage`, `NextPageLink` and `PreviousPageLink`.

### Get Current Company
```go
res, err := client.Company()
company := res.Company
fmt.Println("Company: ", company.Name)
```
### List Users
```go
// Default pagination parameters
res, err := client.Users(aircall.Paginate{})
for _, user := range res.Users {
    fmt.Println("User: ", user.ID)
}

// With pagination
res, err := client.Users(aircall.Paginate{
    Page: 1,
    PerPage: 5,
    Order: "asc"
})
```
### Get User Info
```go
res, err := client.User(123456)
user := res.User
fmt.Println("User: ", user.ID)
```
### List Phone Numbers
```go
res, err := client.Numbers(aircall.Paginate{})
for _, number := range res.Numbers {
    fmt.Println("Number: ", number.ID, number.Digits)
}

```
### Get Phone Number Info
```go
res, err := client.Number(153434)
number := res.Number
fmt.Println("Number: ", number.ID, number.Digits)
```
### List Calls
```go
res, err := client.Calls(aircall.Paginate{})
for _, call := range res.Calls {
    fmt.Println("Call: ", call.ID)
}
```
### Search Calls
```go
res, err := client.SearchCalls(
    aircall.Paginate{},
    aircall.Search{
        PhoneNumber: "+44163748849",
    },
)
```