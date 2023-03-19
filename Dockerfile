FROM golang:1.20-alpine3.17

WORKDIR /app

COPY . .
 
RUN go mod download
RUN go build

EXPOSE 1800

ENV S3HOST s3.my.domain
ENV S3USER username
ENV S3SECRET secret
ENV S3BUCKET mybucket

CMD [ "/app/exm" ]
