FROM golang:1.12 as builder

WORKDIR /src
ADD go.mod /src
ADD main.go /src
ADD handlers /src/handlers

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags="-w -s" -o /go/bin/app

FROM scratch

COPY --from=builder /go/bin/app /go/bin/app
ADD movie.mp4 /
ADD passwd /etc/passwd

USER 1000
EXPOSE 8080

ENTRYPOINT ["/go/bin/app"]




