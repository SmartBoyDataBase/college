FROM golang:alpine as builder
RUN apk add git
COPY . /go/src/sbdb-college
ENV GO111MODULE on
WORKDIR /go/src/sbdb-college
RUN go get && go build

FROM alpine
MAINTAINER longfangsong@icloud.com
COPY --from=builder /go/src/sbdb-college/sbdb-college /
WORKDIR /
CMD ./sbdb-college
ENV PORT 8000
EXPOSE 8000