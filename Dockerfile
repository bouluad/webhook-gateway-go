FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o gateway ./cmd/server

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/gateway .
COPY jenkins_map.yaml .
EXPOSE 8080
CMD ["./gateway"]
