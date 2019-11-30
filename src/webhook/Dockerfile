###################
# Build go binary #
###################
 
FROM golang:1.12.1 as builder
COPY . .
WORKDIR /app
RUN go get && CGO_ENABLED=0 GOOS=linux go build -a




##################################
# Build small image without code #
##################################

FROM alpine:3.6
WORKDIR /app
RUN apk update && apk add curl
EXPOSE 8080
CMD ["sh", "/app/service"]