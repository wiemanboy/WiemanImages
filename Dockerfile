FROM golang AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./main.go ./
COPY .env ./
COPY ./config ./config
COPY ./src ./src
RUN CGO_ENABLED=0 GOOS=linux go build -o /build

FROM golang:alpine

COPY --from=builder /build /build

ENV PORT=8000

EXPOSE $PORT

CMD ["/build"]