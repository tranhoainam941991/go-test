FROM golang:1.14.0

ENV GO111MODULE=on

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
EXPOSE 3000
RUN chmod +x /app
ENTRYPOINT [ "/app/m" ]