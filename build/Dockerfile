FROM golang:1.21-bullseye AS builder

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 go build -o /bin/app

FROM debian:bullseye-slim

COPY --from=builder /bin/app /bin/app

RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

EXPOSE 5000

CMD [ "/bin/app", "/etc/env/.env" ] 