FROM golang:1.23.1 AS builder

WORKDIR /user-app

COPY . ./
RUN go mod download

COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -a -installsuffix cgo -o ./../myapp

FROM alpine:latest

WORKDIR /user-app

COPY --from=builder /user-app/myapp .
COPY --from=builder /user-app/internal/pkg/logs/app.log ./internal/pkg/logs/
COPY --from=builder /user-app/.env .

EXPOSE 8179

CMD [ "./myapp" ]