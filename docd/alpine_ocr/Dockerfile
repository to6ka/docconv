FROM alpine:3.9

RUN apk add build-base abuild binutils make cmake gcc g++
RUN apk add zip
RUN apk add poppler-utils
RUN apk add wv
RUN apk add lynx
RUN apk add leptonica-dev
RUN apk add tesseract-ocr-dev
RUN apk add go
RUN apk add bash

ENV CC=/usr/bin/gcc
ENV CXX=/usr/bin/g++

# Build context must be the host GOPATH
COPY . /goworkspace

WORKDIR /goworkspace/src/code.sajari.com/docconv/docd

COPY src/code.sajari.com/docconv/docd/alpine_ocr/dependencies/* /

RUN /install.sh

ENV GOPATH=/goworkspace

RUN GOOS=linux GOARCH=amd64 go build -tags ocr -o /docd

RUN apk del build-base abuild binutils make cmake gcc g++

EXPOSE 8888

CMD ["/docd"]
