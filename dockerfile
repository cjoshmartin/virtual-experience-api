FROM golang:1.16.5-alpine as builder

RUN apk --no-cache add ca-certificates git

WORKDIR /build/api

COPY go.mod ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o api

# post build stage

FROM alpine

WORKDIR /root

COPY --from=builder /build/api/api .

EXPOSE 8080

CMD ["./api"]

# https://medium.com/@She_Daddy/building-a-simple-mongo-api-service-with-go-eedc3af5ac99