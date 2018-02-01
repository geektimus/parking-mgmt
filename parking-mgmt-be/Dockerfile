# build stage
FROM golang:1.9.3
WORKDIR /go/src/github.com/geektimus/parking-mgmt/parking-mgmt-be
COPY . .
RUN go get -d ./... && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o parking-be .

# final stage
FROM alpine:3.7
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/geektimus/parking-mgmt/parking-mgmt-be/parking-be .
ENTRYPOINT ./parking-be