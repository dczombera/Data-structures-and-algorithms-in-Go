FROM golang:alpine as builder
COPY . $GOPATH/src/github.com/dczombera/data-structures-and-algorithms-in-go 
WORKDIR $GOPATH/src/github.com/dczombera/data-structures-and-algorithms-in-go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /go/bin/app .

FROM scratch
WORKDIR /app
COPY --from=builder /go/bin/app . 
CMD ["./app"]