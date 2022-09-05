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

- Failure to set `Content-Type` to `application/json` returns `Invalid currency` error
- Failure to provide `Callback URL` returns an error.
- The `checkout url` might expire before the user is done keying in the credit card details. This happens when the user attempts a payment with the wrong details at first, the second attempt with the correct details throws `session expired` error.
- Amount fields in `verification response` object should be in `float` and not `string`.
- Suggestion: Introduction of response codes could be a great way of summarizing the transaction response.

### Contributions
- Highly welcome
