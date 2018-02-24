# Digital Ocean Spaces Golang AWS SDK example

This is an example showing how to use the aws golang sdk
with digitalocean spaces. To use replace the variables
`endpoint` and `region` with your digitalocean endpoint
and `myBucket` with your D.O. bucket name.

View `main.go` and how the program works should be obvious.

The program when run with `go run main.go` will upload `main.go`
to your digital ocean space `myBucket`.

To install aws sdk `go get -u github.com/aws/aws-sdk-go`
Note: your API key for D.O. must be in ~/.aws/credentials
