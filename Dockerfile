FROM  golang:latest as builder

WORKDIR /go/src/github.com/fearful-symmetry/hellologdriver
COPY . .

ARG GOOS=linux
ARG GOARCH=amd64
ARG GOARM=

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hellologdriver


FROM alpine:3.7 as final
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/src/github.com/fearful-symmetry/hellologdriver/hellologdriver /usr/bin/hellologdriver