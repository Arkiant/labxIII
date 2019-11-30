##################################
# Build small image without code #
##################################

FROM alpine:3.6
WORKDIR /app
RUN apk update && apk add curl
EXPOSE 8080
CMD ["sh", "/app/service"]