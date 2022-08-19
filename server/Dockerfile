FROM golang:1.19-alpine as dependencies
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

FROM dependencies as builder
COPY cmd/ ./cmd
COPY internal/ ./internal
COPY pkg/ ./pkg
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/library cmd/server/main.go

FROM alpine:latest AS production
WORKDIR /app
COPY --from=builder /bin/library .
COPY migrations/ /migrations
CMD ["/app/library"]