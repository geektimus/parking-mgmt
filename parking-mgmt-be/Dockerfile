# build stage
FROM golang:1.9.3 AS build-env

ADD . /go/src/github.com/geektimus/parking-mgmt/parking-mgmt-be
WORKDIR /go/src/github.com/geektimus/parking-mgmt/parking-mgmt-be
RUN go get ./... && GOOS=linux go build -o app .

# final stage
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build-env /go/src/github.com/geektimus/parking-mgmt/parking-mgmt-be/app /app/
ENTRYPOINT ./app