FROM golang:1.15.8-alpine3.13 as builder

CMD /bin/sh

COPY go.mod go.sum /go/src/fibonacci/

WORKDIR /go/src/fibonacci/

RUN go mod download

COPY . /go/src/fibonacci/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fibonacci /go/src/fibonacci

FROM alpine

RUN apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=builder /go/src/fibonacci/ /usr/bin/fibonacci

EXPOSE 8080 8080

WORKDIR /usr/bin/fibonacci/

ENTRYPOINT ["/usr/bin/fibonacci/fibonacci"]
