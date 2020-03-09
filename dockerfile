FROM golang:alpine3.11

RUN mkdir -p /app
WORKDIR /go/src/build_cv

COPY . .
RUN apk add --no-cache ca-certificates git
RUN go mod download

WORKDIR /go/src/build_cv/cmd
RUN GOOS=linux GOARCH=amd64 go build -o cv

EXPOSE 5000

RUN mv cv ../../../../app/
RUN rm -rf /../../../../build_cv
ENTRYPOINT ["/app/cv"]