FROM golang AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build:linux

FROM golang:alpine

COPY --from=builder /app/build /build
COPY read_secrets.sh /read_secrets.sh
RUN chmod +x /read_secrets.sh

ENV PORT=8000

EXPOSE $PORT

ENTRYPOINT ["/read_secrets.sh", "/build"]