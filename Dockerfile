FROM golang:1.26.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/main.go

FROM alpine:3.20

WORKDIR /app
ENV POLICY_PATH=./policy/auth.rego
ENV PORT=8080

COPY --from=builder /app/server .
COPY --from=builder /app/policy ./policy

EXPOSE 8080

CMD ["./server"]