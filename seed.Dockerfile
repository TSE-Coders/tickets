FROM golang:1.22.5 AS builder
WORKDIR /src
COPY . /src
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /build/app cmd/seed/main.go

FROM scratch AS runner
WORKDIR /app
COPY --from=builder /build/app .
ENTRYPOINT ["./app"]