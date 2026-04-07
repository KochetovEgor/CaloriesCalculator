FROM golang:1.26 AS builder
WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .
# RUN go build -o /build/calories_calculator cmd/backend/main.go

RUN --mount=type=bind,target=. \
    --mount=type=cache,id=go-mod,target=/go/pkg/mod \
    --mount=type=cache,id=go-cache,target=/root/.cache/go-build \
    go build -o /build/calories_calculator cmd/backend/main.go

FROM ubuntu
WORKDIR /app
COPY --from=builder /build/calories_calculator calories_calculator

EXPOSE 8000
CMD ["./calories_calculator"]