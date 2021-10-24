FROM golang:1.14-alpine AS base
RUN mkdir /fastws
WORKDIR /fastws
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN apk add --no-cache tzdata
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/fastws ./cmd/fastws/main.go
FROM scratch
COPY --from=base /go/bin/fastws /go/bin/fastws
COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo
COPY ./config.json .
EXPOSE 6789
ENTRYPOINT ["/go/bin/fastws"]