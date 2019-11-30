##################################
# Build small image without code #
##################################

FROM alpine:3.10
WORKDIR /dialog

# copy local "dialog" binary to docker.
COPY dialog .

ENV apikey ""

EXPOSE 8080
CMD ["sh", "dialog"]