# Build stage
FROM golang:1.20-alpine AS builder

WORKDIR /src
COPY . .
RUN go build -o flugo cmd/main.go

# Run stage
FROM alpine
WORKDIR /src
COPY --from=builder /src/flugo .
COPY .env .
COPY ./migration ./migration

EXPOSE 3000
CMD [ "./flugo" ]