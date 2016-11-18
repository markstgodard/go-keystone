# go-keystone

Go library for Keystone v3.0 API

### Installation

Install using `go get github.com/markstgodard/go-keystone`.


### Usage

```go
// create new client
client, err := keystone.NewClient("http://192.168.56.101:5000")
if err != nil {
    log.Fatal(err)
}

// get token
auth := keystone.NewAuth("admin", "password1", "Default")
token, err := client.Tokens(auth)
if err != nil {
    log.Fatal(err)
}
```
