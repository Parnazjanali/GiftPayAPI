From golang:alpine as builder 

WORKDIR /app

copy . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go


FROM alpine:latest AS runtime

WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8082
CMD ["./main"]


