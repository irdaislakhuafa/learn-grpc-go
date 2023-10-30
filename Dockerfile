FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o build/main src/cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/build /app/build
COPY --from=builder /app/etc /app/etc
EXPOSE 50051
CMD [ "./build/main", "-env", "local" ]