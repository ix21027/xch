FROM golang:1.22.3-alpine3.19 AS builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /src
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=1 GOOS=linux go build -o app -a -ldflags '-linkmode external -extldflags "-static"' .

FROM scratch
COPY --from=builder /src/app .
EXPOSE 9090

ENTRYPOINT ["/app"]
