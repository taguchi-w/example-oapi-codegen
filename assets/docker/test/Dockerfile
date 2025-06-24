FROM golang:1.24.4
WORKDIR /go/src/myapp
COPY . .
RUN sudo chown -R circleci:circleci /go/src/myapp
CMD ["go", "test", "--tags=db", "./..."]
