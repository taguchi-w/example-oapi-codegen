FROM cimg/go:1.23.1
WORKDIR /go/src/myapp
COPY . .
RUN sudo chown -R circleci:circleci /go/src/myapp
CMD ["go", "test", "--tags=db", "./..."]
