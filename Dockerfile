FROM golang:1.17.5-alpine3.14

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV APP=app

WORKDIR /dev/$APP

COPY . .

RUN apk --no-cache add curl

# https://github.com/cosmtrek/air
RUN curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
    && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air

RUN set -eux && \
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 && \
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1