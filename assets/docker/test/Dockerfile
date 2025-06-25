FROM golang:1.24.4
WORKDIR /go/src/myapp
COPY . .
CMD ["go", "test", "--tags=db", "./..."]
