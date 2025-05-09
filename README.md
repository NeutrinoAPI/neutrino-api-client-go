# NeutrinoAPI Go Native SDK

Neutrino API Go client using the native HTTP library

The official API client and SDK built by [NeutrinoAPI](https://www.neutrinoapi.com/)

| Feature          |         |
|------------------|---------|
| Platform Version | >= 1.11 |
| HTTP Library     | Native  |
| JSON Library     | Native  |
| HTTP/2           | Yes     |
| HTTP/3           | No      |
| CodeGen Version  | 4.7.1   |

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

Set the __'your-user-id'__ and __'your-api-key'__ values in the example to retrieve real API responses

## For Support 
[Contact us](https://www.neutrinoapi.com/contact-us/)