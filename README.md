# Chapa-Golang
Unofficial Golang SDK for Chapa ET API

## Todo:
- [ ] We could add nice validations on demand.


## Quirks
- Failure to set `Content-Type` to `application/json` returns `Invalid currency` error
- Failure to provide `Callback URL` returns an error.
- The `checkout url` might expire before the user is done keying in the credit card details. This happens when the user attempts a payment with the wrong details at first, the second attempt with the correct details throws `session expired` error.
- Amount fields in `verification response` object should be in `float` and not `string`.
- Suggestion: Introduction of response codes could be a great way of summarizing the transaction response. 
