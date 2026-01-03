# Build stage
FROM golang:1.24-bookworm AS build
RUN apt-get update && apt-get install -y --no-install-recommends \
    git build-essential ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Install templ
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Generate templ code (no-op if already generated)
RUN templ generate

# Build static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Runtime stage
FROM gcr.io/distroless/base-debian12
WORKDIR /app
COPY --from=build /app/server /app/server
# copy runtime assets if needed
COPY --from=build /app/static /app/static
COPY --from=build /app/views /app/views
COPY --from=build /app/posts /app/posts
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app/server"]
