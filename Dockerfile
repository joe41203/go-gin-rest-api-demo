# builder
FROM golang:1.15.2-alpine as builder
RUN apk update && apk add git
WORKDIR /app
COPY ./go.mod ./go.sum /app/
COPY ./src/ /app/src/
RUN go mod download
RUN go build -o /api /app/src/main.go

# main
FROM alpine:3.11.12
EXPOSE 8080
COPY --from=builder /api .
USER nobody
ENTRYPOINT ["/api"]%