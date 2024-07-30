FROM golang AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./main.go ./
COPY ./config ./config
COPY ./src ./src
RUN CGO_ENABLED=0 GOOS=linux go build -o /build

FROM golang:alpine

COPY --from=builder /build /build
COPY read_secrets.sh /read_secrets.sh

ENV PORT=8000

EXPOSE $PORT

ENTRYPOINT ["/read_secrets.sh", "/build"]