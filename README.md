# storing
Take care of writing and reading on AWS S3 and others storage providers

## Environment setting

To use with AWS it is necessary to set the following environment variables:
```shell
export AWS_ACCESS_KEY_ID=XXXXXXXXX
export AWS_SECRET_ACCESS_KEY=XXXXXXXXX
export AWS_REGION=XXXXXXXXX
export AWS_BUCKET=XXXXXXXXX
export AWS_ACL=XXXXXXXXX
```

## Installation

```shell
go get github.com/materialagora/storing
```

## Examples

### Upload
```go
// read file to upload
bUp, err := ioutil.ReadFile("../testdata/testfile.txt")
if err != nil {
	panic(err)
}

// Upload
var path string
path, err = Upload("testfile.txt", "text/plain", bUp)
if err != nil {
	panic(err)
}
println(path)
```

### Download
```go
bDown, err := Download("testfile.txt")
if err != nil {
	panic(err)
}
println(string(bDown))
```

### Delete
```go
err := Delete("testfile.txt")
if err != nil {
	panic(err)
}
```
