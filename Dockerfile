# -------- Stage 1: Download dependencies --------
FROM golang:1.24 as deps

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# -------- Stage 2: Build application --------
FROM golang:1.24 as builder

WORKDIR /app
COPY --from=deps /go/pkg /go/pkg
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /main ./cmd/app

# -------- Stage 3: Run the binary --------
FROM alpine:latest

WORKDIR /root/
RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /main .
COPY config ./config
COPY app.env .

RUN chmod +x ./main

EXPOSE 8080
CMD ["./main"]