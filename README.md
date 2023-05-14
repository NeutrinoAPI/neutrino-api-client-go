# NeutrinoAPI Go Native SDK

Go client using the native HTTP client

The official API client and SDK built by [NeutrinoAPI](https://www.neutrinoapi.com/)

| Feature          |         |
|------------------|---------|
| Platform Version | >= 1.11 |
| HTTP Library     | Native  |
| JSON Library     | Native  |
| HTTP/2           | Yes     |
| HTTP/3           | No      |
| CodeGen Version  | 4.6.11  |

## Getting started

First you will need a user ID and API key pair: [SignUp](https://www.neutrinoapi.com/signup/)

## To Build 
```sh
# No build command necessary
```

## To Initialize 
```go
neutrinoAPIClient := neutrinoapi.NewNeutrinoAPIClient("<your-user-id>", "<your-api-key>")
```

## Running Examples

```sh
$ go run examples/bad_word_filter.go
```
You can find examples of all APIs in _/examples/<endpoint_name>/main.go_

## For Support 
[Contact us](https://www.neutrinoapi.com/contact-us/)