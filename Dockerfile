# syntax=docker/dockerfile:1
# check=skip=SecretsUsedInArgOrEnv

FROM golang:1.25.0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /netcup-cli

VOLUME /secrets
ENV NETCUP_TOKEN_PATH="/secrets/token.json"


CMD ["/netcup-cli"]