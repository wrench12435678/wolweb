# docker build -t wolweb .
FROM golang:alpine AS builder

LABEL org.label-schema.vcs-url="https://github.com/alchemistake/wolweb" \
    org.label-schema.url="https://github.com/alchemistake/wolweb/blob/master/README.md"

RUN mkdir /wolweb
WORKDIR /wolweb

COPY ./ /wolweb/

# Install Dependecies
RUN apk update && apk upgrade && \
    apk add --no-cache git && \
    # git clone https://github.com/alchemistake/wolweb . && \
    go mod init wolweb && \
    go get -d github.com/gorilla/handlers && \
    go get -d github.com/gorilla/mux && \
    go get -d github.com/ilyakaznacheev/cleanenv

# Build Source Files
RUN go build -o wolweb . 

# Create 2nd Stage final image
FROM alpine
WORKDIR /wolweb
COPY --from=builder /wolweb/index.html .
COPY --from=builder /wolweb/wolweb .
COPY --from=builder /wolweb/static ./static

ARG WOLWEBPORT=8089

CMD ["/wolweb/wolweb"]

EXPOSE ${WOLWEBPORT}
