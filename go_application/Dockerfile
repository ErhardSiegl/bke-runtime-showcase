FROM golang:1.17.5-alpine3.15 as builder

COPY . /src/app
WORKDIR /src/app

RUN CGO_ENABLED=0 go build -v -o /main /src/app/main.go

########
FROM scratch

WORKDIR /app
COPY --from=builder /main /app/
EXPOSE 8080

CMD ["/app/main"]