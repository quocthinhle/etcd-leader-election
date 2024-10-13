# TODO: Add a multi-stage build to reduce the size of the final image

FROM golang:1.23.1-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /out/main ./

ENTRYPOINT ["/out/main"]