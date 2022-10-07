FROM golang:1.18-alpine as builder

WORKDIR /api

COPY go.mod go.sum ./

COPY .env ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o ./akademik

#FROM alpine
#
#RUN apk add --no-cache \
#    ca-certificates \
#    && rm -rf /var/cache/apk/*
#
#WORKDIR /api
#
#COPY --from=build /api/akademik /api/akademik
##uncomment below if .env needed
#COPY --from=build /api/.env /api/.env

EXPOSE 8888
ENTRYPOINT ["/api/akademik"]