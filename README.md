## Chapa Golang
Unofficial Golang SDK for Chapa ET API

### Todo:
- [ ] We could add nice validations on demand.


### Usage
##### 1. Installation
```
    go get github.com/vonmutinda/chapa
```

##### 2. Setup

```go
    package main

    import (
        "os"
        "github.com/vonmutinda/chapa"
    )

    func main(){
        chapaAPI := chapa.New(os.Getenv("CHAPA_API_KEY"))
    }
```

##### 3. Accept Payments
```go
    request := &chapa.ChapaPaymentRequest{
        Amount:         10,
        Currency:       "ETB",
        FirstName:      "Chapa",
        LastName:       "ET",
        Email:          "chapa@et.io",
        CallbackURL:    "https://posthere.io/e631-44fe-a19e",
        TransactionRef: faker.RandomString(20),
        Customization: map[string]interface{}{
            "title":       "A Unique Title",
            "description": "This a perfect description",
            "logo":        "https://your.logo",
        },
    }

    response, err := chapaAPI.PaymentRequest(request)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Printf("payment response: %+v\n", response)
```

##### 4. Verify Payment Transactions
```go
    response, err := chapa.Verify("your-txn-ref")
    if err != nil {
         fmt.Println(err)
         os.Exit(1)
    }

    fmt.Printf("verification response: %+v\n", response)
```

### Resources
- https://developer.chapa.co/docs/overview/

### Quirks
Suggestions on how to improve the API:
- Introduction of `status codes` would be a nice to have in the future. Status codes are better than the `message` in a way considering there are so many reasons a transaction could fail.
e.g 
```shell
    1001: Success
    4001: DuplicateTransaction
    4002: InvalidCurrency
    4003: InvalidAmount 
    4005: InsufficientBalance  
    5000: InternalServerError
    5001: GatewayError
    5002: RejectedByGateway
```
Juuust an example!
### Contributions
- Highly welcome
