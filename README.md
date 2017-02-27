# Aircall API wrapper for Go

This Go package provides a wrapper to interact with the [Aircall API](http://developer.aircall.io/)

## How to install

```
go get -u github.com/leoht/aircall
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/leoht/aircall"
)

func main() {
    client := aircall.NewClient("<your app ID>", "<your app secret")

    // Get current company info
    res, err := client.Company()

    if err != nil {
      panic(err)
    }

    company = res.Company
    fmt.Println("Company name", company.Name)

    // Get users list
    res, err := client.Users()

    if err != nil {
      panic(err)
    }

    for _, u := range res.Users {
        fmt.Println("User", u.ID, u.Name)

        // Get specific user
        userRes, err := client.User(user.ID)
    }
    
    // Get numbers
    res, err := client.Numbers()

    if err != nil {
      panic(err)
    }

    for _, num := range res.Numbers {
        fmt.Println("Number", num.ID, num.Name, num.Digits)
    }
}

```
