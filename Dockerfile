# BUILD STAGE
FROM golang:latest as build

# copy
WORKDIR /go/src/github.com/mchmarny/klogo/
COPY . /src/

# dependancies
WORKDIR /src/
ENV GO111MODULE=on
RUN go mod download

# build
WORKDIR /src/cmd/
RUN CGO_ENABLED=0 go build -v -o /klogo


# RUN STAGE
FROM alpine as release
RUN apk add --no-cache ca-certificates

# app executable
COPY --from=build /klogo /app/klogo

# run
WORKDIR /app/
ENTRYPOINT ["./klogo"]

