FROM golang:1.19.2 as build

WORKDIR /go/src/github.com/cyruzin/golang-study/docker

COPY . .

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' .

FROM alpine:latest  

RUN apk add ca-certificates

COPY --from=build /go/bin/docker .

EXPOSE 8000

CMD ["./docker"]